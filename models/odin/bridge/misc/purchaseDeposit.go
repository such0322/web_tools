package misc

import (
	"time"
	AA "web_tools/models/odin/api/user"
)

const ORDER_STEP = 20

type PurchaseDeposit struct {
	DepositId        string
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
	User             AA.User `gorm:"-"`
}

func GetOrderByPage(pager int) Orders {
	var pds []PurchaseDeposit
	offset := 0
	if pager > 0 {
		offset = (pager - 1) * ORDER_STEP
	}
	DB.Limit(ORDER_STEP).Offset(offset).Find(&pds)
	return pds
}
func GetOrderCount() (count int) {
	var pd PurchaseDeposit
	DB.Model(&pd).Count(&count)
	return count
}

type Orders []PurchaseDeposit

func (orders Orders) GetUserInfo() {
	var guids []string
	for _, vo := range orders {
		guids = append(guids, vo.Guid)
	}
	users := AA.GetUsersByGuids(guids)
	var userMap = make(map[string]AA.User)
	for _, vo := range users {
		userMap[vo.Muid] = vo
	}
	for ko, vo := range orders {
		orders[ko].User = userMap[AA.UserMuidPrefix+vo.Guid]
	}
}
