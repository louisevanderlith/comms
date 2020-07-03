package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk"
	"log"
	"net/http"

	"github.com/louisevanderlith/comms/core"
	"github.com/louisevanderlith/droxolite/context"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	result, err := core.GetMessages(1, 10)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusFound)
		return
	}

	ctx.Serve(http.StatusOK, mix.JSON(result))
}

// @Title SendMessage
// @Description Sends a Message
// @Param	body	body	comms.Message	true	"body for message content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func CreateMessage(smtpUser, smtpPass, smtpHost string, smtpPort int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		var message core.Message
		err := ctx.Body(&message)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		message.TemplateName = "default.html"
		message.To = ctx.GetTokenInfo().GetClaimString("profile.contact.email")

		err = message.SendMessage(smtpUser, smtpPass, smtpHost, smtpPort)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		ctx.Serve(http.StatusOK, mix.JSON("Message Sent"))
	}
}

// @Title GetMessages
// @Description Gets all Messages
// @Success 200 {[]comms.Message]} []comms.Message]
// @router /all/:pagesize [get]
func SearchMessages(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	page, size := ctx.GetPageData()
	result, err := core.GetMessages(page, size)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusFound)
		return
	}

	ctx.Serve(http.StatusOK, mix.JSON(result))
}

// @Title GetMessage
// @Description Gets a comms message
// @Param	key			path	string 	true		"comms key"
// @Success 200 {core.Message} core.Message
// @router /:key [get]
func ViewMessage(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	siteParam := ctx.FindParam("key")

	key, err := husk.ParseKey(siteParam)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusFound)
		return
	}

	result, err := core.GetMessage(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	ctx.Serve(http.StatusOK, mix.JSON(result))
}
