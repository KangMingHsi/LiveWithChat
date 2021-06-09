package chat

import (
	"chat_subsystem"
)

// Service is the interface that provides chat methods.
type Service interface {
	// Get all messages stored on server
	GetMessages(vid string) []*chat_subsystem.Message
	// Create message to server
	CreateMessage(text, videoID, ownerID string) error
	// Update message information
	UpdateMessage(id int64, text, uid string) error
	// Delete message from server
	DeleteMessage(id int64, uid string) error
}

type service struct {
	messageDB chat_subsystem.MessageRepository
	factory *chat_subsystem.MessageFactory
}

func (s *service) GetMessages(vid string) []*chat_subsystem.Message {
	return s.messageDB.FindAll(map[string]interface{}{
		"video_id": vid,
	})
}

func (s *service) CreateMessage(text, videoID, ownerID string) error {
	msg := s.factory.NewMessage(text, videoID, ownerID)

	return s.messageDB.Store(msg)
}

func (s *service) UpdateMessage(id int64, text, uid string) error {
	msg, err := s.messageDB.Find(chat_subsystem.MessageID(id))
	if err != nil {
		return err
	}

	if msg.OwnerID != uid {
		return chat_subsystem.ErrNoAuthorityMessage
	}

	msg.Text = text
	return s.messageDB.Store(msg)
}

func (s *service) DeleteMessage(id int64, uid string) error {
	msg, err := s.messageDB.Find(chat_subsystem.MessageID(id))
	if err != nil {
		return err
	}

	if msg.OwnerID != uid {
		return chat_subsystem.ErrNoAuthorityMessage
	}

	return s.messageDB.Delete(chat_subsystem.MessageID(id))
}

func NewService(
	messageDB chat_subsystem.MessageRepository,
) Service{
	messages := messageDB.FindAll(nil)
	return &service{
		messageDB: messageDB,
		factory: chat_subsystem.NewMessageFactory(int64(len(messages))),
	}
}