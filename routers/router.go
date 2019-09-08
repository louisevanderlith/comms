package routers

import (
	"github.com/louisevanderlith/comms/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e resins.Epoxi) {
	e.JoinBundle("/", roletype.User, mix.JSON, &controllers.Messages{})
	//Message
	/*msgCtrl := &controllers.MessageController{}
	msgGroup := routing.NewRouteGroup("message", mix.JSON)
	msgGroup.AddRoute("Create Message", "", "POST", roletype.Unknown, msgCtrl.Post)
	msgGroup.AddRoute("All Messages", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, msgCtrl.Get)
	msgGroup.AddRoute("View Message", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, msgCtrl.GetOne)
	e.AddBundle(msgGroup)*/
}
