package postgres

import (
	"fmt"
	"stream_subsystem"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Message struct {
	ID int64		`gorm:"primaryKey"`
	VideoID string	`gorm:"not null;index"`
	OwnerID string  `gorm:"not null;index"`
	Text string 	`gorm:"not null"`
	CreatedAt time.Time `gorm:"<-:create"`
	IsValid bool    `gorm:"default:true"`
}

func toMessageDB(m *stream_subsystem.Message) *Message {
	return &Message{
		ID: int64(m.ID),
		VideoID: m.VideoID,
		Text: m.Text,
		CreatedAt: m.CreatedAt,
		OwnerID: m.OwnerID,
		IsValid: true,
	}
}

func toMessageModel(m *Message) *stream_subsystem.Message {
	return &stream_subsystem.Message{
		ID: stream_subsystem.MessageID(m.ID),
		VideoID: m.VideoID,
		Text: m.Text,
		CreatedAt: m.CreatedAt,
		OwnerID: m.OwnerID,
	}
}

type messageRepository struct {
	db *gorm.DB
}

func (r *messageRepository) Store(v *stream_subsystem.Message) error {
	result := r.db.Model(&Message{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(toMessageDB(v))

	return result.Error
}

func (r *messageRepository) Find(id stream_subsystem.MessageID) (*stream_subsystem.Message, error) {
	var vRow *Message

	result := r.db.Model(&Message{}).First(&vRow, "id = ? AND is_valid = ?", int64(id), true)
	if result.Error != nil {
		return nil, stream_subsystem.ErrUnknownMessage
	}

	return toMessageModel(vRow), nil
}

func (r *messageRepository) FindAll(conditions map[string]interface{}) []*stream_subsystem.Message {
	vRows := []*Message{}
	query := r.db.Model(&Message{}).Where("is_valid = ?", true)
	if conditions != nil && len(conditions) > 0 {
		for key, condVal := range conditions {
			query = query.Where(
				fmt.Sprintf("%s = ?", stream_subsystem.Underscore(key)),
				condVal,
			)
		}	
	}
	query.Find(&vRows)

	vs := make([]*stream_subsystem.Message, len(vRows))
	for index, vRow := range vRows {
		vs[index] = toMessageModel(vRow)
	}

	return vs
}

func (r *messageRepository) Delete(id stream_subsystem.MessageID) error {
	result := r.db.Model(
		&Message{}).Where(
			"id = ? AND is_valid = ?", int64(id), true).Update(
				"is_valid", false,
			)

	return result.Error
}

// NewMessageRepository returns a new instance of a postgres message repository.
func NewMessageRepository (client *gorm.DB) stream_subsystem.MessageRepository {
	r := &messageRepository{}
	r.db = client.Session(&gorm.Session{NewDB: true})
	return r
}