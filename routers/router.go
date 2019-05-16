package routers

import (
	"gumeiOPH/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/escape",&controllers.EscapeController{})
	beego.Router("/room",&controllers.RoomController{})
	beego.Router("/introd",&controllers.IntrodController{})
	beego.Router("/service",&controllers.ServiceController{})
	beego.Router("/flowpath",&controllers.FlowpathController{})
	beego.Router("/bulletin",&controllers.BulletinController{})
	beego.Router("/contact",&controllers.ContactController{})
}
