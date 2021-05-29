package main

import (
	"flag"
	"fmt"
	"os"
	"stream_subsystem"
	"stream_subsystem/chat"
	"stream_subsystem/cmd"
	"stream_subsystem/inmem"
	"stream_subsystem/internal"
	"stream_subsystem/jwt"
	"stream_subsystem/local"
	"stream_subsystem/postgres"
	"stream_subsystem/server"
	"stream_subsystem/stream"

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

	defaultContentUrl		 = "http://content_subsystem:8080/api/v1/content/videos"
)

func main() {
	var (
		addr   = cmd.EnvString("PORT", defaultPort)
		secretKey = cmd.EnvString("SECRET", defaultSecret)

		dbHost = cmd.EnvString("DB_HOST", defaultDBHost)
		dbPort = cmd.EnvString("DB_PORT", defaultDBPort)
		dbUser = cmd.EnvString("DB_USER", defaultDBUser)
		dbPassword = cmd.EnvString("DB_PASSWORD", defaultDBPassword)
		dbName = cmd.EnvString("DB_NAME", defaultDBName)

		contentUrl		  = cmd.EnvString("CONTENT_URL", defaultContentUrl)

		inmemory          = flag.Bool("inmem", false, "use in-memory repositories")
		localContent      = flag.Bool("local", false, "use local storage")
	)
	flag.Parse()

	var (
		messageDB stream_subsystem.MessageRepository
		videoDB stream_subsystem.VideoRepository
		contentController stream_subsystem.ContentController
		tokenManager stream_subsystem.TokenManager
	)

	tokenManager = jwt.NewTokenManager(secretKey)

	if *localContent {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		contentController = local.NewContentController(fmt.Sprintf("%s/storage", path))
	} else {
		contentController = internal.NewContentController(contentUrl)
	}

	if *inmemory {
		videoDB = inmem.NewVideoRepository()
		messageDB = inmem.NewMessageRepository()
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
		messageDB = postgres.NewMessageRepository(client)
	}

	var st stream.Service
	st = stream.NewService(videoDB, contentController)

	var ch chat.Service
	ch = chat.NewService(messageDB)

	srv := server.New(ch, st, tokenManager)
	srv.Host.Logger.Fatal(
		srv.Host.Start(fmt.Sprintf(":%s", addr)))
}