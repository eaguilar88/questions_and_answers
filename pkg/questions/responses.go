package questions

import (
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
)

//GetAllQuestionsResponse _
type GetAllQuestionsResponse struct {
	Questions entities.Questions `json:"questions"`
}

type ReverseResponse struct {
	Word string `json:"reversed_word"`
}
