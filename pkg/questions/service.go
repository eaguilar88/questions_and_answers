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
	GetQuestionByID(ctx context.Context, questionID int) (*entities.Question, error)
	UpdateQuestion(ctx context.Context, questionID int, q *entities.Question) (string, error)
	DeleteQuestion(ctx context.Context, questionID int) (string, error)
	GetAll(ctx context.Context) ([]entities.Question, error)
}

type questionsService struct {
	repository Repository
	log        log.Logger
}

//NewService _
func NewService(repo Repository, logger log.Logger) Service {
	return &questionsService{repository: repo, log: logger}
}

//CreateQuestion _
func (svc *questionsService) CreateQuestion(ctx context.Context, q *entities.Question) (string, error) {
	return "", nil
}

//GetQuestionByID _
func (svc *questionsService) GetQuestionByID(ctx context.Context, questionID int) (*entities.Question, error) {
	q, err := svc.repository.ReadQuestion(ctx, questionID)
	if err != nil {
		level.Error(svc.log).Log("msg", "error getting question by id", "error", err)
		return nil, err
	}
	return q, nil
}

//UpdateQuestion _
func (svc *questionsService) UpdateQuestion(ctx context.Context, questionID int, q *entities.Question) (string, error) {
	return "", nil
}

//DeleteQuestion _
func (svc *questionsService) DeleteQuestion(ctx context.Context, questionID int) (string, error) {
	level.Info(svc.log).Log("msg", "delete not implemented")
	return "not implemented yet", nil
}

//GetAll _
func (svc *questionsService) GetAll(ctx context.Context) ([]entities.Question, error) {
	questions, err := svc.repository.GetAll(ctx)
	if err != nil {
		level.Error(svc.log).Log("msg", "error getting all questions", "error", err)
		return nil, err
	}
	return questions, nil
}
