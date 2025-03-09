package main

import (
	"encoding/xml"
)

type Channel struct {
	id string `json:id`
	secret string `json:secret`
}

type YoutubeFeed struct {
	XMLName  xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
	LinkHub  string   `xml:"link[rel='hub'],attr"`
	LinkSelf string   `xml:"link[rel='self'],attr"`
	Title    string   `xml:"title"`
	Updated  string   `xml:"updated"`
	Entries  []Entry  `xml:"entry"`
}

type Entry struct {
	ID        string  `xml:"id"`
	VideoID   string  `xml:"yt:videoId"`
	ChannelID string  `xml:"yt:channelId"`
	Title     string  `xml:"title"`
	Link      string  `xml:"link[rel='alternate'],attr"`
	Author    Author  `xml:"author"`
	Published string  `xml:"published"`
	Updated   string  `xml:"updated"`
}

type Author struct {
	Name string `xml:"name"`
	URI  string `xml:"uri"`
}