package controller

import (
	"go2/uid"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MsgController interface {
	GetAll() []uid.Person
	ShowAll(ctx *gin.Context)
}

type controller struct {
	messages []uid.Person
}

func (c *controller) GetAll() []uid.Person {
	return c.messages
}

func (c *controller) ShowAll(ctx *gin.Context) {
	datas := c.GetAll()
	data := gin.H{
		"message": datas,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
