package controllers

import "github.com/astaxie/beego"

// 简介
type IntrodController struct {
	beego.Controller
}

func (c *IntrodController) Get() {
	c.TplName = "introd.html"
	c.Layout = "header.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "scripts/css.html"
	c.LayoutSections["Scripts"] = "scripts/scripts.html"
}
