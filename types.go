package main

import (
	"encoding/xml"
)

type Channel struct {
	id string `json:id`
	secret string `json:secret`
}

type YoutubeFeed struct {
	XMLName    xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
	Namespace  string   `xml:"xmlns:yt,attr"`
	Namespace2 string   `xml:"xmlns:media,attr"`
	ID         string   `xml:"id"`
	ChannelID  string   `xml:"yt:channelId"`
	Title      string   `xml:"title"`
	Link       string   `xml:"link"`
	Author     Author   `xml:"author"`
	Published  string   `xml:"published"`
	Entries    []Entry  `xml:"entry"`
}

type Author struct {
	Name string `xml:"name"`
	URI  string `xml:"uri"`
}

type Entry struct {
	ID         string       `xml:"id"`
	VideoID    string       `xml:"yt:videoId"`
	ChannelID  string       `xml:"yt:channelId"`
	Title      string       `xml:"title"`
	Link       string       `xml:"link"`
	Author     Author       `xml:"author"`
	Published  string       `xml:"published"`
	Updated    string       `xml:"updated"`
	MediaGroup MediaGroup   `xml:"media:group"`
}

type MediaGroup struct {
	MediaTitle   string          `xml:"media:title"`
	MediaContent MediaContent    `xml:"media:content"`
	Thumbnail    MediaThumbnail  `xml:"media:thumbnail"`
	Description  string          `xml:"media:description"`
	Community    MediaCommunity  `xml:"media:community"`
}

type MediaContent struct {
	URL    string `xml:"url,attr"`
	Type   string `xml:"type,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}

type MediaThumbnail struct {
	URL    string `xml:"url,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}

type MediaCommunity struct {
	StarRating  MediaStarRating `xml:"media:starRating"`
	Statistics  MediaStatistics `xml:"media:statistics"`
}

type MediaStarRating struct {
	Count   int     `xml:"count,attr"`
	Average float64 `xml:"average,attr"`
	Min     int     `xml:"min,attr"`
	Max     int     `xml:"max,attr"`
}

type MediaStatistics struct {
	Views int `xml:"views,attr"`
}