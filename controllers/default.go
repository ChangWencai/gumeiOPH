package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"

	c.Data["Home"] = "HOME"
	c.Data["What"] = "素材"
	c.TplName = "index.html"
}
