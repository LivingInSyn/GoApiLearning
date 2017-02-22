package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
)
import _ "github.com/denisenkom/go-mssqldb"

//this should all get moved to a conf file and removed from git, doesn't matter right now
//because it's not on the internet
var debug = flag.Bool("debug", true, "enable debugging")
var password = flag.String("password", "GoTest!23", "the database password")
var port *int = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "localhost", "the database server")
var user = flag.String("user", "GoTester", "the database user")
var database = flag.String("database", "ErrorLogGo", "the default database")

func createConnection() *sql.DB {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
		fmt.Printf(" database name:%s\n", *database)
	}

	connString := fmt.Sprintf("server=%s;Database=%s;user id=%s;password=%s;port=%d", *server, *database, *user, *password, *port)

	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	//open the connection
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	//defer db.Close()
	return db
}

func closeConnection(db *sql.DB) {
	db.Close()
}

// GetAllEvents returns all events in the database
func GetAllEvents() []*LogEvent {
	db := createConnection()

	rows, err := db.Query("SELECT * FROM LogEvent")
	if err != nil {
		log.Fatal("query failed", err.Error())
	}
	defer rows.Close()

	var LogEvents []*LogEvent
	for rows.Next() {
		le := new(LogEvent)
		if err := rows.Scan(&le.Id, &le.Severity, &le.Message, &le.Time); err != nil {
			log.Fatal("scanning failed", err.Error())
		}
		LogEvents = append(LogEvents, le)
	}

	closeConnection(db)

	return LogEvents

}

// AddEvent adds a new event into the database
func AddEvent(le *LogEvent) *LogEvent {
	db := createConnection()
	result, err := db.Exec(
		"INSERT INTO LogEvent VALUES ($1,$2,$3)",
		le.Severity,
		le.Message,
		le.Time,
	)
	if err != nil {
		log.Fatal("insert failed", err.Error())
	}
	lastId, _ := result.LastInsertId()
	le.Id = int(lastId)

	return le
}
