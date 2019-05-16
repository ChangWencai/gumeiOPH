package controllers

import "github.com/astaxie/beego"


// 逃生图
type EscapeController struct {
	beego.Controller
}

func (c *EscapeController) Get() {
	c.TplName = "escape.html"
	c.Layout = "header.html"
}