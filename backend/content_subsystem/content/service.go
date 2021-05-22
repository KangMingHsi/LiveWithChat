package content

import (
	"content_subsystem"
)

// Service is the interface that provides stream methods.
type Service interface {
	// Save content with the id and type
	Save(id, contentType string, content interface{}) error
	// Get content information
	GetContentInfo(id string) (map[string]interface{}, error)
	// Delete content by id
	Delete(id string) error
}

type service struct {
	videoDB content_subsystem.VideoRepository
	videoStorage content_subsystem.VideoStorage
	videoScheduler content_subsystem.VideoScheduler
}

func (s *service) Save(id, contentType string, content interface{}) error {
	vid := content_subsystem.VideoID(id)
	err := s.videoStorage.Save(vid, contentType, content)
	if err != nil {
		return err
	}

	v := content_subsystem.NewVideo(id)
	err = s.videoDB.Store(v)
	if err != nil {
		s.videoStorage.Delete(vid)
	}

	s.videoScheduler.Submit(id)

	return err
}

func (s *service) GetContentInfo(id string) (map[string]interface{}, error) {
	vid := content_subsystem.VideoID(id)
	_, err := s.videoDB.Find(vid)
	if err != nil {
		return nil, err
	}

	return s.videoStorage.GetContentInfo(content_subsystem.VideoID(id))
}

func (s *service) Delete(id string) error {
	vid := content_subsystem.VideoID(id)
	err := s.videoDB.Delete(vid)
	if err != nil {
		return err
	}

	err = s.videoStorage.Delete(vid)
	if err != nil {
		return err
	}

	return nil
}

// NewService creates a content service with necessary dependencies.
func NewService(
	videoDB content_subsystem.VideoRepository,
	videoStorage content_subsystem.VideoStorage,
	videoScheduler content_subsystem.VideoScheduler,
) Service {
	return &service{
		videoDB: videoDB,
		videoStorage: videoStorage,
		videoScheduler: videoScheduler,
	}
}