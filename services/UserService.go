package services

import (
	"mvc/databases"
	"mvc/models"
	"mvc/requests"
	"mvc/utils"
)

func CreateUser(params requests.CreateUserParam) (err error)  {
	user := models.User{
		Username: params.Username,
		Password: utils.EncryptPassword(params.Password),
	}
	err = databases.MySqlClient.Create(&user).Error
	return
}

func GetUserList(params requests.GetUserListParam) (userList []*models.User, err error) {
	tx := databases.MySqlClient.Select("id", "username", "status")
	if params.Username != "" {
		tx.Where("username LIKE ?", params.Username + "%")
	}
	err = tx.Find(&userList).Error
	return
}

func GetUserById(id string) (user *models.User, err error)  {
	err = databases.MySqlClient.Select("id", "username", "status").First(&user, id).Error
	return
}

func UpdateUser(params requests.UpdateUserParam, id string) (err error)  {
	user := models.User{
		Username: params.Username,
		Password: utils.EncryptPassword(params.Password),
		Status: params.Status,
	}
	err = databases.MySqlClient.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
	return
}

func DeleteUserById(id string) (err error) {
	err = databases.MySqlClient.Delete(&models.User{}, id).Error
	return
}
