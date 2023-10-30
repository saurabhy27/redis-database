package main

import (
	"log"
	"strconv"

	"github.com/saurabhy27/redis-database/datastore"
	"github.com/saurabhy27/redis-database/processor"
	"github.com/saurabhy27/redis-database/server"
	"github.com/saurabhy27/redis-database/utils"
)

func main() {
	log.Println("Project Execting Started..")
	datastore := datastore.New()
	commandProcessor := &processor.RequestProcessor{DataStore: datastore}
	// running the project on default 80 port
	port, err := strconv.Atoi(utils.GetEnv("PORT", "80"))
	if err != nil {
		panic(err)
	}
	args := server.ServerArgs{Port: port}
	server.New(args, commandProcessor).Start()
}
