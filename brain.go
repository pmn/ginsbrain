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

// Add a memory
func (brain *Brain) Add(m Memory) Memory {
	m.Id = brain.GetNextId()
	m.Active = true
	m.CreatedAt = time.Now()

	brain.Memories = append(brain.Memories, m)

	// Return the memory since it's been given an Id and a CreatedAt
	return m
}

// Update a memory
func (brain *Brain) Update(m Memory) Memory {
	var idx int
	for k, v := range brain.Memories {
		if v.Id == m.Id {
			idx = k
			brain.Memories[k].Text = m.Text
		}
	}
	return brain.Memories[idx]
}

// Return the current brain
func getBrain() Brain {
	return brain
}

// Get the next Id for the brain
func (brain *Brain) GetNextId() int {
	i := 0
	for _, v := range brain.Memories {
		if v.Id > i {
			i = v.Id
		}
	}

	return i+1
}
