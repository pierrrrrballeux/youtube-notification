package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleCallback(c *gin.Context){
	challenge, _ := c.GetQuery("hub.challenge")
	log.Printf("Challenge: %s\n", challenge)
	c.String(200, challenge)
	return
}

func handleNotification(c *gin.Context){
	body := c.Request.Body
	defer body.Close()

	result := YoutubeFeed{};
	err := xml.NewDecoder(body).Decode(&result)
	
	if err != nil {
		log.Println(err)
		return
	}

	jsonData, _ := json.Marshal(result)
	ch <- string(jsonData)

	log.Printf("%sが\n「%s」をアップロードしました!!!\n", result.Entries[0].Author.Name,result.Entries[0].Link)

	c.String(200, "success")
	return
}

func handleSubscribe(c *gin.Context){
	_, err := http.Post("https://pubsubhubbub.appspot.com/", "application/x-www-form-urlencoded", nil)
	if err != nil {
		log.Println(err)
		return 
	}

	return
}

func serve(){
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) { ctx.String(200, "Hello World!!!") })
	r.GET("/hub", handleCallback)
	r.POST("/hub", handleNotification)
	r.POST("/subscribe", handleSubscribe)

	r.Run(":8080")
}