package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "gumeiOPH/routers"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/gumei.log"}`)
	logs.EnableFuncCallDepth(true)

	//beego.SetStaticPath("/end", "/static/end")
	beego.Run()
}
