package stream_subsystem

import (
	"errors"
	"reflect"
	"strings"
	"time"
)

type MessageID int64

type Message struct {
	ID MessageID
	VideoID string
	Text string
	OwnerID string
	CreatedAt time.Time
}

// Convert Message type to map
func (m Message) ConvertToMap() map[string]interface{} {
	return map[string]interface{}{
		"ID": m.ID,
		"OwnerID": m.OwnerID,
		"Text": m.Text,
		"VideoID": m.VideoID,
		"CreatedAt": m.CreatedAt,
	}
}

// Create Message from map
func (m *Message) ConvertFromMap(data map[string]interface{}) *Message {
	if data != nil && len(data) > 0 {
		val := reflect.ValueOf(m).Elem()
		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)
			typeField := val.Type().Field(i)
			if v, ok := data[strings.ToLower(typeField.Name)]; ok {
				valueField.Set(reflect.ValueOf(v))
			}
		}
	}
	return m
}


type MessageFactory struct {
	base int64
}

// NewMessage creates a message instance
func (f *MessageFactory) NewMessage(
	text, videoID, ownerID string,
) *Message {
	f.base = f.base + 1
	return &Message{
		ID: MessageID(f.base),
		Text: text,
		OwnerID: ownerID,
		VideoID: videoID,
		CreatedAt: time.Now().UTC(),
	}
}

// NewMessageFactory creates a message factory instance
func NewMessageFactory(base int64) *MessageFactory {
	return &MessageFactory{
		base: base - 1,
	}
}

// MessageRepository provides access to a message store
type MessageRepository interface {
	Store(message *Message) error
	Find(id MessageID) (*Message, error)
	FindAll(map[string]interface{}) []*Message
	Delete(id MessageID) error
}

// ErrUnknownMessage is used when a message could not be found.
var ErrUnknownMessage = errors.New("unknown message")