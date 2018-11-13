package services

import (
	"fmt"
	"github.com/WildanGarviandi/survey-db/config"
	"github.com/WildanGarviandi/survey-db/model"
)

func GetAllUser() ([]model.UserSurvey, error) {

	listOfUser := []model.UserSurvey{}

	query := `select * from user_survey`

	fmt.Println("query : " + query)
	err := config.DbSurvey.Raw(query).Find(&listOfUser).Error

	return listOfUser, err
}

func GetOneUser(uuid string) (model.UserSurvey, error) {
	user := model.UserSurvey{}

	query := `select * from user_survey where user_uuid = ?`

	fmt.Println("query : " + query + uuid)
	err := config.DbSurvey.Raw(query, uuid).Find(&user).Error

	return user, err
}

func GetUserPosition(uuid string) (model.UserPosition, error) {
	userPosition := model.UserPosition{}

	query := `select user_survey.fullname, user_position.position_desc, user_survey.is_able_survey 
		from user_survey 
		left join user_position on user_position.position_code = user_survey.position_code
		WHERE user_survey.user_uuid = ?`
		
	fmt.Println("query : " + query)

	err := config.DbSurvey.Raw(query, uuid).Find(&userPosition).Error

	return userPosition, err
}

func AddUser() () {
	// query := `INSERT INTO 'user_survey'('user_uuid', 'username', 'fullname', 'position_code', 'is_able_survey', 'createdAt', 'updatedAt') VALUES (?,?,?,?,?,?,?)`
}
