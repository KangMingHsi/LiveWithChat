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

	video := stream_subsystem.NewVideo(stream_subsystem.NextVideoID(), "", "", "", "", time.Minute)

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
	// var (
	// 	videoDB mockVideoRepository
	// 	s = NewService(&videoDB)
	// )

}

func TestUpdateVideo(t *testing.T) {

}

func TestDeleteVideo(t *testing.T) {
	
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