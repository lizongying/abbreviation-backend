package main

import (
	"abbreviation/app"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	app.InitConfig()
	app.InitDb()
	server := app.Conf.Server
	gin.SetMode(server.Mode)
	r := gin.New()
	r.Use(cors.Default())

	r.Static("/static", "./dist/static")
	r.GET("/api/abbr/:keyword", func(c *gin.Context) {
		var reqAbbr app.ReqAbbr
		if err := c.ShouldBindUri(&reqAbbr); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": err})
			return
		}
		keyword := c.Param("keyword")
		data := app.GetAbbr(keyword)
		c.String(http.StatusOK, data)
	})
	if err := r.Run(server.Url); err != nil {
		log.Fatalln(err)
	}
}
