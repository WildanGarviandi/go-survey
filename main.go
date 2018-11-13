package main

import (
	"fmt"
	"github.com/WildanGarviandi/survey-db/config"
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/iris-contrib/middleware/cors"
	"github.com/WildanGarviandi/survey-db/controllers"
)

func initApp() *iris.Application {
	api := iris.New()
	api.Use(logger.New())
	api.Use(cors.Default())
	api.Use(recover.New())
	api.Use(controllers.LogHandler)

	v1 := api.Party("/api/v1")
	{
		// get data
		v1.Get("/user/getAll", controllers.GetAllUsers)
		v1.Get("/user/getUser", controllers.GetOneUser)
		v1.Get("/user/getUserPosition", controllers.GetUserPosition)

		// insert data
		v1.Get("/user/addUser", controllers.AddUser)
	}
	return api
}

func main() {
	api := initApp()
	config.InitConfig()

	defer config.DbSurvey.Close()
	fmt.Println("Survey backend in running")
	irisConfiguration := iris.Configuration{
		DisableBodyConsumptionOnUnmarshal: true,
	}

	api.Run(iris.Addr(config.Port), iris.WithConfiguration(irisConfiguration))
}