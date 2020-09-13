package answers

import (
	"context"
	"database/sql"

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
	db  *sql.DB
	log log.Logger
}

//NewRepository _
func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &answerRepository{db: db, log: logger}
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
