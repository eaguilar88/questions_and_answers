package questions

import "bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"

type CreateQuestionRequest struct {
	Question entities.Question `json:"question"`
}
