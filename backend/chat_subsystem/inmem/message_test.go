package inmem

import (
	"chat_subsystem"
	"testing"
)

func TestStoreMessage(t *testing.T) {
	var(
		v1 = chat_subsystem.Message{ID: chat_subsystem.MessageID(0)}
		v2 = chat_subsystem.Message{ID: chat_subsystem.MessageID(1)}
	)

	r := NewMessageRepository()
	err := r.Store(&v1)
	if err != nil {
		t.Error(err)
	}

	err = r.Store(&v2)
	if err != nil {
		t.Error(err)
	}
}

func TestFindMessage(t *testing.T) {
	var(
		v1 = chat_subsystem.Message{ID: chat_subsystem.MessageID(0)}
		v2 = chat_subsystem.Message{ID: chat_subsystem.MessageID(1)}
	)

	r := NewMessageRepository()
	r.Store(&v1)
	r.Store(&v2)

	dbVideo1, err := r.Find(chat_subsystem.MessageID(0))
	if err != nil {
		t.Error(err)
	}

	if dbVideo1.ID != chat_subsystem.MessageID(0) {
		t.Errorf("ID should be the same")
	}

	_, err = r.Find(chat_subsystem.MessageID(2))
	if err == nil {
		t.Errorf("Shouldn't find any user")
	}
}


func TestFindAllMessages(t *testing.T) {
	var(
		v1 = chat_subsystem.Message{ID: chat_subsystem.MessageID(0)}
		v2 = chat_subsystem.Message{ID: chat_subsystem.MessageID(1)}
	)

	r := NewMessageRepository()
	r.Store(&v1)
	r.Store(&v2)

	dbVideos := r.FindAll(nil)
	if len(dbVideos) != 2 {
		t.Errorf("There should be two videos")
	}

	dbVideos = r.FindAll(map[string]interface{}{
		"ID": chat_subsystem.MessageID(0),
	})

	if len(dbVideos) != 1 {
		t.Errorf("There should be one Message")
	}
}

func TestDeleteMessage(t *testing.T) {
	var(
		v1 = chat_subsystem.Message{ID: chat_subsystem.MessageID(0)}
		v2 = chat_subsystem.Message{ID: chat_subsystem.MessageID(1)}
	)

	r := NewMessageRepository()
	r.Store(&v1)
	r.Store(&v2)

	r.Delete(chat_subsystem.MessageID(0))
	dbVideos := r.FindAll(nil)
	if len(dbVideos) != 1 {
		t.Errorf("There should be one video")
	}
}