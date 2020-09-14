package questions

import (
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
	"context"
	"github.com/stretchr/testify/mock"
)

type QuestionRepositoryMock struct {
	mock.Mock
}

func (rm *QuestionRepositoryMock) CreateQuestion(ctx context.Context, q *entities.Question) (string, error) {
	args := rm.Called(ctx, q)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

func (rm *QuestionRepositoryMock) ReadQuestion(ctx context.Context, questionID string) (*entities.Question, error) {
	args := rm.Called(ctx, questionID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Question), args.Error(1)
}

func (rm *QuestionRepositoryMock) UpdateQuestion(ctx context.Context, questionID string, q *entities.Question) (string, error) {
	args := rm.Called(ctx, questionID, q)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

func (rm *QuestionRepositoryMock) DeleteQuestion(ctx context.Context, questionID string) (string, error) {
	args := rm.Called(ctx, questionID)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

func (rm *QuestionRepositoryMock) GetAll(ctx context.Context, limit int) (*entities.Questions, error) {
	args := rm.Called(ctx, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Questions), args.Error(1)
}

func (rm *QuestionRepositoryMock) CreateAnswer(ctx context.Context, questionID string, a *entities.Answer) (string, error) {
	args := rm.Called(ctx, questionID, a)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

func (rm *QuestionRepositoryMock) UpdateAnswer(ctx context.Context, questionID, answerID, newAnswer string) (string, error) {
	args := rm.Called(ctx, questionID, answerID, newAnswer)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}
