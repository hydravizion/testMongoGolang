package handler

import (
	"go2/uid"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonPost struct {
	Uid   string `json:"UID"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type UidPost struct {
	Uid string `json:"Uid"`
}

func GetPerson(p uid.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		allPersons := p.ShowAll()
		data := gin.H{
			"datas": allPersons,
		}
		// c.JSON(http.StatusOK, data)
		c.HTML(http.StatusOK, "index.html", data)
	}
}

func AddPerson(p uid.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := PersonPost{}
		c.Bind(&requestBody)
		println(requestBody.Email)
		persondata := uid.Person{
			Uid:   requestBody.Uid,
			Name:  requestBody.Name,
			Email: requestBody.Email,
		}

		p.Add(persondata)
		c.Status(http.StatusAccepted)
	}
}

func Pingg(p uid.C2) gin.HandlerFunc {
	return func(c *gin.Context) {
		// personemel := p.GetByid(staring)
		requestBody := UidPost{}
		c.Bind(&requestBody)
		personemel := p.GetByid(requestBody.Uid)

		data := gin.H{
			"datas": personemel,
		}
		c.JSON(http.StatusOK, data)
		// c.HTML(http.StatusOK, "index.html", data)
	}
}
