package cmd

import (
	"web_tools/libs/context"
	"web_tools/libs/logger"
	"web_tools/libs/setting"
	"web_tools/libs/template"
	"web_tools/models"
	"web_tools/routes"
	"web_tools/routes/auth"
	"web_tools/routes/index"
	OdinOrder "web_tools/routes/odin/order"

	"github.com/go-macaron/macaron"
	"github.com/go-macaron/session"
	"github.com/urfave/cli"
)

var Web = cli.Command{
	Name:   "web",
	Usage:  "后台网页工具",
	Action: runWeb,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "port,p",
			Value: "3000",
			Usage: "链接用的端口",
		},
	},
}

func runWeb(c *cli.Context) error {
	//设置日志启用
	logger.InitLogger()
	//使用经典的macaron实例
	m := macaron.New()
	//m.Use(logger.SetExectimeLog())
	//m.Use(macaron.Logger())
	m.Use(logger.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("public"))

	funcMap := template.NewFuncMap()
	m.Use(macaron.Renderer(macaron.RenderOptions{
		IndentJSON: false,
		Funcs:      funcMap,
	}))
	setting.LoadCfg()
	models.NewEngines()

	sessionOptions := session.Options{
		Provider: "memory",
		//ProviderConfig:"",
		CookieName:  "OdinToolSession",
		Gclifetime:  3600,
		Maxlifetime: 86400,
	}
	m.Use(session.Sessioner(sessionOptions))
	m.Use(context.Contexter())

	router(m)
	port := 4000
	if c.IsSet("port") {
		port = c.Int("port")
	}

	m.Run(port)
	return nil
}

func router(m *macaron.Macaron) {
	//needLogin:=context.NeedLogin()

	//路由
	m.Get("/", index.Index)
	m.Get("/debug", index.Debug)

	m.Get("/auth/login", auth.Login)
	m.Post("/auth/postLogin", auth.PostLogin)
	m.Get("/auth/logout", auth.Logout)

	m.Group("", func() {
		m.Get("/order/list", OdinOrder.List)
		m.Get("/order/replacement", OdinOrder.Replacement)
		m.Post("/order/replacementProcess", OdinOrder.ReplacementProcess)
	})

	m.NotFound(routes.NotFound)
}
