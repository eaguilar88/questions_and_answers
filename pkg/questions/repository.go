package questions

import (
	"context"
	"database/sql"

	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
	"github.com/go-kit/kit/log"
)

//Repository contract
type Repository interface {
	CreateQuestion(ctx context.Context, q *entities.Question) (string, error)
	ReadQuestion(ctx context.Context, questionID int) (*entities.Question, error)
	UpdateQuestion(ctx context.Context, questionID int, q *entities.Question) (string, error)
	DeleteQuestion(ctx context.Context, questionID int) (string, error)
	GetAll(ctx context.Context) ([]entities.Question, error)
}

type questionRepository struct {
	db  *sql.DB
	log log.Logger
}

//NewRepository _
func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &questionRepository{db: db, log: logger}
}

func (r *questionRepository) CreateQuestion(ctx context.Context, q *entities.Question) (string, error) {
	return "", nil
}
func (r *questionRepository) ReadQuestion(ctx context.Context, questionID int) (*entities.Question, error) {
	return nil, nil
}
func (r *questionRepository) UpdateQuestion(ctx context.Context, questionID int, q *entities.Question) (string, error) {
	return "", nil
}
func (r *questionRepository) DeleteQuestion(ctx context.Context, questionID int) (string, error) {
	return "", nil
}
func (r *questionRepository) GetAll(ctx context.Context) ([]entities.Question, error) {
	return []entities.Question{}, nil
}
