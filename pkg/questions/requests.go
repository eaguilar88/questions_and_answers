package questions

import "bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"

type GetAllQuestionsRequest struct {
	ItemsPerPage int
}

type GetQuestionRequest struct {
	QuestionID string
}
type CreateQuestionRequest struct {
	Title       string
	Description string
}

type UpdateQuestionRequest struct {
	ID          string
	Title       string
	Description string
	Answers     []*entities.Answer
}

type DeleteQuestionRequest struct {
	QuestionID string
}

type AddAnswerRequest struct {
	QuestionID string
	Answer     string
}

type UpdateAnswerRequest struct {
	QuestionID string
	AnswerID   string
	Answer     string
}
