package inmem

import (
	"stream_subsystem"
	"sync"
)

type messageRepository struct {
	mtx    sync.RWMutex
	messages map[stream_subsystem.MessageID]*stream_subsystem.Message
}

func (r *messageRepository) Store(message *stream_subsystem.Message) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.messages[message.ID] = message
	return nil
}

func (r *messageRepository) Find(id stream_subsystem.MessageID) (*stream_subsystem.Message, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if val, ok := r.messages[id]; ok {
		return val, nil
	}
	return nil, stream_subsystem.ErrUnknownMessage
}

func (r *messageRepository) FindAll(conditions map[string]interface{}) []*stream_subsystem.Message {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	messages := []*stream_subsystem.Message{}
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

func (r *messageRepository) Delete(id stream_subsystem.MessageID) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.messages, id)
	return nil
}

// NewMessageRepository returns a new instance of a in-memory message repository.
func NewMessageRepository () stream_subsystem.MessageRepository {
	return &messageRepository{
		messages: make(map[stream_subsystem.MessageID]*stream_subsystem.Message),
	}
}