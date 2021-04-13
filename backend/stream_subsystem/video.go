package stream_subsystem

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// VideoID uniquely identifies a particular video.
type VideoID string


// Video is the central class in domain model
type Video struct {
	ID VideoID
	Title string
	Description string
	Duration   time.Duration
	Type    string
	Likes   int
	Dislikes int
	OwnerID string
}

func (v Video) ConvertToMap() map[string]interface{} {
	return map[string]interface{}{
		"ID": v.ID,
		"Title": v.Title,
		"Description": v.Description,
		"Duration": v.Duration,
		"Likes": v.Likes,
		"Dislikes": v.Dislikes,
		"OwnerID": v.OwnerID,
		"Type": v.Type,
	}
}

// NewVideo creates a video instance
func NewVideo() *Video {
	return &Video{

	}
}


// VideoRepository provides access to a video store
type VideoRepository interface {
	Store(video *Video) error
	Find(id VideoID) (*Video, error)
	FindAll(map[string]interface{}) []*Video
	Delete(id VideoID) error
}

// ErrUnknownVideo is used when a video could not be found.
var ErrUnknownVideo = errors.New("unknown video")

// NextVideoID generates a new video ID.
// TODO: Move to infrastructure(?)
func NextVideoID() VideoID {
	return VideoID(strings.ToUpper(uuid.New().String()))
}

