package main

import (
	"time"

	"github.com/Liza-Developer/apiGO"
)

type embeds struct {
	Content interface{} `json:"content"`
	Embeds  []embed     `json:"embeds"`
}

type embed struct {
	Description interface{} `json:"description"`
	Color       interface{} `json:"color"`
	Footer      footer      `json:"footer"`
	Time        interface{} `json:"timestamp"`
}

type footer struct {
	Text interface{} `json:"text"`
	Icon interface{} `json:"icon_url"`
}

type skinUrls struct {
	Url     interface{} `json:"url"`
	Varient interface{} `json:"variant"`
}

type Data struct {
	Name   string `json:"name"`
	Bearer string `json:"bearer"`
	Unix   int64  `json:"unix"`
	Config string `json:"config"`
	Id     string `json:"id"`
}

type checkDetails struct {
	Error string `json:"error"`
	Sent  string `json:"sent"`
}

type SentRequests struct {
	Requests []Details
}

type Details struct {
	Bearer     string
	SentAt     time.Time
	RecvAt     time.Time
	StatusCode string
	UnixRecv   int64
	Success    bool
	Email      string
	Type       string
	Cloudfront bool
}

var (
	Bearers apiGO.MCbearers
	Acc     apiGO.Config
)
