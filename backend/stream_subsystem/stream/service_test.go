package stream

import (
	"stream_subsystem"
	"testing"
	"time"
)

func TestGetVideos(t *testing.T) {
	var (
		videoDB mockVideoRepository
		s = NewService(&videoDB, nil)
	)

	videos := s.GetVideos()
	if len(videos) > 0 {
		t.Errorf("There should be no any video")
	}

	video := stream_subsystem.NewVideo(
		stream_subsystem.NextVideoID(), "", "", "", "", time.Minute)

	err := videoDB.Store(video)
	if err != nil {
		t.Error(err)
	}

	videos = s.GetVideos()
	if len(videos) != 1 {
		t.Errorf("There should be one video")
	}

	err = videoDB.Delete(video.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestUploadVideos(t *testing.T) {
	var (
		ownerID = "test123"
		videoDB mockVideoRepository
		contentController mockContentController
		s = NewService(&videoDB, &contentController)
	)

	err := s.UploadVideo(
		"title", "description", ownerID, ".mp4", nil)
	if err != nil {
		t.Error(err)
	}

	videos := s.GetVideos()
	if len(videos) != 1 {
		t.Errorf("There should be one video")
	}
}

func TestUpdateVideo(t *testing.T) {
	var (
		videoID = stream_subsystem.NextVideoID()
		ownerID = "test123"
		videoDB mockVideoRepository
		contentController mockContentController
		s = NewService(&videoDB, &contentController)
	)

	video := stream_subsystem.NewVideo(
		videoID, "", "", ownerID, "", time.Minute)

	err := videoDB.Store(video)
	if err != nil {
		t.Error(err)
	}

	err = s.UpdateVideo(string(videoID), ownerID, map[string]interface{}{
		"title": "newTitle",
	})
	if err != nil {
		t.Error(err)
	}

	v, err := videoDB.Find(videoID)
	if v.Title != "newTitle" {
		t.Errorf("the title should be newTitle")
	}
}

func TestDeleteVideo(t *testing.T) {
	var (
		videoID = stream_subsystem.NextVideoID()
		ownerID = "test123"
		videoDB mockVideoRepository
		contentController mockContentController
		s = NewService(&videoDB, &contentController)
	)

	video := stream_subsystem.NewVideo(
		videoID, "", "", ownerID, "", time.Minute)

	err := videoDB.Store(video)
	if err != nil {
		t.Error(err)
	}

	err = s.DeleteVideo(string(videoID), ownerID)
	if err != nil {
		t.Error(err)
	}

	videos := s.GetVideos()
	if len(videos) > 0 {
		t.Errorf("There should be no any video")
	}
}

type mockVideoRepository struct {
	video *stream_subsystem.Video
}

func (r *mockVideoRepository) Store(video *stream_subsystem.Video) error {
	r.video = video
	return nil
}

func (r *mockVideoRepository) Find(id stream_subsystem.VideoID) (*stream_subsystem.Video, error) {
	if r.video == nil {
		return nil, stream_subsystem.ErrUnknownVideo
	}
	return r.video, nil
}

func (r *mockVideoRepository) FindAll(map[string]interface{}) []*stream_subsystem.Video {
	if r.video == nil {
		return []*stream_subsystem.Video{}
	}

	return []*stream_subsystem.Video{
		r.video,
	}
}

func (r *mockVideoRepository) Delete(id stream_subsystem.VideoID) error {
	r.video = nil
	return nil
}

type mockContentController struct {}

func (c *mockContentController) Save(id, contentType string, content interface{}) error {
	return nil
}

func (c *mockContentController) GetContentInfo(id string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"duration": "1000000000",
	}, nil
}

func (c *mockContentController) Delete(id string) error {
	return nil
}