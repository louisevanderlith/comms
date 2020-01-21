package message

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/comms/core"
	"github.com/louisevanderlith/droxo"
	"github.com/louisevanderlith/husk"
	"net/http"
)

type Messages struct {
}

func Get(c *gin.Context) {
	result := core.GetMessages(1, 10)

	c.JSON(http.StatusOK, result)
}

// @Title SendMessage
// @Description Sends a Message
// @Param	body	body	comms.Message	true	"body for message content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func Create(c *gin.Context) {
	var message core.Message
	err := c.Bind(&message)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	//Can be detected from Origin
	message.TemplateName = "default.html"

	err = message.SendMessage()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, "Message has been sent.")
}

// @Title GetMessages
// @Description Gets all Messages
// @Success 200 {[]comms.Message]} []comms.Message]
// @router /all/:pagesize [get]
func Search(c *gin.Context) {
	page, size := droxo.GetPageData(c.Param("pagesize"))
	result := core.GetMessages(page, size)

	c.JSON(http.StatusOK, result)
}

// @Title GetMessage
// @Description Gets a comms message
// @Param	key			path	string 	true		"comms key"
// @Success 200 {core.Message} core.Message
// @router /:key [get]
func View(c *gin.Context) {
	siteParam := c.Param("key")

	key, err := husk.ParseKey(siteParam)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	result, err := core.GetMessage(key)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, result)
}
