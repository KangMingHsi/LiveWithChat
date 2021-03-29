package postgres

import (
	"authentication"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type user struct {
	ID  		authentication.MemberID `gorm:"primaryKey"`
	HashedPassword    string
	IsOnline bool `gorm:"default:false"`
	IsBlocked bool `gorm:"default:false"`
	IpAddr  string
	LastLoginTime time.Time
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Store(u *authentication.User) error {
	r.db.Model(&user{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(u.ConvertToMap())
	return nil
}

func (r *userRepository) Find(id authentication.MemberID) (*authentication.User, error) {
	var u *authentication.User
	result := r.db.Model(&user{}).First(&u, "id = ?", id)
	if result.Error != nil {
		return nil, authentication.ErrUnknownUser
	}
	return u, nil
}

func (r *userRepository) FindAll() []*authentication.User {
	us := []*authentication.User{}
	r.db.Model(&user{}).Find(&us)
	return us
}

// NewUserRepository returns a new instance of a postgres user repository.
func NewUserRepository (client *gorm.DB) authentication.UserRepository {
	r := &userRepository{}
	client.Migrator().AutoMigrate(&user{})
	r.db = client.Session(&gorm.Session{NewDB: true})
	return r
}