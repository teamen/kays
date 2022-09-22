package main

import (
	"log"

	"github.com/teamen/kays/internal/apiserver"
)

func main() {
	command := apiserver.NewAPIServerCommand()
	if err := command.Execute(); err != nil {
		log.Fatal("fail to execute the command.")
		// os.Exit(1)
	}
}
