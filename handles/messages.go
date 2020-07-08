package handles

import (
	"errors"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/kong/tokens"
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

		email, err := emailFromClaims(ctx.GetTokenInfo())

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		message.TemplateName = "default.html"
		message.To = email

		err = message.SendMessage(smtpUser, smtpPass, smtpHost, smtpPort)

		if err != nil {
			log.Println("Send Message Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		err = ctx.Serve(http.StatusOK, mix.JSON("Message Sent"))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func emailFromClaims(clms tokens.Claimer) (string, error) {
	contacts := clms.GetClaim(tokens.KongContacts).([]interface{})

	for _, v := range contacts {
		cnt := v.(map[string]interface{})

		if cnt["Name"] == "email" {
			return cnt["Value"].(string), nil
		}
	}

	return "", errors.New("no email claim")
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
