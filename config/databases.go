package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func loadMysqlDB(databases string, hostAddress string) (db *gorm.DB, err error) {
	con, err := gorm.Open("mysql", hostAddress)

	if err != nil {
		fmt.Println("error on opening DB on port = " + hostAddress)
		return con, err
	}

	con.LogMode(LogMode)
	con.SingularTable(true)
	con.DB().SetMaxIdleConns(10)

	if Env != "testing" {
		fmt.Println("Connected to mysql db = " + databases)
	}

	return con, err
}

func InitDb(c Config) {
	conString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", c.MySql[0].Username, c.MySql[0].Password, c.MySql[0].Db)
	if c.MySql[0].Password != "" {
		conString += " password=" + c.MySql[0].Password
	}

	var err error
	
	DbSurvey, err = loadMysqlDB(c.MySql[0].Db, conString)
	if err != nil {
		panic(err)
	}
}
