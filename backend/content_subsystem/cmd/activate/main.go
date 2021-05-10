package main

import (
	"content_subsystem"
	"content_subsystem/cmd"
	"content_subsystem/content"
	"content_subsystem/inmem"
	"content_subsystem/local"
	"content_subsystem/postgres"
	"content_subsystem/server"
	"flag"
	"fmt"
	"os"

	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultPort              = "8080"
	defaultSecret			 = "secret"
	defaultDBHost			 = "database"
	defaultDBUser			 = "livewithchat"
	defaultDBPassword		 = "default"
	defaultDBName			 = "livewithchat"
	defaultDBPort			 = "5432"
)

func main() {
	var (
		addr   = cmd.EnvString("PORT", defaultPort)
		// secretKey = cmd.EnvString("SECRET", defaultSecret)

		dbHost = cmd.EnvString("DB_HOST", defaultDBHost)
		dbPort = cmd.EnvString("DB_PORT", defaultDBPort)
		dbUser = cmd.EnvString("DB_USER", defaultDBUser)
		dbPassword = cmd.EnvString("DB_PASSWORD", defaultDBPassword)
		dbName = cmd.EnvString("DB_NAME", defaultDBName)

		inmemory          = flag.Bool("inmem", false, "use in-memory repositories")
		localContent      = flag.Bool("local", false, "use local storage")
	)
	flag.Parse()

	var (
		videoDB content_subsystem.VideoRepository
		videoStorage content_subsystem.VideoStorage
		// tokenManager stream_subsystem.TokenManager
	)

// 	tokenManager = jwt.NewTokenManager(secretKey)

	if *localContent {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		println("Current Path: " + path)
		videoStorage = local.NewVideoStorage(fmt.Sprintf("%s/storage", path))
	}

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

	var st content.Service
	st = content.NewService(videoDB, videoStorage)

	srv := server.New(st)
	srv.Host.Logger.Fatal(
		srv.Host.Start(fmt.Sprintf(":%s", addr)))
}