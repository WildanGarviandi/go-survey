package controllers

import (
	"github.com/kataras/iris"
	"github.com/WildanGarviandi/survey-db/services"
)

func GetAllUsers(ctx iris.Context) {
	listOfUser, err := services.GetAllUser()

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map {
			"status": "Failed",
			"code": iris.StatusInternalServerError,
			"message": "Error",
			"data":   nil,
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map {
		"status": "OK",
		"code": iris.StatusOK,
		"message": "Success",
		"data":   listOfUser,
	})
}

func GetOneUser(ctx iris.Context) {
	uuid := ctx.URLParam("userId")

	user, err := services.GetOneUser(uuid)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map {
			"status": "Failed",
			"code": iris.StatusInternalServerError,
			"message": err.Error(),
			"data":   nil,
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map {
		"status": "OK",
		"code": iris.StatusOK,
		"message": "Success",
		"data":   user,
	})
}

func GetUserPosition(ctx iris.Context) {
	uuid := ctx.URLParam("userId")

	userPosition, err := services.GetUserPosition(uuid)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map {
			"status": "Failed",
			"code": iris.StatusInternalServerError,
			"message": err.Error(),
			"data":   nil,
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map {
		"status": "OK",
		"code": iris.StatusOK,
		"message": "Success",
		"data":   userPosition,
	})
}

func AddUser(ctx iris.Context) {
	type AddUserReq struct {
		UserUUID				string 			`json:"user_id"`
		Username				string			`json:"username"`
		Fullname				string			`json:"fullname"`
		PositionCode		string			`json:"position_code"`
		IsAbleSurvey		bool				`json:"is_able_survey"`
	}

	addUserReq := AddUserReq{}
	err := ctx.ReadJSON(&addUserReq)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map {
			"status": "Failed",
			"code": iris.StatusInternalServerError,
			"message": err.Error(),
			"data":   nil,
		})
		return
	}
	
}
