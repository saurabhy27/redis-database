package main

import (
	"fmt"

	"github.com/saurabhy27/redis-database/datastore"
	"github.com/saurabhy27/redis-database/processor"
	"github.com/saurabhy27/redis-database/server"
)

func main() {
	fmt.Println("Project Execting Started..")
	datastore := datastore.NewDataStore()
	commandProcessor := &processor.RequestProcessor{DataStore: datastore}
	args := server.ServerArgs{Port: 80}
	server.NewServer(args, commandProcessor).Start()
}
