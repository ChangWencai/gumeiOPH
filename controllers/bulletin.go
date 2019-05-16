package controllers

import "github.com/astaxie/beego"

// 公告栏
type BulletinController struct {
	beego.Controller
}

func (c *BulletinController) Get() {
	c.TplName = "bulletin.html"
	c.Layout = "header.html"
}