package master

import (
	"time"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

const (
	ChannelHOUTU  = "houtu"
	ChannelBILI   = "bili"
	ChannelHUAWEI = "huawei"
	ChannelIOS    = "ios"
	ChannelMI     = "mi"
	ChannelMZ     = "mz"
	ChannelOPD2C  = "opd2c"
	ChannelOPPO   = "oppo"
	ChannelQIHOO  = "qihoo"
	ChannelQUICK  = "quick"
	ChannelUC     = "uc"
	ChannelVIVO   = "vivo"
)

type Model struct {
	InsDate time.Time
}

func GetChannels() []string {
	channels := []string{
		ChannelHOUTU,
		ChannelBILI,
		ChannelHUAWEI,
		ChannelIOS,
		ChannelMI,
		ChannelMZ,
		ChannelOPD2C,
		ChannelOPPO,
		ChannelQIHOO,
		ChannelQUICK,
		ChannelUC,
		ChannelVIVO,
	}
	return channels
}
