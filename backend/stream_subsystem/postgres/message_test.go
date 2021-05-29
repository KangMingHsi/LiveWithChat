package postgres

import (
	"fmt"
	"stream_subsystem"
	"testing"

	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupMessageEnv() (stream_subsystem.MessageRepository, func())  {
	dsn := fmt.Sprint(
		"host=database_test user=livewithchat password=default",
		" dbname=mock port=5432 sslmode=disable TimeZone=Asia/Taipei",
	)
	var client *gorm.DB
	client, err := gorm.Open(psql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return NewMessageRepository(client), func () {
		db, err := client.DB()
		if err != nil {
			print(err)
			return
		}
		db.Close()
	}
}

func TestStoreMessage(t *testing.T) {
	var(
		v1 = stream_subsystem.Message{ID: stream_subsystem.MessageID(0)}
		v2 = stream_subsystem.Message{ID: stream_subsystem.MessageID(1)}
	)

	r, close := SetupMessageEnv()
	defer close()
	err := r.Store(&v1)
	if err != nil {
		t.Error(err)
	}

	err = r.Store(&v2)
	if err != nil {
		t.Error(err)
	}
}

func TestFindMessage(t *testing.T) {
	r, close := SetupMessageEnv()
	defer close()

	dbVideo1, err := r.Find(stream_subsystem.MessageID(0))
	if err != nil {
		t.Error(err)
	}

	if dbVideo1.ID != stream_subsystem.MessageID(0) {
		t.Errorf("ID should be the same")
	}

	_, err = r.Find(stream_subsystem.MessageID(2))
	if err == nil {
		t.Errorf("Shouldn't find any user")
	}
}

func TestFindAllMessages(t *testing.T) {
	r, close := SetupMessageEnv()
	defer close()

	dbMessages := r.FindAll(nil)
	if len(dbMessages) != 2 {
		t.Errorf("There should be two messages")
	}

	dbMessages = r.FindAll(map[string]interface{}{
		"id": stream_subsystem.MessageID(0),
	})

	if len(dbMessages) != 1 {
		t.Errorf("There should be one message")
	}
}

func TestDeleteMessage(t *testing.T) {
	r, close := SetupMessageEnv()
	defer close()

	err := r.Delete(stream_subsystem.MessageID(0))
	if err != nil {
		t.Error(err)
	}

	err = r.Delete(stream_subsystem.MessageID(1))
	if err != nil {
		t.Error(err)
	}

	dbMessages := r.FindAll(nil)
	if len(dbMessages) != 0 {
		t.Errorf("There should be no message")
	}
}