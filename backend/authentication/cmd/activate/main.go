package main

import (
	"authentication"
	"authentication/auth"
	"authentication/cmd"
	"authentication/inmen"
	"authentication/jwt"
	"authentication/postgres"
	"authentication/server"
	"flag"
	"fmt"
	"strconv"
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

	defaultTokenDuration     = "3"
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

		tokenDuration = func() int {
			duration, _ := strconv.Atoi(cmd.EnvString("TOKEN_DURATION", defaultTokenDuration))
			return duration
		}()
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
		time.Second * time.Duration(tokenDuration))

	var au auth.Service
	au = auth.NewService(users, users, token)

	srv := server.New(au)
	srv.Host.Logger.Fatal(
		srv.Host.Start(fmt.Sprintf(":%s", addr)))
}