package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App container for router and db
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize init Database
func (a *App) Initialize(user, password, dbname string) {
	// connectionString := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run Start the Application
func (a *App) Run(addr string) {}
