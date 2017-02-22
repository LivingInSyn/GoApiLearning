package main

import "time"

// LogEvent is the struct for our log events
type LogEvent struct {
	Id       int       `json:"id"`
	Severity int       `json:"severity"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
}

// LogEvents is an array of the LogEvent struct
type LogEvents []LogEvent
