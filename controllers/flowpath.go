package controllers

import "github.com/astaxie/beego"

type FlowpathController struct {
	beego.Controller
}

func (c *FlowpathController) Get() {
	c.TplName = "flowpath.html"
	c.Layout = "header.html"
}
