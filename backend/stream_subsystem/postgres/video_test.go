package postgres

import (
	"fmt"
	"stream_subsystem"
	"testing"

	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupVideoEnv() (stream_subsystem.VideoRepository, func())  {
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

func TestStoreVideo(t *testing.T) {
	var(
		v1 = stream_subsystem.Video{ID: stream_subsystem.VideoID("A")}
		v2 = stream_subsystem.Video{ID: stream_subsystem.VideoID("B")}
	)

	r, close := SetupVideoEnv()
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

func TestFindVideo(t *testing.T) {
	r, close := SetupVideoEnv()
	defer close()

	dbVideo1, err := r.Find(stream_subsystem.VideoID("A"))
	if err != nil {
		t.Error(err)
	}

	if dbVideo1.ID != stream_subsystem.VideoID("A") {
		t.Errorf("ID should be the same")
	}

	_, err = r.Find(stream_subsystem.VideoID("C"))
	if err == nil {
		t.Errorf("Shouldn't find any user")
	}
}

func TestFindAllVideos(t *testing.T) {
	r, close := SetupVideoEnv()
	defer close()

	dbVideos := r.FindAll(nil)
	if len(dbVideos) != 2 {
		t.Errorf("There should be two videos")
	}

	dbVideos = r.FindAll(map[string]interface{}{
		"id": stream_subsystem.VideoID("A"),
	})

	if len(dbVideos) != 1 {
		t.Errorf("There should be one video")
	}
}

func TestDeleteVideo(t *testing.T) {
	r, close := SetupVideoEnv()
	defer close()

	err := r.Delete(stream_subsystem.VideoID("A"))
	if err != nil {
		t.Error(err)
	}

	err = r.Delete(stream_subsystem.VideoID("B"))
	if err != nil {
		t.Error(err)
	}

	dbVideos := r.FindAll(nil)
	if len(dbVideos) != 0 {
		t.Errorf("There should be no video")
	}
}