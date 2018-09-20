package bonus

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BonusGacha struct {
	ID    int    `json:"id"`
	Title string `json:"name"`
}
type BonusGachas struct {
	Data []BonusGacha `json:"data"`
}

func (BonusGacha) TableName() string {
	return "feature"
}

func (m *BonusGachas) GetRewardNames() {
	DB.Select("id, title").Where("feature_type = ? ", "gacha").Find(&m.Data)
}

func (m *BonusGachas) GetType() string {
	return TypeGachaTicket
}

func (m *BonusGachas) ToJson() string {
	json, err := json.Marshal(m.Data)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
