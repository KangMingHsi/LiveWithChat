package postgres

import (
	"authentication"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type user struct {
	ID  		authentication.MemberID `gorm:"primaryKey"`

	Email string `gorm:"unique;index"`
	Nickname string
	Gender            string
	HashedPassword    string
	Role		      string
	IsOnline bool `gorm:"default:false"`
	IsBlocked bool `gorm:"default:false"`
	IpAddr  pq.StringArray `gorm:"type:text[];default:{}"`

	LimitationPeriod time.Time
	LoginTime time.Time
}

func toUserDB(u *authentication.User) *user {
	return &user{
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

func toUser(u *user) *authentication.User {
	return &authentication.User{
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

func (r *userRepository) Store(u *authentication.User) error {
	r.db.Model(&user{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(toUserDB(u))
	return nil
}

func (r *userRepository) Find(email string) (*authentication.User, error) {
	var uRow *user

	result := r.db.Model(&user{}).First(&uRow, "email = ?", email)
	if result.Error != nil {
		return nil, authentication.ErrUnknownUser
	}

	return toUser(uRow), nil
}

func (r *userRepository) FindAll() []*authentication.User {
	uRows := []*user{}
	r.db.Model(&user{}).Find(&uRows)

	us := make([]*authentication.User, len(uRows))
	for index, uRow := range uRows {
		us[index] = toUser(uRow)
	}

	return us
}

// NewUserRepository returns a new instance of a postgres user repository.
func NewUserRepository (client *gorm.DB) authentication.UserRepository {
	r := &userRepository{}
	client.Migrator().AutoMigrate(&user{})
	r.db = client.Session(&gorm.Session{NewDB: true})
	return r
}