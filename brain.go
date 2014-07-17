package main

import (
	"bytes"
	"encoding/gob"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"log"
	"time"
)

const filename = "ginsbrain"
const bucketName = "david-egg-roll"

type Memory struct {
	Id      int       `json:"id"`
	Text    string    `json:"text"`
	AddedBy string    `json:"added_by"`
	AddedAt time.Time `json:"added_at"`
	Active  bool      `json:"active"`
}

type Brain struct {
	Memories []Memory `json:"memories"`
}

// Load a brain
func (brain *Brain) Load() {
	// The AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables are used.
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err.Error())
	}

	// Open Bucket
	s := s3.New(auth, aws.USEast)

	// Load the database from an S3 bucket
	bucket := s.Bucket(bucketName)

	// Create a bytes.Buffer
	n, err := bucket.Get(filename)
	if err != nil {
		panic(err)
	}

	p := bytes.NewBuffer(n)
	dec := gob.NewDecoder(p)

	err = dec.Decode(&brain)

	if err != nil {
		log.Print("There was an error loading the brain. Using a blank one.")
	}
}

// Persist a brain
func (brain *Brain) Save() {
	// Persist the database to file
	var data bytes.Buffer
	contents := gob.NewEncoder(&data)
	err := contents.Encode(brain)
	if err != nil {
		panic(err)
	}

	// The AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables are used.
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err.Error())
	}

	// Open Bucket
	s := s3.New(auth, aws.USEast)

	// Load the database from an S3 bucket
	bucket := s.Bucket(bucketName)

	err = bucket.Put(filename, data.Bytes(), "text/plain", s3.BucketOwnerFull)
	if err != nil {
		panic(err.Error())
	}
}

// Add a memory
func (brain *Brain) Add(m Memory) Memory {
	m.Id = brain.GetNextId()
	m.Active = true
	m.AddedAt = time.Now()

	brain.Memories = append(brain.Memories, m)

	// Save the brain
	brain.Save()
	// Return the memory since it's been given an Id and a AddedAt
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

// Get the next Id for the brain
func (brain *Brain) GetNextId() int {
	i := 0
	for _, v := range brain.Memories {
		if v.Id > i {
			i = v.Id
		}
	}

	return i + 1
}
