package order

import (
	"bytes"
	"encoding/json"
	"strconv"
	"web_tools/libs/context"
	"web_tools/libs/pages"
	BM "web_tools/models/odin/bridge/master"
	BMi "web_tools/models/odin/bridge/misc"
	"web_tools/models/tool"
)

func Replacement(c *context.Context) {
	depositId := c.Req.FormValue("deposit_id")
	order := BMi.LoadOrderByID(depositId).LoadUser().LoadProduct()
	channels := BM.GetChannels()
	c.Data["Order"] = order
	c.Data["Channels"] = channels
	c.HTML(200, "odin/order/replacement")
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func ReplacementProcess(c *context.Context) {
	depositId := c.Req.FormValue("deposit_id")
	channelId := c.Req.FormValue("channel_id")
	order := BMi.LoadOrderByID(depositId).LoadUser().LoadProduct()
	channels := BM.GetChannels()
	bool := false
	for _, vo := range channels {
		if vo == channelId {
			bool = true
			break
		}
	}
	if !bool {
		//	session.Flash{}
		c.Redirect("/order/replacement?deposit_id="+depositId, 200)
		return
	}

	//记录补单操作
	tool.NewReplacementLog(order.DepositId, tool.DefaultOperator)
	//向支付服务器请求补单
	order.CallAPIReplacement(channelId)
	//如果是直购商品，开启一个携程，每10秒对订单状态进行校验，finish进行补发商品

}

func List(c *context.Context) {
	pager, err := strconv.Atoi(c.Req.FormValue("pager"))
	if err != nil {
		pager = 1
	}
	productList := BM.GetProductList()
	orders := BMi.GetOrderByPage(pager)
	orders.LoadUserInfo()
	count := BMi.GetOrderCount()
	pages := &pages.Pages{Count: count, Page: pager, PrePage: BMi.ORDER_STEP, Url: "list"}
	//获取最新的订单
	c.Data["Title"] = "orderList"
	c.Data["Orders"] = orders
	c.Data["Pages"] = pages.Get()
	c.Data["ProductList"] = productList
	c.HTML(200, "odin/order/list")
}
