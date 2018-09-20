package bonus

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
)

const TypeArtifact = "artifact"
const TypeEP = "ep"
const TypeEXP = "exp"
const TypeGachaTicket = "gacha_ticket"
const TypeGuildCoin = "guild_coin"
const TypeGuildEXP = "guild_exp"
const TypeItem = "item"
const TypeItemBundle = "item_bundle"
const TypeLUPI = "lupi"
const TypeMap = "map"
const TypeMileage = "mileage"
const TypeORB = "orb"
const TypeRandomBox = "random_box"
const TypeRune = "rune"
const TypeToken = "token"
const TypeVIP = "vip"
const TypeMissionItem = "mission_item"
const TypeLetter = "letter"
const TypeChest = "chest"

var DB *gorm.DB

var RewardType = map[string]string{
	TypeORB:         "水晶",
	TypeLUPI:        "金币",
	TypeMileage:     "稀有勋章",
	TypeArtifact:    "装备",
	TypeRune:        "符石",
	TypeMap:         "地图",
	TypeGachaTicket: "扭蛋券",
	TypeItem:        "道具",
	TypeItemBundle:  "道具包",
	TypeEXP:         "经验值",
}

type BonusData struct {
	Type        string `json:"type"`
	ID          int    `json:"id,string"`
	Quantity    int    `json:"quantity,string"`
	AddContents string `json:"add_contents"`
}

// func (bd *BonusData) MarshalJSON() ([]byte, error) {
// 	return []byte{}, nil
// }

func (b *BonusData) CreateRewards(types []string, ids []string, qs []string) []BonusData {
	var bds []BonusData
	for k, v := range types {
		id, err := strconv.Atoi(ids[k])
		if err != nil {
			// TODO
			fmt.Println(err)
		}
		q, err := strconv.Atoi(qs[k])
		if err != nil {
			// TODO
			fmt.Println(err)
		}
		//if v == "maps" {
		//	v = "map"
		//}
		bd := BonusData{Type: v, ID: id, Quantity: q, AddContents: "[]"}
		bds = append(bds, bd)
	}
	return bds
}

type Bonus struct {
	Type string `json:"type"`
	Name string `json:"name"`

	Obj Bonuser `json:"obj"`
}

type Bonuser interface {
	GetRewardNames()
	GetType() string
}

func NewBonus(bonusType string) (*Bonus, error) {
	name, ok := RewardType[bonusType]
	if !ok {
		return nil, errors.New("NewBonus: type error")
	}
	bonus := Bonus{Type: bonusType, Name: name}
	switch bonusType {
	case TypeItem:
		obj := BonusItems{}
		bonus.Obj = &obj
		return &bonus, nil
	case TypeItemBundle:
		obj := BonusItemBundles{}
		bonus.Obj = &obj
		return &bonus, nil
	case TypeArtifact:
		obj := BonusArtifacts{}
		bonus.Obj = &obj
		return &bonus, nil
	case TypeEXP:
		obj := BonusExp{}
		bonus.Obj = &obj
		return &bonus, nil
	case TypeLUPI:
		obj := BonusLupi{}
		bonus.Obj = &obj
		return &bonus, nil
	case TypeGachaTicket:
		obj := BonusGachas{}
		bonus.Obj = &obj
		return &bonus, nil
	case TypeRune:
		obj := BonusRunes{}
		bonus.Obj = &obj
		return &bonus, nil
	case TypeMileage:
		obj := BonusMileage{}
		bonus.Obj = &obj
		return &bonus, nil
	case TypeORB:
		obj := BonusOrbs{}
		bonus.Obj = &obj
		return &bonus, nil
	case TypeMap:
		obj := BonusMaps{}
		bonus.Obj = &obj
		return &bonus, nil
	}

	return nil, errors.New("NewBonus:error")
}

func (b *Bonus) GetRewardNames() {
	b.Obj.GetRewardNames()
}

func (b *Bonus) ToJson() string {
	json, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
