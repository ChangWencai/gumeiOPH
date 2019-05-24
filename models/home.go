package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Home struct {
	Id        int
	Content   string
	ImageName string // OneToOne relation
}


func init() {

	orm.RegisterModel(new(Home))
	orm.Debug = true
}

func GetContent() {
	qs := orm.NewOrm().QueryTable("home")
	fmt.Println(qs)
}

func GetAll(homes []*Home)  {
	sql := `select * from Home`
	_, err := orm.NewOrm().Raw(sql).QueryRows(&homes)
	if err != nil {
		beego.Error("get das_conf error:", err.Error())
	}
	beego.Info(homes)
	return
}