package context

import (
	"fmt"
	"net/http"
	"web_tools/models/tool"

	"github.com/go-macaron/macaron"
	"github.com/go-macaron/session"
)

type Context struct {
	*macaron.Context
	Session session.Store

	User *tool.User
}

//todo 这边的依赖注入还不是很懂，暂时这样可以使用,有点懂了
func Contexter() macaron.Handler {
	return func(ctx *macaron.Context, sess session.Store) {
		c := &Context{
			Context: ctx,
			Session: sess,
		}

		c.User = tool.UserSignin(sess)
		c.Data["User"] = c.User
		ctx.Map(c)
	}
}

//404页面
func (c *Context) NotFound() {
	c.Handle(http.StatusNotFound, "", nil)
}

//500页面
func (c *Context) ServerError(title string, err error) {
	c.Handle(http.StatusInternalServerError, title, err)
}

//错误页面通用接口
func (c *Context) Handle(status int, title string, err error) {
	switch status {
	case http.StatusNotFound:
		c.Data["Title"] = "Page Not Found"
	case http.StatusInternalServerError:
		c.Data["Title"] = "Internal Server Error"
		//log.Error("%s: %v", title, err)
		//log.Error(3, "%s: %v", title, err)
		//if !setting.ProdMode || (c.IsLogged && c.User.IsAdmin) {
		//	c.Data["ErrorMsg"] = err
		//}
	}
	c.HTML(status, fmt.Sprintf("status/%d", status))
}

func NeedLogin() macaron.Handler {
	return func(c *Context) {
		if c.User == nil {
			fmt.Println("bbb")
			c.Redirect("auth/login")
			return
		}
	}
}
