package main

import (
	"authentication"
	"authentication/auth"
	"authentication/inmen"
	"authentication/jwt"
	"authentication/postgres"
	"authentication/server"
	"flag"
	"fmt"
	"os"
	"time"

	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultPort              = "8080"
	defaultSecret            = "secret"

	defaultDBHost			 = "database"
	defaultDBUser			 = "livewithchat"
	defaultDBPassword		 = "default"
	defaultDBName			 = "livewithchat"
	defaultDBPort			 = "5432"
)

func main() {
	var (
		addr   = envString("PORT", defaultPort)
		secretKey = envString("SECRET", defaultSecret)

		dbHost = envString("DB_HOST", defaultDBHost)
		dbPort = envString("DB_PORT", defaultDBPort)
		dbUser = envString("DB_USER", defaultDBUser)
		dbPassword = envString("DB_PASSWORD", defaultDBPassword)
		dbName = envString("DB_NAME", defaultDBName)

		inmemory          = flag.Bool("inmem", false, "use in-memory repositories")
	)
	flag.Parse()
	var (
		users authentication.UserRepository
		token authentication.TokenManager
	)

	if *inmemory {
		users = inmen.NewUserRepository()
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
		users = postgres.NewUserRepository(client)
	}

	token = jwt.NewTokenManager(
		secretKey,
		time.Second * 20,
		time.Second * 60)

	var au auth.Service
	au = auth.NewService(users, users, token)

	srv := server.New(au)
	srv.Host.Logger.Fatal(
		srv.Host.Start(fmt.Sprintf(":%s", addr)))
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}