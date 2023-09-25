package cmd

import (
	"log"

	"github.com/InvalidJokerDE/fpm/internal/server"
	"github.com/InvalidJokerDE/fpm/internal/server/utils"
)

func Execute() {

	if err := utils.LoadProcesses(); err != nil {
		panic(err)
	}

	log.Println("Starting server")

	err2 := server.StartServer()

	if err2 != nil {
		log.Fatal(err2)
	} else {
		log.Println("Server closed peacefully")
	}
}
