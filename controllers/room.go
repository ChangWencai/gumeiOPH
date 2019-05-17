package controllers

import "github.com/astaxie/beego"

// 房间布局图
type RoomController struct {
	beego.Controller
}

func (c *RoomController) Get() {
	c.TplName = "room.html"
	c.Layout = "header.html"
}
