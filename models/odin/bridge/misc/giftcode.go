package misc

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//可重复使用账号
const TYPE_REPEAT_USE = 1

//单批次只可使用一次
const TYPE_ONLY_ONE = 2

//单批次可使用多个
const TYPE_MULTI_USE = 3

const GIFTCODE_STEP = 100
const CODE_LEN = 10

const STATUS_OPEN = 1

const QUANTITY_DEFAULT = 1

type GiftCode struct {
	Model
	Code        string
	Batch       int
	Channel     string
	Type        int
	Quantity    int
	Package     string
	Status      int
	StartDate   time.Time
	EndDate     time.Time
	LastModDate time.Time
}

func GetGiftCodeByPage(pager int) []GiftCode {
	var giftcodes []GiftCode
	offset := 0
	if pager > 0 {
		offset = (pager - 1) * GIFTCODE_STEP
	}
	DB.Limit(GIFTCODE_STEP).Offset(offset).Find(&giftcodes)
	return giftcodes
}

func GetGiftCodeCount() (count int) {
	var gc GiftCode
	DB.Model(&gc).Count(&count)
	return count
}
func GetRandomCode() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	rs := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < CODE_LEN; i++ {
		rs = append(rs, bytes[r.Intn(len(bytes))])
	}
	return string(rs)
}

func CreateGiftCodes(data map[string]interface{}) {
	switch data["type"] {
	case TYPE_REPEAT_USE:
		m := GiftCode{}
		m.createRepeat(data)
	case TYPE_ONLY_ONE:
		fallthrough
	case TYPE_MULTI_USE:
		m := GiftCodes{}
		m.createBatch(data)
	}
}

func (m *GiftCode) createRepeat(data map[string]interface{}) int64 {
	if v, ok := data["code"]; !ok || v == "" {
		data["code"] = GetRandomCode()
	}
	raNum, err := m.create(data)
	if err != nil {
		log.Fatal(err)
	}
	return raNum
}

func (m *GiftCode) create(data map[string]interface{}) (int64, error) {
	sql := fmt.Sprintf("insert into gift_code (code,batch,channel,type,quantity,package,status,start_date,end_date,last_mod_date,ins_date)values (?,?,?,?,?,?,?,?,?,?,?)")
	if data["type"] != TYPE_REPEAT_USE {
		data["quantity"] = QUANTITY_DEFAULT
	}
	data["status"] = STATUS_OPEN
	t := time.Now()
	data["last_mod_date"] = t.Format("2006-01-02 15:04:05")
	data["ins_date"] = t.Format("2006-01-02 15:04:05")

	err := DB.Exec(sql, data["code"], data["batch"], data["channel"], data["type"], data["quantity"], data["package"],
		data["status"], data["start_date"], data["end_date"], data["last_mod_date"], data["ins_date"]).Error
	if err != nil {
		fmt.Println(err)
	}
	return DB.RowsAffected, nil
}

type GiftCodes struct {
	Data []GiftCode
}

var wg sync.WaitGroup

func (m *GiftCodes) createBatch(data map[string]interface{}) int64 {
	var codes []string
	// 生成codes[]
	for i := 0; i < data["quantity"].(int); i++ {
		code := GetRandomCode()
		if !inCodes(codes, code) {
			codes = append(codes, code)
		}
	}

	DB.Where(" code IN (?)", codes).Find(&m.Data)

	for _, row := range m.Data {
		codes = codesUnset(codes, row.Code)
	}
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	for _, code := range codes {
		wg.Add(1)
		dataTmp := make(map[string]interface{})
		dataTmp["code"] = code
		dataTmp["batch"] = data["batch"]
		dataTmp["channel"] = data["channel"]
		dataTmp["type"] = data["type"]
		dataTmp["quantity"] = data["quantity"]
		dataTmp["package"] = data["package"]
		dataTmp["start_date"] = data["start_date"]
		dataTmp["end_date"] = data["end_date"]

		go func(dataTmp map[string]interface{}) {
			gift := new(GiftCode)
			_, err := gift.create(dataTmp)
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(dataTmp)
	}
	wg.Wait()
	return 0
}

func codesUnset(codes []string, code string) []string {
	var i int
	for k, v := range codes {
		if v == code {
			i = k
			break
		}
	}
	return append(codes[:i], codes[i+1:]...)

}

func inCodes(codes []string, code string) bool {
	for _, v := range codes {
		if v == code {
			return true
		}
	}
	return false
}
