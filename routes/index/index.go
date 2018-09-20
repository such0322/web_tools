package index

import (
	"web_tools/libs/context"

	log "github.com/sirupsen/logrus"
)

func Index(c *context.Context) {
	if c.User != nil {
		c.Data["Name"] = c.User.Name
	} else {
		c.Data["Name"] = "Guest"
	}
	c.Data["Title"] = "index"
	c.HTML(200, "index/index")
	//c.JSON(200, []string{"a", "b"})
}

func Debug(c *context.Context) {
	log.Info("test debug 4")
	log.Error("test debug 5")
}
