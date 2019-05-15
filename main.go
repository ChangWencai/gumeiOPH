package main

import (
	"github.com/astaxie/beego/logs"
	_ "gumeiOPH/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/gumei.log"}`)
	logs.EnableFuncCallDepth(true)

	beego.SetStaticPath("static", "static")
	beego.Run()
}

