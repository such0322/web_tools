package bonus

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BonusMap struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type BonusMaps struct {
	Data []BonusMap `json:"data"`
}

func (BonusMap) TableName() string {
	return "mission_explore"
}

func (m *BonusMaps) GetRewardNames() {
	DB.Select("id, name").Find(&m.Data)
}
func (m *BonusMaps) GetType() string {
	return TypeMap
}

func (m *BonusMaps) ToJson() string {
	json, err := json.Marshal(m.Data)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
