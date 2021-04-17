package main

import (
	"flag"
	"fmt"
	"stream_subsystem"
	"stream_subsystem/cmd"
	"stream_subsystem/inmem"
	"stream_subsystem/local"
	"stream_subsystem/postgres"
	"stream_subsystem/server"
	"stream_subsystem/stream"

	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultPort              = "8080"
	defaultDBHost			 = "database"
	defaultDBUser			 = "livewithchat"
	defaultDBPassword		 = "default"
	defaultDBName			 = "livewithchat"
	defaultDBPort			 = "5432"
)

func main() {
	var (
		addr   = cmd.EnvString("PORT", defaultPort)

		dbHost = cmd.EnvString("DB_HOST", defaultDBHost)
		dbPort = cmd.EnvString("DB_PORT", defaultDBPort)
		dbUser = cmd.EnvString("DB_USER", defaultDBUser)
		dbPassword = cmd.EnvString("DB_PASSWORD", defaultDBPassword)
		dbName = cmd.EnvString("DB_NAME", defaultDBName)

		inmemory          = flag.Bool("inmem", false, "use in-memory repositories")
	)
	flag.Parse()

	var (
		videoDB stream_subsystem.VideoRepository
		contentController stream_subsystem.ContentController
	)

	contentController = local.NewContentController("/go/src/server/stream_subsystem/storage")

	if *inmemory {
		videoDB = inmem.NewVideoRepository()
	} else {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
			dbHost, dbUser, dbPassword, dbName, dbPort,
		)

		var client *gorm.DB
		client, err := gorm.Open(psql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		defer func () {
			db, err := client.DB()
			if err != nil {
				print(err)
				return
			}
			db.Close()
		}()

		videoDB = postgres.NewVideoRepository(client)
	}

	var st stream.Service
	st = stream.NewService(videoDB, contentController)

	srv := server.New(st)
	srv.Host.Logger.Fatal(
		srv.Host.Start(fmt.Sprintf(":%s", addr)))
}