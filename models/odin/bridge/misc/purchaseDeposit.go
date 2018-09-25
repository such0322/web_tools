package misc

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"web_tools/libs/setting"
	AA "web_tools/models/odin/api/user"
	BM "web_tools/models/odin/bridge/master"
	"web_tools/models/tool"
)

const ORDER_STEP = 20
const DefaultProjectId = "twod"

type PurchaseDeposit struct {
	DepositId        string `gorm:"primary_key"`
	Guid             string
	AppId            string
	Platform         string
	Store            string
	Status           string
	VerifyStatus     string
	StoreEnv         string
	ProductId        string
	ProductPrimaryId int
	ManagedType      string `gorm:"type:enum('consumable','non-consumable','subscription')"`
	PurchaseType     string `gorm:"DEFAULT:'normal'"`
	IsRestored       int
	Payment          int
	OrderId          string
	VerifyToken      string
	InsDate          time.Time
	LastModDate      time.Time
	ChannelId        string
	User             *AA.User    `gorm:"-"`
	Product          *BM.Product `gorm:"-"`
}

func LoadOrderByID(depositId string) *PurchaseDeposit {
	var order PurchaseDeposit
	DB.Where("deposit_id = ?", depositId).First(&order)
	return &order
}

func (order *PurchaseDeposit) LoadProduct() *PurchaseDeposit {
	order.Product = BM.LoadProduct(order.ProductPrimaryId)
	return order
}
func (order *PurchaseDeposit) LoadUser() *PurchaseDeposit {
	order.User = AA.LoadUserByGuid(order.Guid)
	return order
}

func GetOrderByPage(pager int, cid int) Orders {
	var pds []PurchaseDeposit
	offset := 0
	if pager > 0 {
		offset = (pager - 1) * ORDER_STEP
	}
	db := DB.Limit(ORDER_STEP).Offset(offset).Order("ins_date desc")
	if cid > 0 {
		user := AA.LoadUserByCid(cid)
		if user.Cid == 0 {
			return pds
		}
		guid := strings.Split(user.Muid, ":")[1]
		db = db.Where("guid = ?", guid)
	}
	db.Find(&pds)
	return pds
}
func GetOrderCount(cid int) (count int) {
	var pd PurchaseDeposit
	db := DB.Model(&pd)
	if cid > 0 {
		user := AA.LoadUserByCid(cid)
		if user.Cid == 0 {
			return 0
		}
		guid := strings.Split(user.Muid, ":")[1]
		db = db.Where("guid = ?", guid)
	}
	db.Count(&count)
	return count
}

type Orders []PurchaseDeposit

func (orders Orders) LoadUserInfo() {
	var guids []string
	for _, vo := range orders {
		guids = append(guids, vo.Guid)
	}
	users := AA.GetUsersByGuids(guids)
	//users.LoadCharaInfo()
	var userMap = make(map[string]AA.User)
	for _, vo := range users {
		userMap[vo.Muid] = vo
	}
	for ko, vo := range orders {
		user := userMap[AA.UserMuidPrefix+vo.Guid]
		orders[ko].User = &user
	}
}

func (order *PurchaseDeposit) CallAPIReplacement(channelId string) error {
	postUrl := setting.Cfg.Section("").Key("PAY_API").String()
	pp := tool.PostParams{
		ProjectId: DefaultProjectId,
		ChannelId: channelId,
		UserId:    order.Guid,
		OrderId:   order.DepositId,
		ProductId: order.ProductId,
		Money:     order.Product.Payment * 100,
	}
	jsonPP, err := json.Marshal(pp)
	if err != nil {
		return err
	}

	params := strings.NewReader(string(jsonPP))
	resp, err := http.Post(postUrl+"/orderReissued", "application/x-www-form-urlencoded", params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("http status not ok!")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respJson := tool.RespData{}
	err = json.Unmarshal(body, &respJson)
	if err != nil {
		return err
	}
	return nil
}

//直购商品进行补偿
