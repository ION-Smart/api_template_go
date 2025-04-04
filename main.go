package main

import (
	"log"

	"github.com/ION-Smart/api_template_go/internal/api"
	"github.com/ION-Smart/api_template_go/internal/db"
)

var (
	addr = ":8888"
)
var logger *log.Logger

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	sv := api.NewAPIServer(addr, db)
	if err := sv.Run(); err != nil {
		log.Fatal(err)
	}
}
