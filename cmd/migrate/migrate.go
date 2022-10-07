package main

import (
	"log"

	"github.com/teamen/kays/internal/migrate"
)

func main() {
	command := migrate.NewMigrateCommand()
	if err := command.Execute(); err != nil {
		log.Fatal("fail to execute the command.")
		// os.Exit(1)
	}
}
