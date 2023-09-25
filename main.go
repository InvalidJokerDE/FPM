package main

import (
	"log"
	"os"

	"github.com/InvalidJokerDE/fpm/cmd"
)

func main() {
	// get args
	args := os.Args[1:]

	if len(args) > 0 {
		log.Println("Args:", args)
	} else {
		log.Println("No args")
	}
	log.Println("Starting FPM")

	cmd.Execute()
}
