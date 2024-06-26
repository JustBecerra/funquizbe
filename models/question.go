package models

type QuestionText struct {
	Text string `json:"text"`
}

type Question struct {
	Category         string       `json:"category"`
	Id               string       `json:"id"`
	Tags             []string     `json:"tags"`
	Difficulty       string       `json:"difficulty"`
	Regions          []string     `json:"regions"`
	IsNiche          bool         `json:"isNiche"`
	Question         QuestionText `json:"question"`
	CorrectAnswer    string       `json:"correctAnswer"`
	IncorrectAnswers []string     `json:"incorrectAnswers"`
	Type             string       `json:"type"`
}