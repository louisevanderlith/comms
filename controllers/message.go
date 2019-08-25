package controllers

import (
	"net/http"

	"github.com/louisevanderlith/comms/core"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type MessageController struct {
}

// @Title SendMessage
// @Description Sends a Message
// @Param	body	body	comms.Message	true	"body for message content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *MessageController) Post(ctx context.Contexer) (int, interface{}) {
	var message core.Message
	err := ctx.Body(&message)

	if err != nil {
		return http.StatusBadRequest, err
	}

	//Can be detected from Origin
	message.TemplateName = "default.html"

	err = message.SendMessage()

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, "Message has been sent."
}

// @Title GetMessages
// @Description Gets all Messages
// @Success 200 {[]comms.Message]} []comms.Message]
// @router /all/:pagesize [get]
func (req *MessageController) Get(ctx context.Contexer) (int, interface{}) {
	page, size := ctx.GetPageData()
	result := core.GetMessages(page, size)

	return http.StatusOK, result
}

// @Title GetMessage
// @Description Gets a comms message
// @Param	key			path	string 	true		"comms key"
// @Success 200 {core.Message} core.Message
// @router /:key [get]
func (req *MessageController) GetOne(ctx context.Contexer) (int, interface{}) {
	siteParam := ctx.FindParam("key")

	key, err := husk.ParseKey(siteParam)

	if err != nil {
		return http.StatusNotFound, err
	}

	result, err := core.GetMessage(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, result
}
