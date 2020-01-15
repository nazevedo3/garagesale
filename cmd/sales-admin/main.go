package main

import (
	"flag"
	"log"
	"os"

	"github.com/nazevedo3/garagesale/internal/platform/database"
	"github.com/nazevedo3/garagesale/internal/schema"
)

func main() {

	flag.Parse()

	// =========================================================================
	// Setup Dependencies

	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	switch flag.Arg(0) {
	case "migrate":
		if err := schema.Migrate(db); err != nil {
			log.Fatal("applying migrations", err)
			os.Exit(1)
		}
		log.Println("Migrations complete")
		return
	case "seed":
		if err := schema.Seed(db); err != nil {
			log.Fatal("applying seed data", err)
		}
		log.Println("Seed data complete")
		return

	}
}
