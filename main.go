package main

import (
	_ "gumeiOPH/routers"
	"github.com/astaxie/beego"
)

func main() {

	beego.SetStaticPath("../", "static")

	beego.Run()
}

