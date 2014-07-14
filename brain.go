package main

import (
	"time"
)

type Memory struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
