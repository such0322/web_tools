package bonus

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BonusItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type BonusItems struct {
	Data []BonusItem `json:"data"`
}

func (BonusItem) TableName() string {
	return "item"
}

func (m *BonusItems) GetRewardNames() {
	DB.Select("id, name").Find(&m.Data)
}
func (m *BonusItems) GetType() string {
	return TypeItem
}

func (m *BonusItems) ToJson() string {
	json, err := json.Marshal(m.Data)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
