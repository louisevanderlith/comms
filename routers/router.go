package routers

import (
	"github.com/louisevanderlith/comms/controllers"
	"github.com/louisevanderlith/droxolite"

	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(poxy *droxolite.Epoxy) {
	//Message
	msgCtrl := &controllers.MessageController{}
	msgGroup := droxolite.NewRouteGroup("message", msgCtrl)
	msgGroup.AddRoute("Create Message", "", "POST", roletype.Unknown, msgCtrl.Post)
	msgGroup.AddRoute("All Messages", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, msgCtrl.Get)
	msgGroup.AddRoute("View Message", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, msgCtrl.GetOne)
	poxy.AddGroup(msgGroup)
	/*ctrlmap := EnableFilters(s, host)

	msgCtrl := controllers.NewMessageCtrl(ctrlmap)

	beego.Router("/v1/message", msgCtrl, "post:Post")
	beego.Router("/v1/message/:key", msgCtrl, "get:GetOne")
	beego.Router("/v1/message/all/:pageSize", msgCtrl, "get:Get")*/
}

/*
func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.Admin

	ctrlmap.Add("/v1/message", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}
*/
