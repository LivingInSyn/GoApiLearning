package main

import "time"

type LogEvent struct {
        Id int `json:"id"`
        Severity int `json:"severity"`
        Message string `json:"message"`
        Time time.Time `json:"time"`
}

type LogEvents []LogEvent