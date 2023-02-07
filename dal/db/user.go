package db

import (
	"context"
	"github.com/TheSaltiestFish/EasyDouyinApp/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID        int64  `json:"user_id"`
	UserName      string `json:"user_name"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
}

func (n *User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(ctx context.Context, user *User) error {
	if err := DB.Debug().WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	var res []*User
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.Debug().WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

//// UpdateUser update user info
//func UpdateUser(ctx context.Context, noteID, userID int64, title, content *string) error {
//	params := map[string]interface{}{}
//	if title != nil {
//		params["title"] = *title
//	}
//	if content != nil {
//		params["content"] = *content
//	}
//	return DB.WithContext(ctx).Model(&Note{}).Where("id = ? and user_id = ?", noteID, userID).
//		Updates(params).Error
//}

//// DeleteNote delete note info
//func DeleteNote(ctx context.Context, noteID, userID int64) error {
//	return DB.WithContext(ctx).Where("id = ? and user_id = ? ", noteID, userID).Delete(&Note{}).Error
//}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userid *string) (*User, error) {

	var res *User
	conn := DB.Debug().WithContext(ctx).Model(&User{}).Where("user_id = ?", userid)

	if err := conn.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
