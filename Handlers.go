package main

import (
        "encoding/json"
        "fmt"
        "net/http"
        "time"
        "io"
        "io/ioutil"

        "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome")
}

func EventIndex(w http.ResponseWriter, r *http.Request) {
        events := LogEvents{
                LogEvent{Message: "oh no!", Time: time.Now()},
                LogEvent{Message: "oh yeah!", Time: time.Now()},
        }

        if err := json.NewEncoder(w).Encode(events); err != nil {
                panic(err)
        }
}

func EventShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    eventId := vars["eventId"]
    fmt.Fprintln(w, "Event Id show:", eventId)
}

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

        //here: write event to the DB
        e := LogEvent{Id: 1, Severity: 1, Message: "one", Time: time.Now()}

        //now, build a response with a status created object
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusCreated)
        if err := json.NewEncoder(w).Encode(e); err != nil {
                panic(err)
        }
}