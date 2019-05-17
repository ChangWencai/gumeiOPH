package afterEnd

import "github.com/astaxie/beego"

// 后台登陆
type EndController struct {
	beego.Controller
}

func (c *EndController) Get() {
	c.TplName = "end/index.html"
}