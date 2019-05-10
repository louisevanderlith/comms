// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/comms/controllers"
	"github.com/louisevanderlith/mango"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilters(s)

	msgCtrl := controllers.NewMessageCtrl(ctrlmap)

	beego.Router("/v1/message", msgCtrl, "post:Post")
	beego.Router("/v1/message/:key", msgCtrl, "get:GetOne")
	beego.Router("/v1/message/all/:pageSize", msgCtrl, "get:Get")
}

func EnableFilters(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.Admin

	ctrlmap.Add("/v1/message", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type", "User-Agent"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}), false)

	return ctrlmap
}
