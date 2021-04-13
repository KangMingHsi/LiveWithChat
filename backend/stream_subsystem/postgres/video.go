package postgres

import (
	"fmt"
	"stream_subsystem"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Video struct {
	ID string		`gorm:"primaryKey"`
	Title string	`gorm:"not null"`
	Description string `gorm:"default:''"`
	Duration   time.Duration `gorm:"not null"`
	Type    string `gorm:"not null"`
	Likes   int    `gorm:"default:0"`
	Dislikes int   `gorm:"default:0"`
	OwnerID string  `gorm:"not null;index"`
	IsValid bool    `gorm:"default:true"`
}

func toVideoDB(v *stream_subsystem.Video) *Video {
	return &Video{
		ID: string(v.ID),
		Title: v.Title,
		Description: v.Description,
		Duration: v.Duration,
		Type: v.Type,
		Likes: v.Likes,
		Dislikes: v.Dislikes,
		OwnerID: v.OwnerID,
		IsValid: true,
	}
}

func toVideoModel(v *Video) *stream_subsystem.Video {
	return &stream_subsystem.Video{
		ID: stream_subsystem.VideoID(v.ID),
		Title: v.Title,
		Description: v.Description,
		Duration: v.Duration,
		Type: v.Type,
		Likes: v.Likes,
		Dislikes: v.Dislikes,
		OwnerID: v.OwnerID,
	}
}

type videoRepository struct {
	db *gorm.DB
}

func (r *videoRepository) Store(v *stream_subsystem.Video) error {
	result := r.db.Model(&Video{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(toVideoDB(v))

	return result.Error
}

func (r *videoRepository) Find(id stream_subsystem.VideoID) (*stream_subsystem.Video, error) {
	var vRow *Video

	result := r.db.Model(&Video{}).First(&vRow, "id = ? AND is_valid = ?", string(id), true)
	if result.Error != nil {
		return nil, stream_subsystem.ErrUnknownVideo
	}

	return toVideoModel(vRow), nil
}

func (r *videoRepository) FindAll(conditions map[string]interface{}) []*stream_subsystem.Video {
	vRows := []*Video{}
	query := r.db.Model(&Video{}).Where("is_valid = ?", true)
	if conditions != nil && len(conditions) > 0 {
		for key, condVal := range conditions {
			query = query.Where(
				fmt.Sprintf("%s = ?", stream_subsystem.Underscore(key)),
				condVal,
			)
		}	
	}
	query.Find(&vRows)
	print(len(vRows))
	vs := make([]*stream_subsystem.Video, len(vRows))
	for index, vRow := range vRows {
		vs[index] = toVideoModel(vRow)
	}

	return vs
}

func (r *videoRepository) Delete(id stream_subsystem.VideoID) error {
	result := r.db.Model(
		&Video{}).Where(
			"id = ? AND is_valid = ?", string(id), true).Update(
				"is_valid", false,
			)

	return result.Error
}

// NewVideoRepository returns a new instance of a postgres video repository.
func NewVideoRepository (client *gorm.DB) stream_subsystem.VideoRepository {
	r := &videoRepository{}
	r.db = client.Session(&gorm.Session{NewDB: true})
	return r
}