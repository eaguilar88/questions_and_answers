package questions

import (
	"context"

	"github.com/go-kit/kit/log/level"

	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
	"github.com/go-kit/kit/log"
)

//Service interface
type Service interface {
	CreateQuestion(ctx context.Context, q *entities.Question) (string, error)
	GetQuestionByID(ctx context.Context, questionID string) (*entities.Question, error)
	UpdateQuestion(ctx context.Context, questionID string, q *entities.Question) (string, error)
	DeleteQuestion(ctx context.Context, questionID string) (string, error)
	GetAll(ctx context.Context, limit int) (*entities.Questions, error)
	AddAnswer(ctx context.Context, questionID string, a *entities.Answer) (string, error)
	UpdateAnswer(ctx context.Context, questionID, answerID string, a *entities.Answer) (string, error)
}

type questionService struct {
	repository Repository
	log        log.Logger
}

//NewService _
func NewService(repo Repository, logger log.Logger) *questionService {
	return &questionService{repository: repo, log: logger}
}

//CreateQuestion _
func (svc *questionService) CreateQuestion(ctx context.Context, q *entities.Question) (string, error) {
	created, err := svc.repository.CreateQuestion(ctx, q)
	if err != nil {
		level.Error(svc.log).Log("msg", "error creating question", "error", err)
		return "", err
	}
	return created, nil
}

//GetQuestionByID _
func (svc *questionService) GetQuestionByID(ctx context.Context, questionID string) (*entities.Question, error) {
	q, err := svc.repository.ReadQuestion(ctx, questionID)
	if err != nil {
		level.Error(svc.log).Log("msg", "error getting question by id", "error", err)
		return nil, err
	}
	return q, nil
}

//UpdateQuestion _
func (svc *questionService) UpdateQuestion(ctx context.Context, questionID string, q *entities.Question) (string, error) {
	questionForUpdate, err := svc.repository.ReadQuestion(ctx, questionID)
	if err != nil {
		level.Error(svc.log).Log("msg", "error getting question by id", "error", err)
		return "", err
	}
	if q.Title != "" {
		questionForUpdate.Title = q.Title
	}
	if q.Description != "" {
		questionForUpdate.Description = q.Description
	}
	updatedId, err := svc.repository.UpdateQuestion(ctx, questionID, questionForUpdate)
	if err != nil {
		level.Error(svc.log).Log("msg", "error updating question by id", "error", err)
		return "", err
	}
	return updatedId, nil
}

//DeleteQuestion _
func (svc *questionService) DeleteQuestion(ctx context.Context, questionID string) (string, error) {
	result, err := svc.repository.DeleteQuestion(ctx, questionID)
	if err != nil {
		level.Error(svc.log).Log("msg", "error getting question by id", "error", err)
		return "", err
	}
	return result, nil
}

//GetAll _
func (svc *questionService) GetAll(ctx context.Context, limit int) (*entities.Questions, error) {
	questions, err := svc.repository.GetAll(ctx, limit)
	if err != nil {
		level.Error(svc.log).Log("msg", "error getting all questions", "error", err)
		return nil, err
	}
	return questions, nil
}

func (svc *questionService) AddAnswer(ctx context.Context, questionID string, a *entities.Answer) (string, error) {
	result, err := svc.repository.CreateAnswer(ctx, questionID, a)
	if err != nil {
		level.Error(svc.log).Log("msg", "error adding answer to question", "error", err)
		return "", err
	}
	return result, nil
}

func (svc *questionService) UpdateAnswer(ctx context.Context, questionID, answerID string, a *entities.Answer) (string, error) {
	result, err := svc.repository.UpdateAnswer(ctx, questionID, answerID, a.Answer)
	if err != nil {
		level.Error(svc.log).Log("msg", "error updating answer", "error", err)
		return "", err
	}
	return result, nil
}
