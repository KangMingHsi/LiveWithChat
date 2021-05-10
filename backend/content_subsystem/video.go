package content_subsystem

import (
	"errors"
)

// VideoID uniquely identifies a particular video.
type VideoID string

// Video is the central class in domain model
type Video struct {
	ID VideoID
	IsDirty bool
}

func (v Video) ConvertToMap() map[string]interface{} {
	return map[string]interface{}{
		"ID": v.ID,
		"IsDirty": v.IsDirty,
	}
}

// NewVideo returns a new instance of a video model.
func NewVideo(id string) *Video {
	return &Video{
		ID: VideoID(id),
		IsDirty: true,
	}
}

// VideoRepository provides access to a video store
type VideoRepository interface {
	Store(video *Video) error
	Find(id VideoID) (*Video, error)
	FindAll(map[string]interface{}) []*Video
	Delete(id VideoID) error
}

// VideoStorage provides access to a video files store
type VideoStorage interface {
	Save(id VideoID, contentType string, content interface{}) error
	Delete(id VideoID) error
	GetContentInfo(id VideoID) (map[string]interface{}, error)
}

// ErrUnknownVideo is used when a video could not be found.
var ErrUnknownVideo = errors.New("unknown video")

// ErrNoAuthority is used when do something to a video that you cannot.
var ErrNoAuthority = errors.New("no authority")