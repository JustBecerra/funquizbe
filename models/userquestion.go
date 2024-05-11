package models

type UserQuestion struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Question      string `json:"question"`
	CorrectAnswer string `json:"correctanswer"`
	WrongAnswer1  string `json:"wronganswer1"`
	WrongAnswer2  string `json:"wronganswer2"`
	WrongAnswer3  string `json:"wronganswer3"`
}