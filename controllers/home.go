package controllers

import (
	"github.com/astaxie/beego"
)

// 首页
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.TplName = "home.html"
	c.Layout = "header.html"
}
