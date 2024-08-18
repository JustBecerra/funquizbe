package models

type UserQuestion struct {
	ID               int      `json:"id" gorm:"primaryKey"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correctanswer"`
	IncorrectAnswers []string `json:"incorrectAnswers"`
}