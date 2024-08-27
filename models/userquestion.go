package models

type UserQuestion struct {
	ID               int      `json:"id" gorm:"primaryKey"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correctAnswer"`
	IncorrectAnswers []string `json:"incorrectAnswers"`
}