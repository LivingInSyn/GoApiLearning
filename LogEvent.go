package main

import "time"

type LogEvent struct {
        Severity int `json:"severity"`
        Message string `json:"message"`
        Time time.Time `json:"time"`
}

type LogEvents []LogEvent