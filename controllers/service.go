package controllers

import "github.com/astaxie/beego"


// 服务
type ServiceController struct {
	beego.Controller
}

func (c *ServiceController) Get() {
	c.TplName = "service.html"
	c.Layout = "header.html"
}