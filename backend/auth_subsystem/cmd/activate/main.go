package main

import (
	"auth_subsystem"
	"auth_subsystem/auth"
	"auth_subsystem/cmd"
	"auth_subsystem/inmem"
	"auth_subsystem/jwt"
	"auth_subsystem/postgres"
	"auth_subsystem/server"
	"flag"
	"fmt"
	"strconv"
	"time"

	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		users auth_subsystem.UserRepository
		token auth_subsystem.TokenManager
	)

	if *inmemory {
		users = inmem.NewUserRepository()
	} else {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
			dbHost, dbUser, dbPassword, dbName, dbPort,
		)
		var client *gorm.DB
		client, err := gorm.Open(psql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
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
