package postgres

import (
	"content_subsystem"
	"fmt"
	"testing"

	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupEnv() (content_subsystem.VideoRepository, func())  {
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
	return NewVideoRepository(client), func () {
		db, err := client.DB()
		if err != nil {
			print(err)
			return
		}
		db.Close()
	}
}

func TestStore(t *testing.T) {
	var(
		v1 = content_subsystem.Video{ID: content_subsystem.VideoID("A")}
		v2 = content_subsystem.Video{ID: content_subsystem.VideoID("B")}
	)

	r, close := SetupEnv()
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

func TestFind(t *testing.T) {
	r, close := SetupEnv()
	defer close()

	dbVideo1, err := r.Find(content_subsystem.VideoID("A"))
	if err != nil {
		t.Error(err)
	}

	if dbVideo1.ID != content_subsystem.VideoID("A") {
		t.Errorf("ID should be the same")
	}

	_, err = r.Find(content_subsystem.VideoID("C"))
	if err == nil {
		t.Errorf("Shouldn't find any user")
	}
}

func TestFindAll(t *testing.T) {
	r, close := SetupEnv()
	defer close()

	dbVideos := r.FindAll(nil)
	if len(dbVideos) != 2 {
		t.Errorf("There should be two videos")
	}

	dbVideos = r.FindAll(map[string]interface{}{
		"id": content_subsystem.VideoID("A"),
	})

	if len(dbVideos) != 1 {
		t.Errorf("There should be one video")
	}
}

func TestDelete(t *testing.T) {
	r, close := SetupEnv()
	defer close()

	err := r.Delete(content_subsystem.VideoID("A"))
	if err != nil {
		t.Error(err)
	}

	err = r.Delete(content_subsystem.VideoID("B"))
	if err != nil {
		t.Error(err)
	}

	dbVideos := r.FindAll(nil)
	if len(dbVideos) != 0 {
		t.Errorf("There should be no video")
	}
}