package questions

import (
	"context"

	"github.com/go-kit/kit/log/level"

	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository contract
type Repository interface {
	CreateQuestion(ctx context.Context, q *entities.Question) (string, error)
	ReadQuestion(ctx context.Context, questionID int) (*entities.Question, error)
	UpdateQuestion(ctx context.Context, questionID int, q *entities.Question) (string, error)
	DeleteQuestion(ctx context.Context, questionID int) (string, error)
	GetAll(ctx context.Context) (entities.Questions, error)
}

type questionRepository struct {
	client *mongo.Client
	dbName string
	log    log.Logger
}

//NewRepository _
func NewRepository(client *mongo.Client, name string, logger log.Logger) Repository {
	return &questionRepository{client: client, dbName: name, log: logger}
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
func (r *questionRepository) GetAll(ctx context.Context) (entities.Questions, error) {
	coll := r.client.Database(r.dbName).Collection("questions")
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		level.Error(r.log).Log("msg", "error executing query", "error", err)
		return nil, err
	}
	var questions entities.Questions
	if err = cursor.All(ctx, &questions); err != nil {
		level.Error(r.log).Log("msg", "error decoding query results", "error", err)
		return nil, err
	}
	return questions, nil
}
