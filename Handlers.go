package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Index : the default route.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

//SightingIndex : returns a list of all the available sightings.
func SightingIndex(w http.ResponseWriter, r *http.Request) {
	sightings := Sightings{
		Sighting{ID: 1, Xcoord: 46.6594744, Ycoord: 0.3629372},
		Sighting{ID: 2, Xcoord: 46.6535364, Ycoord: 0.3639805}}

	if err := json.NewEncoder(w).Encode(sightings); err != nil {
		panic(err)
	}
}

//SightingShow : return the info for a specified sighting (placeholder for the time being. waiting for DB connection.)
func SightingShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sightingID, oups := strconv.ParseInt(vars["sightingId"], 10, 64)

	if oups != nil {
		log.Fatal(oups)
	}

	sighting := Sighting{}

	sightings := Sightings{
		Sighting{ID: 1, Xcoord: 46.6594744, Ycoord: 0.3629372},
		Sighting{ID: 2, Xcoord: 46.6535364, Ycoord: 0.3639805}}

	for i := range sightings {
		if sightings[i].ID == sightingID {
			sighting = sightings[i]
		}
	}

	if err := json.NewEncoder(w).Encode(sighting); err != nil {
		panic(err)
	}
}
