package service

import (
	"gomicro/common"
	"gomicro/proto/user"
	"gomicro/user-service/model"
)

func RegisterUser(username string, password string, email string) user.UserLoginResponse {
	UserModel := model.User{
		Username: username,
		Password: password,
		Email:    email,
		Channel:  1,
	}
	message := "Success"
	token := ""

	db := common.GetDB()

	if err := db.Create(&UserModel).Error; err != nil {
		message = err.Error()

	} else {
		token = "Test Token"
	}

	return user.UserLoginResponse{
		Message: message,
		Token:   token,
	}
}

func UserLogin(username string, password string) user.UserLoginResponse {
	db := common.GetDB()
	userModel := &model.User{}
	//存在绑定在model里面
	err := db.Where("username=? and password=?", username, password).Find(userModel).Error
	message := "Success"
	token := ""
	if err != nil {
		message = "Failed"
	} else {
		token = common.GenToken(userModel.Username, userModel.ID)
	}
	return user.UserLoginResponse{
		Message: message,
		Token:   token,
	}
}
