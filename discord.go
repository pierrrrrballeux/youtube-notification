package main

import (
	"log"
	"github.com/bwmarrin/discordgo"
)

func handleReady(s *discordgo.Session, _ *discordgo.Ready){
	log.Printf("[DISOCRD_BOT] %s\n", s.State.User.Username)
}

func handleMessageCreate(s *discordgo.Session, msg *discordgo.MessageCreate){}
func handleClick(s *discordgo.Session, interaction *discordgo.InteractionCreate){
	if interaction.Type == discordgo.InteractionModalSubmit {

	}
}