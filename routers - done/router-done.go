package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louisevanderlith/comms/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e resins.Epoxi) {
	msgCtrl := &controllers.Messages{}
	e.JoinBundle("/", roletype.Admin, mix.JSON, msgCtrl)
	e.JoinPath(e.Router().(*mux.Router), "/submit", "Submit Contact", http.MethodPost, roletype.Unknown, mix.JSON, msgCtrl.Create)

	//Message
	/*msgCtrl := &controllers.MessageController{}
	msgGroup := routing.NewRouteGroup("message", mix.JSON)
	msgGroup.AddRoute("Create Message", "", "POST", roletype.Unknown, msgCtrl.Post)
	msgGroup.AddRoute("All Messages", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, msgCtrl.Get)
	msgGroup.AddRoute("View Message", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, msgCtrl.GetOne)
	e.AddBundle(msgGroup)*/
}
