package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Index just returns "welcome" as a string
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

// EventIndex returns all events in the DB in JSON format
func EventIndex(w http.ResponseWriter, r *http.Request) {
	events := GetAllEvents()

	if err := json.NewEncoder(w).Encode(events); err != nil {
		panic(err)
	}
}

// EventShow shows an individual event by Id
func EventShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	fmt.Fprintln(w, "Event Id show:", eventId)
}

// EventAdd adds an event into the DB
func EventAdd(w http.ResponseWriter, r *http.Request) {
	//create the object
	var event LogEvent
	//read the request body and check for an error
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	//close the body and check for an error
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	//create an error if there was one when we deserialized the object
	if err := json.Unmarshal(body, &event); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//write event to the DB
	e := AddEvent(&event)

	//now, build a response with a status created object
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}
