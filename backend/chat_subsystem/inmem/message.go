package inmem

import (
	"chat_subsystem"
	"sync"
)

type messageRepository struct {
	mtx    sync.RWMutex
	messages map[chat_subsystem.MessageID]*chat_subsystem.Message
}

func (r *messageRepository) Store(message *chat_subsystem.Message) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.messages[message.ID] = message
	return nil
}

func (r *messageRepository) Find(id chat_subsystem.MessageID) (*chat_subsystem.Message, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if val, ok := r.messages[id]; ok {
		return val, nil
	}
	return nil, chat_subsystem.ErrUnknownMessage
}

func (r *messageRepository) FindAll(conditions map[string]interface{}) []*chat_subsystem.Message {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	messages := []*chat_subsystem.Message{}
	if conditions == nil || len(conditions) == 0 {
		for _, val := range r.messages {
			messages = append(messages, val)
		}
	} else {
		for _, val := range r.messages {
			mapVal := val.ConvertToMap()
			for key, condVal := range conditions {
				if v, ok := mapVal[key]; ok && v == condVal {
					messages = append(messages, val)
				}
			}
		}
	}
	return messages
}

func (r *messageRepository) Delete(id chat_subsystem.MessageID) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.messages, id)
	return nil
}

// NewMessageRepository returns a new instance of a in-memory message repository.
func NewMessageRepository () chat_subsystem.MessageRepository {
	return &messageRepository{
		messages: make(map[chat_subsystem.MessageID]*chat_subsystem.Message),
	}
}