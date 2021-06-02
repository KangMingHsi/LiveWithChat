package stream

import (
	"mime/multipart"
	"strconv"
	"stream_subsystem"
	"time"
)

// Service is the interface that provides stream methods.
type Service interface {
	// Get all videos stored on server
	GetVideos() []*stream_subsystem.Video
	// Upload video to server
	UploadVideo(title, description, ownerID, videoType string, video *multipart.FileHeader) error
	// Update video information
	UpdateVideo(vid, uid string, data map[string]interface{}) error
	// Delete video from server
	DeleteVideo(vid, uid string) error
}

type service struct {
	videoDB stream_subsystem.VideoRepository
	contentController stream_subsystem.ContentController
}

func (s *service) GetVideos() []*stream_subsystem.Video {
	return s.videoDB.FindAll(nil)
}

func (s *service) UploadVideo(
	title, description, ownerID, videoType string,
	video *multipart.FileHeader,
) error {
	vid := stream_subsystem.NextVideoID()
	err := s.contentController.Save(string(vid), videoType, video)
	if err != nil {
		return err
	}

	information, err := s.contentController.GetContentInfo(string(vid))
	if err != nil {
		return err
	}

	duration, err := strconv.ParseFloat(information["duration"].(string), 64)
	if err != nil {
		return err
	}

	v := stream_subsystem.NewVideo(
		vid, title, description, ownerID, videoType, time.Duration(float64(time.Second) * duration))

	return s.videoDB.Store(v)
}

func (s *service) UpdateVideo(vid, uid string, data map[string]interface{}) error {
	v, err := s.videoDB.Find(stream_subsystem.VideoID(vid))
	if err != nil {
		return err
	}

	if v.OwnerID != uid {
		return stream_subsystem.ErrNoAuthority
	}

	newV := v.ConvertFromMap(data)
	return s.videoDB.Store(newV)
}

func (s *service) DeleteVideo(
	vid, uid string,
) error {
	v, err := s.videoDB.Find(stream_subsystem.VideoID(vid))
	if err != nil {
		return err
	}

	if v.OwnerID != uid {
		return stream_subsystem.ErrNoAuthority
	}

	err = s.contentController.Delete(vid)
	if err != nil {
		return err
	}
	return s.videoDB.Delete(stream_subsystem.VideoID(vid))
}

// NewService creates a stream service with necessary dependencies.
func NewService(
		videoDB stream_subsystem.VideoRepository,
		contentController stream_subsystem.ContentController) Service {
	return &service{
		videoDB: videoDB,
		contentController: contentController,
	}
}