package order

import (
	"strconv"
	"web_tools/libs/context"
	"web_tools/libs/pages"
	BM "web_tools/models/odin/bridge/master"
	BMi "web_tools/models/odin/bridge/misc"
)

func Replacement(c *context.Context) {
	if c.User != nil {
		c.Data["Name"] = c.User.Name
	} else {
		c.Data["Name"] = "Guest"
	}
	c.Data["Title"] = "index"
	c.HTML(200, "index/index")
	//c.JSON(200, []string{"a", "b"})
}

func List(c *context.Context) {
	pager, err := strconv.Atoi(c.Req.FormValue("pager"))
	if err != nil {
		pager = 1
	}
	productList := BM.GetProductList()
	orders := BMi.GetOrderByPage(pager)
	orders.GetUserInfo()
	count := BMi.GetOrderCount()
	pages := &pages.Pages{Count: count, Page: pager, PrePage: BMi.ORDER_STEP, Url: "list"}
	//获取最新的订单
	c.Data["Title"] = "orderList"
	c.Data["Orders"] = orders
	c.Data["Pages"] = pages.Get()
	c.Data["ProductList"] = productList
	c.HTML(200, "odin/order/list")
}
