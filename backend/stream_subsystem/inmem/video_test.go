package inmem

import (
	"stream_subsystem"
	"testing"
)

func TestStoreVideo(t *testing.T) {
	var(
		v1 = stream_subsystem.Video{ID: stream_subsystem.VideoID("A")}
		v2 = stream_subsystem.Video{ID: stream_subsystem.VideoID("B")}
	)

	r := NewVideoRepository()
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
	var(
		v1 = stream_subsystem.Video{ID: stream_subsystem.VideoID("A")}
		v2 = stream_subsystem.Video{ID: stream_subsystem.VideoID("B")}
	)

	r := NewVideoRepository()
	r.Store(&v1)
	r.Store(&v2)

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
	var(
		v1 = stream_subsystem.Video{ID: stream_subsystem.VideoID("A")}
		v2 = stream_subsystem.Video{ID: stream_subsystem.VideoID("B")}
	)

	r := NewVideoRepository()
	r.Store(&v1)
	r.Store(&v2)

	dbVideos := r.FindAll(nil)
	if len(dbVideos) != 2 {
		t.Errorf("There should be two videos")
	}

	dbVideos = r.FindAll(map[string]interface{}{
		"ID": stream_subsystem.VideoID("A"),
	})

	if len(dbVideos) != 1 {
		t.Errorf("There should be one video")
	}
}

func TestDeleteVideo(t *testing.T) {
	var(
		v1 = stream_subsystem.Video{ID: stream_subsystem.VideoID("A")}
		v2 = stream_subsystem.Video{ID: stream_subsystem.VideoID("B")}
	)

	r := NewVideoRepository()
	r.Store(&v1)
	r.Store(&v2)

	r.Delete(stream_subsystem.VideoID("A"))
	dbVideos := r.FindAll(nil)
	if len(dbVideos) != 1 {
		t.Errorf("There should be one video")
	}
}