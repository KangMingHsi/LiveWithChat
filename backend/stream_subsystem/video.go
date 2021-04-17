package stream_subsystem

import (
	"errors"
	"reflect"
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

func (v *Video) ConvertFromMap(data map[string]interface{}) *Video {
	if data != nil && len(data) > 0 {
		val := reflect.ValueOf(v).Elem()
		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)
			typeField := val.Type().Field(i)
			if v, ok := data[strings.ToLower(typeField.Name)]; ok {
				valueField.Set(reflect.ValueOf(v))
			}
		}
	}
	return v
}

// NewVideo creates a video instance
func NewVideo(
		id VideoID,
		title, description, ownerID, videoType string,
		duration time.Duration) *Video {
	return &Video{
		ID: id,
		Title: title,
		Description: description,
		Duration: duration,
		Likes: 0,
		Dislikes: 0,
		OwnerID: ownerID,
		Type: videoType,
	}
}

// NextVideoID generates a new video ID.
// TODO: Move to infrastructure(?)
func NextVideoID() VideoID {
	return VideoID(strings.ToUpper(uuid.New().String()))
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
