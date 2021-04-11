package postgres

import (
	"auth_subsystem"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// User schema
type User struct {
	ID  		auth_subsystem.MemberID `gorm:"primaryKey"`

	Email string `gorm:"unique;index"`
	Nickname string
	Gender            string
	HashedPassword    string
	Role		      string
	IsOnline bool `gorm:"default:false"`
	IsBlocked bool `gorm:"default:false"`
	IpAddr  pq.StringArray `gorm:"type:text[]"`

	LimitationPeriod time.Time
	LoginTime time.Time
}

func toUserDB(u *auth_subsystem.User) *User {
	return &User{
		ID: u.ID,

		Email: u.Email,
		Gender: u.Gender,
		Nickname: u.Nickname,
		HashedPassword: u.HashedPassword,
		Role: u.Role,
		IsOnline: u.IsOnline,
		IsBlocked: u.IsBlocked,
		IpAddr: u.IpAddr,

		LimitationPeriod: u.LimitationPeriod,
		LoginTime: u.LoginTime,
	}
}

func toUser(u *User) *auth_subsystem.User {
	return &auth_subsystem.User{
		ID: u.ID,

		Email: u.Email,
		Gender: u.Gender,
		Nickname: u.Nickname,
		HashedPassword: u.HashedPassword,
		Role: u.Role,
		IsOnline: u.IsOnline,
		IsBlocked: u.IsBlocked,
		IpAddr: []string(u.IpAddr),

		LimitationPeriod: u.LimitationPeriod,
		LoginTime: u.LoginTime,
	}
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Store(u *auth_subsystem.User) error {
	result := r.db.Model(&User{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(toUserDB(u))
	
	return result.Error
}

func (r *userRepository) Find(email string) (*auth_subsystem.User, error) {
	var uRow *User

	result := r.db.Model(&User{}).First(&uRow, "email = ?", email)
	if result.Error != nil {
		return nil, auth_subsystem.ErrUnknownUser
	}

	return toUser(uRow), nil
}

func (r *userRepository) FindAll() []*auth_subsystem.User {
	uRows := []*User{}
	r.db.Model(&User{}).Find(&uRows)

	us := make([]*auth_subsystem.User, len(uRows))
	for index, uRow := range uRows {
		us[index] = toUser(uRow)
	}

	return us
}

// NewUserRepository returns a new instance of a postgres user repository.
func NewUserRepository (client *gorm.DB) auth_subsystem.UserRepository {
	r := &userRepository{}
	r.db = client.Session(&gorm.Session{NewDB: true})
	return r
}