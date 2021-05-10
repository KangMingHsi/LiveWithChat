package postgres

import (
	"content_subsystem"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Video struct {
	ID string		`gorm:"primaryKey"`
	IsDirty	bool	`gorm:"default:true"`
	IsValid bool    `gorm:"default:true"`
}

func toVideoDB(v *content_subsystem.Video) *Video {
	return &Video{
		ID: string(v.ID),
		IsDirty: v.IsDirty,
		IsValid: true,
	}
}

func toVideoModel(v *Video) *content_subsystem.Video {
	return &content_subsystem.Video{
		ID: content_subsystem.VideoID(v.ID),
		IsDirty: v.IsDirty,
	}
}

type videoRepository struct {
	db *gorm.DB
}

func (r *videoRepository) Store(v *content_subsystem.Video) error {
	result := r.db.Model(&Video{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(toVideoDB(v))

	return result.Error
}

func (r *videoRepository) Find(id content_subsystem.VideoID) (*content_subsystem.Video, error) {
	var vRow *Video

	result := r.db.Model(&Video{}).First(&vRow, "id = ? AND is_valid = ?", string(id), true)
	if result.Error != nil {
		return nil, content_subsystem.ErrUnknownVideo
	}

	return toVideoModel(vRow), nil
}

func (r *videoRepository) FindAll(conditions map[string]interface{}) []*content_subsystem.Video {
	vRows := []*Video{}
	query := r.db.Model(&Video{}).Where("is_valid = ?", true)
	if conditions != nil && len(conditions) > 0 {
		for key, condVal := range conditions {
			query = query.Where(
				fmt.Sprintf("%s = ?", content_subsystem.Underscore(key)),
				condVal,
			)
		}	
	}
	query.Find(&vRows)

	vs := make([]*content_subsystem.Video, len(vRows))
	for index, vRow := range vRows {
		vs[index] = toVideoModel(vRow)
	}

	return vs
}

func (r *videoRepository) Delete(id content_subsystem.VideoID) error {
	result := r.db.Model(
		&Video{}).Where(
			"id = ? AND is_valid = ?", string(id), true).Update(
				"is_valid", false,
			)

	return result.Error
}

// NewVideoRepository returns a new instance of a postgres video repository.
func NewVideoRepository (client *gorm.DB) content_subsystem.VideoRepository {
	r := &videoRepository{}
	r.db = client.Session(&gorm.Session{NewDB: true})
	return r
}