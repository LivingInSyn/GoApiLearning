package main

import (
        "encoding/json"
        "fmt"
        "net/http"
        "time"

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
    todoId := vars["evendId"]
    fmt.Fprintln(w, "Event Id show:", todoId)
}