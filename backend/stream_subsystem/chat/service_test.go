package chat

import (
	"testing"

	"stream_subsystem"
)

var (
	messageDB mockMessageRepository
)

func TestCreateMessage(t *testing.T) {
	var (
		service = NewService(&messageDB)
		id = stream_subsystem.MessageID(0)
		text = "hello"
		videoID = "0"
		ownerID = "0"
	)

	err := service.CreateMessage(text, videoID, ownerID)
	if err != nil {
		t.Error(err.Error())
	}

	msg, err := messageDB.Find(id)
	if err != nil {
		t.Error(err.Error())
	}

	if msg.ID != id || msg.Text != text ||
		msg.VideoID != videoID || msg.OwnerID != ownerID {
		
		t.Errorf("Some fields are wrong %v", msg)
	}
}

func TestUpdateMessage(t *testing.T) {
	var (
		service = NewService(&messageDB)
		id = stream_subsystem.MessageID(0)
		text = "hello"
		nextText = "world"
		videoID = "0"
		ownerID = "0"
	)

	msg, err := messageDB.Find(id)
	if err != nil {
		t.Error(err.Error())
	}

	if msg.ID != id || msg.Text != text ||
		msg.VideoID != videoID || msg.OwnerID != ownerID {
		
		t.Errorf("Some fields are wrong %v", msg)
	}

	err = service.UpdateMessage(int64(id), nextText, ownerID)
	if err != nil {
		t.Error(err.Error())
	}

	msg, err = messageDB.Find(id)
	if err != nil {
		t.Error(err.Error())
	}

	if msg.ID != id || msg.Text != nextText ||
		msg.VideoID != videoID || msg.OwnerID != ownerID {
		
		t.Errorf("Some fields are wrong %v", msg)
	}
}

func TestListMessage(t *testing.T) {
	var (
		service = NewService(&messageDB)
		id = stream_subsystem.MessageID(0)
		nextText = "world"
		videoID = "0"
		ownerID = "0"
	)

	messages := service.GetMessages(videoID)
	if len(messages) != 1 {
		t.Errorf("It should be one message")
	}

	msg := messages[0]
	if msg.ID != id || msg.Text != nextText ||
		msg.VideoID != videoID || msg.OwnerID != ownerID {
		
		t.Errorf("Some fields are wrong %v", msg)
	}
}

func TestDeleteMessage(t *testing.T) {
	var (
		service = NewService(&messageDB)
		id = stream_subsystem.MessageID(0)
		ownerID = "0"
	)

	err := service.DeleteMessage(int64(id), ownerID)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = messageDB.Find(id)
	if err == nil {
		t.Errorf("It should be no message")
	}
}

type mockMessageRepository struct {
	message *stream_subsystem.Message
}

func (r *mockMessageRepository) Store(message *stream_subsystem.Message) error {
	r.message = message
	return nil
}

func (r *mockMessageRepository) Find(id stream_subsystem.MessageID) (*stream_subsystem.Message, error) {
	if r.message == nil {
		return nil, stream_subsystem.ErrUnknownMessage
	}
	return r.message, nil
}

func (r *mockMessageRepository) FindAll(map[string]interface{}) []*stream_subsystem.Message {
	if r.message == nil {
		return []*stream_subsystem.Message{}
	}

	return []*stream_subsystem.Message{
		r.message,
	}
}

func (r *mockMessageRepository) Delete(id stream_subsystem.MessageID) error {
	r.message = nil
	return nil
}
