package routes

import "web_tools/libs/context"

func NotFound(c *context.Context) {
	c.Data["Title"] = "Page Not Found"
	c.NotFound()
}
