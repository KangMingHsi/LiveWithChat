package postgres

import (
	"auth_subsystem"
	"fmt"
	"testing"

	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var client *gorm.DB
func SetupEnv() (auth_subsystem.UserRepository, func())  {
	dsn := fmt.Sprint(
		"host=database_test user=livewithchat password=default",
		" dbname=mock port=5432 sslmode=disable TimeZone=Asia/Taipei",
	)
	client, err := gorm.Open(psql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return NewUserRepository(client), func () {
		db, err := client.DB()
		if err != nil {
			print(err)
			return
		}
		db.Close()
	}
}
 
func TestStore(t *testing.T) {
	var (
		user1 = &auth_subsystem.User{
			ID: auth_subsystem.MemberID("1"),
			Email: "a@a.com",
		}
		user2 = &auth_subsystem.User{
			ID: auth_subsystem.MemberID("2"),
			Email: "b@b.com",
		}
	)

	r, close := SetupEnv()
	defer close()
	err := r.Store(user1)
	if err != nil {
		t.Error(err)
	}

	err = r.Store(user2)
	if err != nil {
		t.Error(err)
	}
}

func TestFind(t *testing.T) {
	r, close := SetupEnv()
	defer close()

	dbUser1, err := r.Find("a@a.com")
	if err != nil {
		t.Error(err)
	}

	if dbUser1.Email != "a@a.com" {
		t.Errorf("Email should be the same")
	}

	_, err = r.Find("a@b.com")
	if err == nil {
		t.Errorf("Shouldn't find any user")
	}
}

func TestFindAll(t *testing.T) {
	r, close := SetupEnv()
	defer close()

	dbUsers := r.FindAll()
	if len(dbUsers) != 2 {
		t.Errorf("There should be two users")
	}
}