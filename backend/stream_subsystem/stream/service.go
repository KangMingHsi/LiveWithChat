package stream

import "stream_subsystem"

// Service is the interface that provides stream methods.
type Service interface {
	// Get all videos stored on server
	GetVideos() []*stream_subsystem.Video
	// Upload video to server
	UploadVideo() error
	// Update video information
	UpdateVideo() error
	// Delete video from server
	DeleteVideo() error
}

type service struct {
	videoDB stream_subsystem.VideoRepository
}

func (s *service) GetVideos() []*stream_subsystem.Video {
	return nil
}

func (s *service) UploadVideo() error {
	return nil
}

func (s *service) UpdateVideo() error {
	return nil
}

func (s *service) DeleteVideo() error {
	return nil
}

// NewService creates a stream service with necessary dependencies.
func NewService(
		videoDB stream_subsystem.VideoRepository) Service {
	return &service{
		videoDB: videoDB,
	}
}