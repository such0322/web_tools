package main

import (
	"os"
	"web_tools/cmd"

	"github.com/urfave/cli"
)

const APP_VAR = "1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "web_tools"
	app.Usage = "通用后台工具"
	app.Version = APP_VAR
	app.Commands = []cli.Command{
		cmd.Web,
	}
	app.Run(os.Args)
}
