package main

import (
	"auth_subsystem/cmd"
	"auth_subsystem/postgres"
	"flag"
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultDBHost			 = "database"
	defaultDBUser			 = "livewithchat"
	defaultDBPassword		 = "default"
	defaultDBName			 = "livewithchat"
	defaultDBPort			 = "5432"
)

func main() {
	var (
		dbHost = cmd.EnvString("DB_HOST", defaultDBHost)
		dbPort = cmd.EnvString("DB_PORT", defaultDBPort)
		dbUser = cmd.EnvString("DB_USER", defaultDBUser)
		dbPassword = cmd.EnvString("DB_PASSWORD", defaultDBPassword)
		dbName = cmd.EnvString("DB_NAME", defaultDBName)

		upgradeTo = flag.String("upgrade", "", "the revision of the db to upgrade to")
		downgradeTo = flag.String("downgrade", "", "the revision of the db to downgrade to")
	)

	flag.Parse()
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
	m := gormigrate.New(client, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202104082324",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&postgres.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
	})

	if *upgradeTo != "" {
		err = m.MigrateTo(*upgradeTo)
	} else if *downgradeTo != "" {
		err = m.RollbackTo(*downgradeTo)
	} else {
		err = m.Migrate()
	}

	if err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}