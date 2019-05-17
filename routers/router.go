package routers

import (
	"github.com/astaxie/beego"
	"gumeiOPH/controllers"
	"gumeiOPH/controllers/afterEnd"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/escape", &controllers.EscapeController{})
	beego.Router("/room", &controllers.RoomController{})
	beego.Router("/introd", &controllers.IntrodController{})
	beego.Router("/service", &controllers.ServiceController{})
	beego.Router("/flowpath", &controllers.FlowpathController{})
	beego.Router("/bulletin", &controllers.BulletinController{})
	beego.Router("/contact", &controllers.ContactController{})
	beego.Router("/admin", &afterEnd.AdminController{})
	beego.Router("/end", &afterEnd.EndController{})
}
