package main

import (
	"encoding/xml"
	"log"
	"github.com/gin-gonic/gin"
)

func handleCallback(c *gin.Context){
	challenge, _ := c.GetQuery("hub.challenge")
	log.Println("Challenge: %s", challenge)
	c.String(200, challenge)
	return
}

func handleNotification(c *gin.Context){
	body := c.Request.Body
	defer body.Close()

	result := Feed{};
	err := xml.NewDecoder(body).Decode(&result)
	
	if err != nil {
		log.Println(err)
	}

	log.Println("%sが\n「%s」をアップロードしました!!!", result.Author.Name,result.Title)

	c.String(200, "success")
	return
}

func handleSubscribe(c *gin.Context){
	return
}

func main(){
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) { ctx.String(200, "Hello World!!!") })
	r.GET("/hub", handleCallback)
	r.POST("/hub", handleNotification)
	r.POST("/subscribe", handleSubscribe)

	r.Run()
}