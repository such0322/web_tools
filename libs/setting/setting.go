package setting

import (
	"fmt"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
)

func LoadCfg() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println(err)
	}
}
