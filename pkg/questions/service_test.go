package questions

import (
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/testhelp"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"testing"
)

type serviceSuite struct {
	ctx    context.Context
	repo   *QuestionRepositoryMock
	logger log.Logger
	svc    *questionService
	suite.Suite
}

func (s *serviceSuite) SetupTest() {
	s.ctx = context.Background()
	s.repo = new(QuestionRepositoryMock)
	s.logger = log.NewNopLogger()
	s.svc = NewService(s.repo, s.logger)
}

func TestServiceSuite(t *testing.T) {
	s := new(serviceSuite)
	suite.Run(t, s)
}

func (s *serviceSuite) TestGetAll_Success() {
	var questions = testhelp.MockQuestionArray()
	limit := rand.Intn(4)
	s.repo.On("GetAll", s.ctx, limit).Return(questions, nil)
	qs, err := s.svc.GetAll(s.ctx, limit)
	assert.Nil(s.T(), err)
	assert.IsType(s.T(), entities.Questions{}, *qs)
}

func (s *serviceSuite) TestGetQuestion_Success() {
	questionID := primitive.NewObjectID()
	var question = &entities.Question{
		ID:          questionID,
		Title:       "Generic question?",
		Description: "Generic description",
		Answers:     nil,
	}
	s.repo.On("ReadQuestion", s.ctx, questionID.Hex()).Return(question, nil)
	q, err := s.svc.GetQuestionByID(s.ctx, questionID.Hex())
	assert.Nil(s.T(), err)
	assert.IsType(s.T(), entities.Question{}, *q)
	assert.Equal(s.T(), questionID, q.ID)
}

func (s *serviceSuite) TestCreateQuestion_Success() {
	questionID := primitive.NewObjectID()
	var question = &entities.Question{
		Title:       "Generic question?",
		Description: "Generic description",
	}
	s.repo.On("CreateQuestion", s.ctx, question).Return(questionID.Hex(), nil)
	_, err := s.svc.CreateQuestion(s.ctx, question)
	assert.Nil(s.T(), err)
}

func (s *serviceSuite) TestUpdateQuestion_Success() {
	questionID := primitive.NewObjectID()
	var question = &entities.Question{
		Title:       "Generic question?",
		Description: "Generic description",
	}
	s.repo.On("ReadQuestion", s.ctx, questionID.Hex()).Return(question, nil)
	s.repo.On("UpdateQuestion", s.ctx, questionID.Hex(), question).Return("success", nil)
	result, err := s.svc.UpdateQuestion(s.ctx, questionID.Hex(), question)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), result, "success")
}

func (s *serviceSuite) TestDeleteQuestion_Success() {
	questionID := primitive.NewObjectID()
	s.repo.On("DeleteQuestion", s.ctx, questionID.Hex()).Return("success", nil)
	result, err := s.svc.DeleteQuestion(s.ctx, questionID.Hex())
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), result, "success")
}

func (s *serviceSuite) TestCreateAnswer_Success() {
	questionID := primitive.NewObjectID().Hex()
	answerID := primitive.NewObjectID().Hex()
	var answer = &entities.Answer{
		Answer: "This is an answer",
	}
	s.repo.On("CreateAnswer", s.ctx, questionID, answer).Return(answerID, nil)
	result, err := s.svc.AddAnswer(s.ctx, questionID, answer)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), answerID, result)
}

func (s *serviceSuite) TestUpdateAnswer_Success() {
	questionID := primitive.NewObjectID().Hex()
	answerID := primitive.NewObjectID().Hex()
	var answer = &entities.Answer{
		Answer: "This is an answer",
	}
	s.repo.On("UpdateAnswer", s.ctx, questionID, answerID, answer.Answer).Return("success", nil)
	result, err := s.svc.UpdateAnswer(s.ctx, questionID, answerID, answer)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), result, "success")
}
