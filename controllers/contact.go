package controllers

import "github.com/astaxie/beego"

//联系方式
type ContactController struct {
	beego.Controller
}

func (c *ContactController) Get() {
	c.TplName = "contact.html"
	c.Layout = "header.html"
}
