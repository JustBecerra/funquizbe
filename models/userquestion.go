package models

type UserQuestion struct {
	Id               string   `json:"id"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correctAnswer"`
	IncorrectAnswers []string `json:"incorrectAnswers"`
}