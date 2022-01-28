package controllers

import (
	"github.com/gin-gonic/gin"
	"mvc/constants"
	"mvc/requests"
	"mvc/services"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var params requests.CreateUserParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": constants.FailValidation,
			"msg": err.Error(),
			"data": nil,
		})
		return
	}

	if err = services.CreateUser(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": constants.Error,
			"msg": err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": constants.SuccessCreate,
		"msg": "success",
		"data": nil,
	})
}

func GetUserList (c *gin.Context) {
	var params requests.GetUserListParam
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": constants.FailValidation,
			"msg": err.Error(),
			"data": nil,
		})
		return
	}

	userList, err := services.GetUserList(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}  else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": userList,
		})
	}
}

func GetUserDetail (c *gin.Context) {
	user, err := services.GetUserById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": constants.Success,
		"msg": "success",
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {
	var params requests.UpdateUserParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": constants.FailValidation,
			"msg": err.Error(),
			"data": nil,
		})
		return
	}

	err = services.UpdateUser(params, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": constants.FailUpdate,
			"msg": "FAIL",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": constants.SuccessUpdate,
		"msg": "OK",
		"data": nil,
	})
}

func DeleteUser(c *gin.Context) {
	err := services.DeleteUserById(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"code": constants.FailDelete,
			"msg": err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": constants.SuccessDelete,
		"msg": "OK",
		"data": nil,
	})
}
