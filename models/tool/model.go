package tool

import (
	"time"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type PostParams struct {
	ProjectId string `json:"projectId"`
	ChannelId string `json:"channelId"`
	UserId    string `json:"userId"`
	OrderId   string `json:"orderId"`
	ProductId string `json:"productId"`
	Money     int    `json:"money"`
}

type RespData struct {
	State int     `json:"state"`
	Data  ApiData `json:"data"`
}

type ApiData struct {
	ChannelId    string `json:"channelId"`
	UserId       string `json:"userId"`
	OrderId      string `json:"orderId"`
	ProductId    string `json:"productId"`
	Money        int    `json:"money"`
	PurchaseDate Time   `json:"purchaseDate"`
	Extension    string `json:"extension"`
	CreateTime   Time   `json:"createTime"`
}

type Time time.Time

const (
	timeFormart = "2006-1-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return err
}
