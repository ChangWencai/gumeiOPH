package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"gumeiOPH/models"
	_ "gumeiOPH/routers"
	"time"
)

func init()  {
	_ = orm.RegisterDriver("sqlite3", orm.DRSqlite)
	_ = orm.RegisterDataBase("default", "sqlite3", "db/content.db")
	orm.DefaultTimeLoc = time.UTC

}


func main() {


	beego.SetLogger("file", `{"filename":"logs/gumei.log"}`)
	logs.EnableFuncCallDepth(true)

	//beego.SetStaticPath("/end", "/static/end")
	//models.GetContent()
	//hs := make([]*models.Home,1)
	//models.GetAll(hs)
	//h := models.Home{2,"agdhsuis","jgfodjgifd"}
	//_,err := orm.NewOrm().Insert(h)
	//if err != nil {
	//	beego.Error("get das_conf error:", err.Error())
	//}
	//models.GetContent()
	//models.GetAll(hs)
	//fmt.Println(hs)
	home := make([]*models.Home,0)
	fmt.Println(orm.NewOrm().Raw("select * from home").QueryRows(&home))
	for k,v := range home{
		fmt.Println("k ",k, " v", v)
	}

	h := models.Home{
		Id:3,
		Content:   "new cccc",
		ImageName: "new imaaaaa",
	}

	//_, _ = orm.NewOrm().Insert(&h)
	_, _ = orm.NewOrm().Update(&h)

	beego.Run()
}
