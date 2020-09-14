package answers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
	"github.com/go-kit/kit/log"
)

//Repository contract
type Repository interface {
	CreateAnswer(ctx context.Context, q *entities.Answer) (string, error)
	ReadAnswer(ctx context.Context, answerID int) (*entities.Answer, error)
	UpdateAnswer(ctx context.Context, answerID int, q *entities.Answer) (string, error)
	DeleteAnswer(ctx context.Context, answerID int) (string, error)
	GetAll(ctx context.Context) ([]entities.Answer, error)
}

type answerRepository struct {
	client *mongo.Client
	dbName string
	log    log.Logger
}

//NewRepository _
func NewRepository(client *mongo.Client, name string, logger log.Logger) Repository {
	return &answerRepository{client: client, dbName: name, log: logger}
}

func (r *answerRepository) CreateAnswer(ctx context.Context, q *entities.Answer) (string, error) {
	return "", nil
}
func (r *answerRepository) ReadAnswer(ctx context.Context, answerID int) (*entities.Answer, error) {
	return nil, nil
}
func (r *answerRepository) UpdateAnswer(ctx context.Context, answerID int, q *entities.Answer) (string, error) {
	return "", nil
}
func (r *answerRepository) DeleteAnswer(ctx context.Context, answerID int) (string, error) {
	return "", nil
}
func (r *answerRepository) GetAll(ctx context.Context) ([]entities.Answer, error) {
	return []entities.Answer{}, nil
}
