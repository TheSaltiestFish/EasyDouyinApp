package db

import (
	"context"
	"github.com/TheSaltiestFish/EasyDouyinApp/constants"
	"gorm.io/gorm"
)

type UserLogin struct {
	gorm.Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (n *UserLogin) TableName() string {
	return constants.UserLoginTableName
}

// CreateUserLogin create user_login info
func CreateUserLogin(ctx context.Context, userLogin *UserLogin) (int64, error) {
	if err := DB.Debug().WithContext(ctx).Create(&userLogin).Error; err != nil {
		return 0, err
	}
	return int64(userLogin.ID), nil
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

// // DeleteNote delete note info
//
//	func DeleteNote(ctx context.Context, noteID, userID int64) error {
//		return DB.WithContext(ctx).Where("id = ? and user_id = ? ", noteID, userID).Delete(&Note{}).Error
//	}

// QueryUserLogin query list of user info
func QueryUserLogin(ctx context.Context, username *string, password *string) (*UserLogin, error) {

	var res *UserLogin
	conn := DB.Debug().WithContext(ctx).Model(&UserLogin{}).Where("user_name = ?", username)
	conn = conn.Where("password = ?", password)

	if err := conn.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func QueryUserToken(ctx context.Context, username *string, token *string) (*UserLogin, error) {

	var res *UserLogin
	conn := DB.Debug().WithContext(ctx).Model(&UserLogin{}).Where("id = ?", username)

	conn = conn.Where("token = ?", token)

	if err := conn.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
