package questions

import (
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
)

//GetAllQuestionsResponse _
type GetAllQuestionsResponse struct {
	Questions *entities.Questions `json:"questions"`
}

type GetQuestionResponse struct {
	Question *entities.Question `json:"question"`
}

type CreateResourceResponse struct {
	CreatedID string `json:"created_id"`
}

type HandleResourceResponse struct {
	Message string `json:"message"`
}
