package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	ch = make(chan string)
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
		return
	}
	
	session, err := discordgo.New("Bot " + os.Getenv("Token"))
	
	if err != nil {
		log.Fatalln(err)
		return 
	}

	session.AddHandler(handleReady)

	go serve()
	go func(){
		select {
		case result := <-ch:
			fmt.Println(result)
		}
	}()
	go func ()  {
		if err = session.Open(); err != nil {
			log.Fatalln(err)
			return
		}	
	}()

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch

	err = session.Close()
	if err != nil {
		log.Printf("could not close session gracefully: %s", err)
	}
}