package content

import (
	"content_subsystem"
	"strconv"
	"testing"
)

var fakeID = "fake"

func TestSave(t *testing.T) {
	var (
		videoDB mockVideoRepository
		videoStorage mockVideoStorage
		videoScheduler mockVideoScheduler
		s = NewService(&videoDB, &videoStorage, &videoScheduler)
	)

	err := s.Save(fakeID, ".mp4", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestGetContentInfo(t *testing.T) {
	var (
		videoDB mockVideoRepository
		videoStorage mockVideoStorage
		videoScheduler mockVideoScheduler
		s = NewService(&videoDB, &videoStorage, &videoScheduler)
	)

	err := s.Save(fakeID, ".mp4", nil)
	if err != nil {
		t.Error(err)
	}

	info, err := s.GetContentInfo(fakeID)
	if err != nil {
		t.Error(err)
	}

	val, ok := info["duration"]
	if !ok {
		t.Error("there should be duration field")
	}

	_, err = strconv.ParseFloat(val.(string), 32)
	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	var (
		videoDB mockVideoRepository
		videoStorage mockVideoStorage
		videoScheduler mockVideoScheduler
		s = NewService(&videoDB, &videoStorage, &videoScheduler)
	)

	err := s.Save(fakeID, ".mp4", nil)
	if err != nil {
		t.Error(err)
	}

	err = s.Delete(fakeID)
	if err != nil {
		t.Error(err)
	}
}

type mockVideoRepository struct {
	video *content_subsystem.Video
}

func (r *mockVideoRepository) Store(video *content_subsystem.Video) error {
	r.video = video
	return nil
}

func (r *mockVideoRepository) Find(id content_subsystem.VideoID) (*content_subsystem.Video, error) {
	if r.video == nil {
		return nil, content_subsystem.ErrUnknownVideo
	}
	return r.video, nil
}

func (r *mockVideoRepository) FindAll(map[string]interface{}) []*content_subsystem.Video {
	if r.video == nil {
		return []*content_subsystem.Video{}
	}

	return []*content_subsystem.Video{
		r.video,
	}
}

func (r *mockVideoRepository) Delete(id content_subsystem.VideoID) error {
	r.video = nil
	return nil
}

type mockVideoStorage struct {}

func (s *mockVideoStorage) Save(id content_subsystem.VideoID, contentType string, content interface{}) error {
	return nil
}

func (s *mockVideoStorage) GetContentInfo(id content_subsystem.VideoID) (map[string]interface{}, error) {
	return map[string]interface{}{
		"duration": "1000000000",
	}, nil
}

func (s *mockVideoStorage) Delete(id content_subsystem.VideoID) error {
	return nil
}

type mockVideoScheduler struct {}

func (s *mockVideoScheduler) Submit(id string) {}

func (s *mockVideoScheduler) WorkerReady(w chan string) {}

func (s *mockVideoScheduler) Run() {}