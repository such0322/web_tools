package auth

import (
	"web_tools/libs/context"
	"web_tools/models/tool"
)

func Login(c *context.Context) {
	c.HTML(200, "auth/login")
}

func PostLogin(c *context.Context) {
	account := c.Req.FormValue("account")
	passwd := c.Req.FormValue("passwd")

	user, err := tool.UserLogin(account, passwd)
	if err != nil {
		c.ServerError("登录失败", err)
		return
	}
	c.Session.Set("uid", user.ID)
	c.Session.Set("uname", user.Name)
	c.Redirect("/")
}

func Logout(c *context.Context) {
	c.Session.Flush()
	c.Redirect("/")
}
