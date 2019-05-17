package afterEnd

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// 后台登陆
type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.TplName = "admin/index.html"
}

func (c *AdminController) Login() {
	name := c.GetString("name")
	psword := c.GetString("password")
	logs.Info(name, psword)
	c.TplName = "end/index"
}