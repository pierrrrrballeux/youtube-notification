package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func GetHMACSecret() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println(err)
	}

	return hex.EncodeToString(bytes)
}

func send_webhook(data Feed){
	
}