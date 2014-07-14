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
	var memories []Memory
	var memory Memory
	memory.Id = 0
	memory.Text = "Some circus subculture shit"
	memory.Active = true

	var memory2 Memory
	memory2.Id = 1
	memory2.Text = "cetacean junk"
	memory2.Active = true

	memories = append(memories, memory)
	memories = append(memories, memory2)

	brain.Memories = memories
}

// Persist a brain
func (brain *Brain) Save() {

}

// Return the current brain
func getBrain() Brain {
	return brain
}
