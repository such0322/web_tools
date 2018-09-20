package bonus

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BonusOrb struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type BonusOrbs struct {
	Data []BonusOrb `json:"data"`
}

func (BonusOrb) TableName() string {
	return "promotion_orb"
}

func (m *BonusOrbs) GetRewardNames() {
	DB.Select("id, name").Find(&m.Data)
}
func (m *BonusOrbs) GetType() string {
	return TypeORB
}

func (m *BonusOrbs) ToJson() string {
	json, err := json.Marshal(m.Data)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
