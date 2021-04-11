package main

import (
	"stream_subsystem"
	"stream_subsystem/cmd"
	"stream_subsystem/inmem"
	"stream_subsystem/stream"
	"stream_subsystem/server"
	"flag"
	"fmt"
)

const (
	defaultPort              = "8080"
	// defaultDBHost			 = "database"
	// defaultDBUser			 = "livewithchat"
	// defaultDBPassword		 = "default"
	// defaultDBName			 = "livewithchat"
	// defaultDBPort			 = "5432"
)

func main() {
	var (
		addr   = cmd.EnvString("PORT", defaultPort)

		// dbHost = cmd.EnvString("DB_HOST", defaultDBHost)
		// dbPort = cmd.EnvString("DB_PORT", defaultDBPort)
		// dbUser = cmd.EnvString("DB_USER", defaultDBUser)
		// dbPassword = cmd.EnvString("DB_PASSWORD", defaultDBPassword)
		// dbName = cmd.EnvString("DB_NAME", defaultDBName)

		inmemory          = flag.Bool("inmem", false, "use in-memory repositories")
	)
	flag.Parse()

	var (
		videoDB stream_subsystem.VideoRepository
	)

	if *inmemory {
		videoDB = inmem.NewVideoRepository()
	} else {

	}

	var st stream.Service
	st = stream.NewService(videoDB)

	srv := server.New(st)
	srv.Host.Logger.Fatal(
		srv.Host.Start(fmt.Sprintf(":%s", addr)))
}