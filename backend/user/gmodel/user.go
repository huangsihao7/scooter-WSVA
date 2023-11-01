package gmodel

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type User struct {
	Id            uint64       `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name          string       `gorm:"column:name;type:varchar(255);comment:用户姓名;NOT NULL" json:"name"`
	Gender        uint         `gorm:"column:gender;type:tinyint(3) unsigned;default:0;comment:用户性别;NOT NULL" json:"gender"`
	Mobile        string       `gorm:"column:mobile;type:varchar(255);comment:用户电话;NOT NULL" json:"mobile"`
	Password      string       `gorm:"column:password;type:varchar(255);comment:用户密码;NOT NULL" json:"password"`
	Dec           string       `gorm:"column:dec;type:varchar(255);comment:个性签名;NOT NULL" json:"dec"`
	Avatar        string       `gorm:"column:avatar;type:varchar(255);comment:头像;NOT NULL" json:"avatar"`
	BackgroundUrl string       `gorm:"column:background_url;type:varchar(255);comment:背景图;NOT NULL" json:"background_url"`
	CreateTime    sql.NullTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime    sql.NullTime `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"update_time"`
}

func (m *User) TableName() string {
	return "user"
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

// FindOneByMobile 根据手机号判断用户是否存在
func (m *UserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	var user User
	result := m.db.WithContext(ctx).Where("mobile = ?", mobile).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// InsertUser 新增一条用户记录
func (m *UserModel) InsertUser(ctx context.Context, user *User) error {
	result := m.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByID 根据用户id获取用户信息
func (m *UserModel) GetUserByID(ctx context.Context, userID uint) (*User, error) {
	var user User
	result := m.db.WithContext(ctx).Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (m *UserModel) UpdateUser(ctx context.Context, user *User) error {
	// Save the updated user information in the database
	existingUser := &User{}
	err := m.db.WithContext(ctx).First(existingUser, user.Id).Error
	if err != nil {
		return err
	}
	existingUser.Gender = user.Gender
	existingUser.Name = user.Name
	existingUser.Avatar = user.Avatar
	existingUser.Dec = user.Dec
	existingUser.BackgroundUrl = user.BackgroundUrl

	// 保存更新后的用户信息到数据库
	err = m.db.WithContext(ctx).Save(existingUser).Error
	if err != nil {
		return err
	}

	return nil
}
