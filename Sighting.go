package main

import "time"

//Sighting : the events stored in our app. The time and location a yellow-legged hornet was spotted.
type Sighting struct {
	ID     int64     `json:"id"`
	Xcoord float64   `json:"xcoord"`
	Ycoord float64   `json:"ycood"`
	Time   time.Time `json:"time"`
}

//Sightings : our list of Sighting objects.
type Sightings []Sighting
