package questions

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-kit/kit/log/level"

	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository contract
type Repository interface {
	CreateQuestion(ctx context.Context, q *entities.Question) (string, error)
	ReadQuestion(ctx context.Context, questionID string) (*entities.Question, error)
	UpdateQuestion(ctx context.Context, questionID string, q *entities.Question) (string, error)
	DeleteQuestion(ctx context.Context, questionID string) (string, error)
	GetAll(ctx context.Context, limit int) (*entities.Questions, error)
	CreateAnswer(ctx context.Context, questionID string, a *entities.Answer) (string, error)
	UpdateAnswer(ctx context.Context, questionID, answerID, newAnswer string) (string, error)
}

type questionRepository struct {
	client *mongo.Client
	dbName string
	log    log.Logger
}

const collectionName = "questions"

//NewRepository _
func NewRepository(client *mongo.Client, name string, logger log.Logger) Repository {
	return &questionRepository{client: client, dbName: name, log: logger}
}

func (r *questionRepository) CreateQuestion(ctx context.Context, q *entities.Question) (string, error) {
	q.ID = primitive.NewObjectID()
	result, err := r.client.Database(r.dbName).Collection(collectionName).InsertOne(ctx, &q)
	if err != nil {
		level.Error(r.log).Log("message", "error inserting question", "error", err)
		return "", errors.New("there was an error inserting question")
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *questionRepository) ReadQuestion(ctx context.Context, questionID string) (*entities.Question, error) {
	objID, err := primitive.ObjectIDFromHex(questionID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"_id": objID,
	}
	result := r.client.Database(r.dbName).Collection(collectionName).FindOne(ctx, filter)
	if result.Err() != nil {
		level.Error(r.log).Log("message", "error running query", "error", result.Err())
		return nil, errors.New("there was an error running the query")
	}
	var question entities.Question
	err = result.Decode(&question)
	if err != nil {
		level.Error(r.log).Log("message", "error decoding query", "error", err)
		return nil, errors.New("there was an error decoding the query")
	}
	return &question, nil
}

func (r *questionRepository) UpdateQuestion(ctx context.Context, questionID string, q *entities.Question) (string, error) {
	objID, err := primitive.ObjectIDFromHex(questionID)
	if err != nil {
		return "", err
	}
	filter := bson.M{
		"_id": objID,
	}
	update := bson.D{{
		"$set",
		q,
	}}
	result := r.client.Database(r.dbName).Collection(collectionName).FindOneAndUpdate(ctx, filter, update)
	if err := result.Err(); err != nil {
		level.Error(r.log).Log("message", "error updating question", "error", err)
		return "", errors.New("there was an error updating question")
	}
	return "", nil
}

func (r *questionRepository) DeleteQuestion(ctx context.Context, questionID string) (string, error) {
	objID, err := primitive.ObjectIDFromHex(questionID)
	if err != nil {
		return "error", err
	}
	filter := bson.M{
		"_id": objID,
	}

	_, err = r.client.Database(r.dbName).Collection(collectionName).DeleteOne(ctx, filter)
	if err != nil {
		level.Error(r.log).Log("message", "error delete question", "error", err)
		return "", errors.New("there was an error delete question")
	}
	return "", nil
}

func (r *questionRepository) GetAll(ctx context.Context, limit int) (*entities.Questions, error) {
	coll := r.client.Database(r.dbName).Collection(collectionName)
	var opts = &options.FindOptions{}
	if limit != 0 {
		opts.SetLimit(int64(limit))
	}
	cursor, err := coll.Find(ctx, bson.D{}, opts)
	if err != nil {
		level.Error(r.log).Log("msg", "error executing query", "error", err)
		return nil, err
	}
	var questions entities.Questions
	if err = cursor.All(ctx, &questions); err != nil {
		level.Error(r.log).Log("msg", "error decoding query results", "error", err)
		return nil, err
	}
	return &questions, nil
}

func (r *questionRepository) CreateAnswer(ctx context.Context, questionID string, a *entities.Answer) (string, error) {
	objID, err := primitive.ObjectIDFromHex(questionID)
	if err != nil {
		return "", err
	}
	filter := bson.M{
		"_id": objID,
	}
	a.ID = primitive.NewObjectID()
	update := bson.M{"$addToSet": bson.M{"answers": a}}
	result := r.client.Database(r.dbName).Collection(collectionName).FindOneAndUpdate(ctx, filter, update)
	if err := result.Err(); err != nil {
		level.Error(r.log).Log("message", "error adding answer to question", "error", err)
		return "", errors.New("there was an error adding answer to question")
	}
	return a.ID.Hex(), nil
}

func (r *questionRepository) UpdateAnswer(ctx context.Context, questionID, answerID, newAnswer string) (string, error) {
	objID, err := primitive.ObjectIDFromHex(questionID)
	if err != nil {
		return "", err
	}
	objAnsID, err := primitive.ObjectIDFromHex(answerID)
	if err != nil {
		return "", err
	}
	filter := bson.M{
		"_id":         objID,
		"answers._id": objAnsID,
	}
	update := bson.D{{
		"$set",
		bson.D{
			{"answers.$.answer", newAnswer},
		},
	}}
	result := r.client.Database(r.dbName).Collection(collectionName).FindOneAndUpdate(ctx, filter, update)
	if err := result.Err(); err != nil {
		level.Error(r.log).Log("message", "error updating answer of question", "error", err)
		return "", errors.New("there was an error updating answer of question")
	}
	return "success", nil
}
