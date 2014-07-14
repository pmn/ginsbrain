package main

import (
	"time"
)

type Memory struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
}

type Brain struct {
	Memories []Memory `json:"memories"`
}

// Load a brain
func (brain *Brain) Load() {

}

// Perisst a brain
func (brain *Brain) Save() {

}

// Return the current brain
func getBrain() Brain {
	return brain
}
