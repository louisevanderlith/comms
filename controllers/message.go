package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/comms/core"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/control"
)

type MessageController struct {
	control.APIController
}

func NewMessageCtrl(ctrlMap *control.ControllerMap) *MessageController {
	result := &MessageController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title SendMessage
// @Description Sends a Message
// @Param	body	body	comms.Message	true	"body for message content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *MessageController) Post() {
	var message core.Message
	json.Unmarshal(req.Ctx.Input.RequestBody, &message)

	if message.To == "" {
		message.To = beego.AppConfig.String("defaultEmail")
	}

	err := message.SendMessage()

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, "Message has been sent.")
}

// @Title GetMessages
// @Description Gets all Messages
// @Success 200 {[]comms.Message]} []comms.Message]
// @router /all/:pagesize [get]
func (req *MessageController) Get() {
	page, size := req.GetPageData()
	result := core.GetMessages(page, size)

	req.Serve(http.StatusOK, nil, result)
}

// @Title GetMessage
// @Description Gets a comms message
// @Param	key			path	string 	true		"comms key"
// @Success 200 {core.Message} core.Message
// @router /:key [get]
func (req *MessageController) GetOne() {
	siteParam := req.Ctx.Input.Param(":key")

	key, err := husk.ParseKey(siteParam)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	result, err := core.GetMessage(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, result)
}
