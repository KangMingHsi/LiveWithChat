package main

import (
	"chat_subsystem"
	"chat_subsystem/chat"
	"chat_subsystem/cmd"
	"chat_subsystem/inmem"
	"chat_subsystem/jwt"
	"chat_subsystem/postgres"
	"chat_subsystem/server"
	"flag"
	"fmt"

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
		secretKey = cmd.EnvString("SECRET", defaultSecret)

		dbHost = cmd.EnvString("DB_HOST", defaultDBHost)
		dbPort = cmd.EnvString("DB_PORT", defaultDBPort)
		dbUser = cmd.EnvString("DB_USER", defaultDBUser)
		dbPassword = cmd.EnvString("DB_PASSWORD", defaultDBPassword)
		dbName = cmd.EnvString("DB_NAME", defaultDBName)

		inmemory          = flag.Bool("inmem", false, "use in-memory repositories")
	)
	flag.Parse()

	var (
		messageDB chat_subsystem.MessageRepository
		tokenManager chat_subsystem.TokenManager
	)

	tokenManager = jwt.NewTokenManager(secretKey)

	if *inmemory {
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

		messageDB = postgres.NewMessageRepository(client)
	}

	var ch chat.Service
	ch = chat.NewService(messageDB)

	srv := server.New(ch, tokenManager)
	srv.Host.Logger.Fatal(
		srv.Host.Start(fmt.Sprintf(":%s", addr)))
}