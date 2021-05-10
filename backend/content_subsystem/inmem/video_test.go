package inmem

import (
	"content_subsystem"
	"testing"
)

func TestStore(t *testing.T) {
	var(
		v1 = content_subsystem.Video{ID: content_subsystem.VideoID("A")}
		v2 = content_subsystem.Video{ID: content_subsystem.VideoID("B")}
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

func TestFind(t *testing.T) {
	var(
		v1 = content_subsystem.Video{ID: content_subsystem.VideoID("A")}
		v2 = content_subsystem.Video{ID: content_subsystem.VideoID("B")}
	)

	r := NewVideoRepository()
	r.Store(&v1)
	r.Store(&v2)

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
	var(
		v1 = content_subsystem.Video{ID: content_subsystem.VideoID("A")}
		v2 = content_subsystem.Video{ID: content_subsystem.VideoID("B")}
	)

	r := NewVideoRepository()
	r.Store(&v1)
	r.Store(&v2)

	dbVideos := r.FindAll(nil)
	if len(dbVideos) != 2 {
		t.Errorf("There should be two videos")
	}

	dbVideos = r.FindAll(map[string]interface{}{
		"ID": content_subsystem.VideoID("A"),
	})

	if len(dbVideos) != 1 {
		t.Errorf("There should be one video")
	}
}

func TestDelete(t *testing.T) {
	var(
		v1 = content_subsystem.Video{ID: content_subsystem.VideoID("A")}
		v2 = content_subsystem.Video{ID: content_subsystem.VideoID("B")}
	)

	r := NewVideoRepository()
	r.Store(&v1)
	r.Store(&v2)

	r.Delete(content_subsystem.VideoID("A"))
	dbVideos := r.FindAll(nil)
	if len(dbVideos) != 1 {
		t.Errorf("There should be one video")
	}
}