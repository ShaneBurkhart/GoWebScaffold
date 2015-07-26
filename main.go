package main

import (
	"encoding/gob"
	"github.com/ShaneBurkhart/PlanSource/config"
	"github.com/ShaneBurkhart/PlanSource/config/db"
	"github.com/ShaneBurkhart/PlanSource/flash"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	// Add time, file name, and line number to log output.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Add Flash type to gob so we can serialize and deserialize it.
	gob.Register(flash.Flash{})

	err := db.SetupDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.CloseDB()

	if err := db.VerifyDB(); err != nil {
		log.Fatal(err)
		return
	}

	r := mux.NewRouter()
	config.SetupRoutes(r)
	config.Serve(r)
}
