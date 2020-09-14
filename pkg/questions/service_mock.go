package questions

import (
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
	"context"
	"github.com/stretchr/testify/mock"
)

type QuestionServiceMock struct {
	mock.Mock
}

//CreateQuestion _
func (sm *QuestionServiceMock) CreateQuestion(ctx context.Context, q *entities.Question) (string, error) {
	args := sm.Called(ctx, q)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

//GetQuestionByID _
func (sm *QuestionServiceMock) GetQuestionByID(ctx context.Context, questionID string) (*entities.Question, error) {
	args := sm.Called(ctx, questionID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Question), args.Error(1)
}

//UpdateQuestion _
func (sm *QuestionServiceMock) UpdateQuestion(ctx context.Context, questionID string, q *entities.Question) (string, error) {
	args := sm.Called(ctx, questionID, q)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

//DeleteQuestion _
func (sm *QuestionServiceMock) DeleteQuestion(ctx context.Context, questionID string) (string, error) {
	args := sm.Called(ctx, questionID)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

//GetAll _
func (sm *QuestionServiceMock) GetAll(ctx context.Context, limit int) (*entities.Questions, error) {
	args := sm.Called(ctx, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Questions), args.Error(1)
}

//AddAnswer _
func (sm *QuestionServiceMock) AddAnswer(ctx context.Context, questionID string, a *entities.Answer) (string, error) {
	args := sm.Called(ctx, questionID, a)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

//UpdateAnswer _
func (sm *QuestionServiceMock) UpdateAnswer(ctx context.Context, questionID, answerID string, a *entities.Answer) (string, error) {
	args := sm.Called(ctx, questionID, answerID, a)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}
