package model

import "time"

type UserSurvey struct {
	UserUUID				string 			`gorm:"column:user_uuid" json:"user_id"`
	Username				string			`gorm:"column:username" json:"username"`
	Fullname				string			`gorm:"column:fullname" json:"fullname"`
	PositionCode		string			`gorm:"column:position_code" json:"position_code"`
	IsAbleSurvey		bool				`gorm:"column:is_able_survey" json:"is_able_survey"`
	CreatedDate			time.Time		`gorm:"column:createdAt" json:"createdAt"`
	UpdatedDate			time.Time		`gorm:"column:updatedAt" json:"updatedAt"`
}

type UserPosition struct {
	Fullname					string		`gorm:"fullname" json:"fullname"`
	PositionDesc			string		`gorm:"position_desc" json:"position_desc"`
	IsEligbleSurvey		bool			`gorm:"is_able_survey" json:"is_able_survey"`
}
