package questions

import (
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/testhelp"
	"context"
	"errors"
	"math/rand"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type endpointsSuite struct {
	ctx    context.Context
	svc    *QuestionServiceMock
	logger log.Logger
	suite.Suite
}

func (s *endpointsSuite) SetupTest() {
	s.ctx = context.Background()
	s.svc = new(QuestionServiceMock)
	s.logger = log.NewNopLogger()
}

func TestQuestionEndpoints(t *testing.T) {
	s := new(endpointsSuite)
	suite.Run(t, s)
}

func (s *endpointsSuite) TestGetAllQuestions_Success() {
	var questions = testhelp.MockQuestionArray()
	limit := rand.Intn(4)
	s.svc.On("GetAll", s.ctx, limit).Return(questions, nil)
	ep := makeGetAllEndpoint(s.svc, s.logger)
	req := &GetAllQuestionsRequest{ItemsPerPage: limit}
	untyped, err := ep(s.ctx, req)
	assert.Nil(s.T(), err)
	resp, ok := untyped.(*GetAllQuestionsResponse)
	assert.True(s.T(), ok)
	assert.NotNil(s.T(), resp)
	assert.IsType(s.T(), GetAllQuestionsResponse{}, *resp)
}

func (s *endpointsSuite) TestGetAllQuestions_BadRequest() {
	var questions = testhelp.MockQuestionArray()
	limit := rand.Intn(4)
	s.svc.On("GetAll", s.ctx, limit).Return(questions, nil)
	ep := makeGetAllEndpoint(s.svc, s.logger)
	req := &GetQuestionRequest{QuestionID: "invalid-id"}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestGetAllQuestions_InternalError() {
	limit := rand.Intn(4)
	s.svc.On("GetAll", s.ctx, limit).Return(nil, errors.New("error getting all questions"))
	ep := makeGetAllEndpoint(s.svc, s.logger)
	req := &GetAllQuestionsRequest{ItemsPerPage: limit}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "error getting all questions")
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestGetQuestion_Success() {
	questionID := primitive.NewObjectID()
	var question = &entities.Question{
		ID:          questionID,
		Title:       "Generic question?",
		Description: "Generic description",
		Answers:     nil,
	}
	s.svc.On("GetQuestionByID", s.ctx, questionID.Hex()).Return(question, nil)
	ep := makeGetQuestionEndpoint(s.svc, s.logger)
	req := &GetQuestionRequest{QuestionID: questionID.Hex()}
	untyped, err := ep(s.ctx, req)
	assert.Nil(s.T(), err)
	resp, ok := untyped.(*GetQuestionResponse)
	assert.True(s.T(), ok)
	assert.NotNil(s.T(), resp)
	assert.IsType(s.T(), GetQuestionResponse{}, *resp)
	assert.Equal(s.T(), questionID, resp.Question.ID)
}

func (s *endpointsSuite) TestGetQuestion_BadRequest() {
	questionID := rand.Intn(4)
	s.svc.On("GetQuestionByID", s.ctx, questionID).Return(nil, errors.New("bad request"))
	ep := makeGetQuestionEndpoint(s.svc, s.logger)
	req := &GetAllQuestionsRequest{ItemsPerPage: 0}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "bad request")
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestGetQuestion_NotFound() {
	questionID := primitive.NewObjectID().Hex()
	s.svc.On("GetQuestionByID", s.ctx, questionID).Return(nil, errors.New("no documents in result"))
	ep := makeGetQuestionEndpoint(s.svc, s.logger)
	req := &GetQuestionRequest{QuestionID: questionID}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "no documents in result")
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestCreateQuestion_Success() {
	questionID := primitive.NewObjectID().Hex()
	var question = &entities.Question{
		Title:       "Generic question?",
		Description: "Generic description",
		Answers:     nil,
	}
	s.svc.On("CreateQuestion", s.ctx, question).Return(questionID, nil)
	ep := makeCreateQuestionEndpoint(s.svc, s.logger)
	req := &CreateQuestionRequest{
		Title:       "Generic question?",
		Description: "Generic description",
	}
	untyped, err := ep(s.ctx, req)
	assert.Nil(s.T(), err)
	resp, ok := untyped.(*CreateResourceResponse)
	assert.True(s.T(), ok)
	assert.NotNil(s.T(), resp)
	assert.IsType(s.T(), CreateResourceResponse{}, *resp)
}

func (s *endpointsSuite) TestCreateQuestion_BadRequest() {
	questionID := rand.Intn(4)
	s.svc.On("CreateQuestion", s.ctx, questionID).Return(nil, errors.New("bad request"))
	ep := makeCreateQuestionEndpoint(s.svc, s.logger)
	req := &GetAllQuestionsRequest{ItemsPerPage: 0}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "bad request")
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestCreateQuestion_InternalError() {
	var question = &entities.Question{
		Title:       "Generic question?",
		Description: "Generic description",
		Answers:     nil,
	}
	s.svc.On("CreateQuestion", s.ctx, question).Return(nil, errors.New("create question error"))
	ep := makeCreateQuestionEndpoint(s.svc, s.logger)
	req := &CreateQuestionRequest{
		Title:       "Generic question?",
		Description: "Generic description",
	}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "create question error")
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestUpdateQuestion_Success() {
	questionID := primitive.NewObjectID()
	var question = &entities.Question{
		Title:       "Generic question?",
		Description: "Generic description",
		Answers:     nil,
	}
	s.svc.On("UpdateQuestion", s.ctx, questionID.Hex(), question).Return("success", nil)
	ep := makeUpdateQuestionEndpoint(s.svc, s.logger)
	req := &UpdateQuestionRequest{
		ID:          questionID.Hex(),
		Title:       "Generic question?",
		Description: "Generic description",
		Answers:     nil,
	}
	untyped, err := ep(s.ctx, req)
	assert.Nil(s.T(), err)
	resp, ok := untyped.(*HandleResourceResponse)
	assert.True(s.T(), ok)
	assert.NotNil(s.T(), resp)
	assert.IsType(s.T(), HandleResourceResponse{}, *resp)
	assert.Equal(s.T(), "success", resp.Message)
}

func (s *endpointsSuite) TestUpdateQuestion_BadRequest() {
	questionID := primitive.NewObjectID()
	var question = &entities.Question{
		ID:          questionID,
		Title:       "Generic question?",
		Description: "Generic description",
		Answers:     nil,
	}
	s.svc.On("UpdateQuestion", s.ctx, questionID, question).Return(nil, errors.New("bad request"))
	ep := makeUpdateQuestionEndpoint(s.svc, s.logger)
	req := &GetAllQuestionsRequest{ItemsPerPage: 0}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "bad request")
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestUpdateQuestion_InternalError() {
	questionID := primitive.NewObjectID()
	var question = &entities.Question{
		Title:       "Generic question?",
		Description: "Generic description",
		Answers:     nil,
	}
	s.svc.On("UpdateQuestion", s.ctx, questionID.Hex(), question).Return(nil, errors.New("update question error"))
	ep := makeUpdateQuestionEndpoint(s.svc, s.logger)
	req := &UpdateQuestionRequest{
		ID:          questionID.Hex(),
		Title:       "Generic question?",
		Description: "Generic description",
	}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "update question error")
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestDeleteQuestion_Success() {
	questionID := primitive.NewObjectID()
	s.svc.On("DeleteQuestion", s.ctx, questionID.Hex()).Return("success", nil)
	ep := makeDeleteQuestionEndpoint(s.svc, s.logger)
	req := &DeleteQuestionRequest{
		QuestionID: questionID.Hex(),
	}
	untyped, err := ep(s.ctx, req)
	assert.Nil(s.T(), err)
	resp, ok := untyped.(*HandleResourceResponse)
	assert.True(s.T(), ok)
	assert.NotNil(s.T(), resp)
	assert.IsType(s.T(), HandleResourceResponse{}, *resp)
	assert.Equal(s.T(), "success", resp.Message)
}

func (s *endpointsSuite) TestDeleteQuestion_BadRequest() {
	questionID := primitive.NewObjectID()
	s.svc.On("DeleteQuestion", s.ctx, questionID.Hex()).Return(nil, errors.New("bad request"))
	ep := makeDeleteQuestionEndpoint(s.svc, s.logger)
	req := &GetAllQuestionsRequest{ItemsPerPage: 0}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "bad request")
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestDeleteQuestion_InternalError() {
	questionID := primitive.NewObjectID()
	s.svc.On("DeleteQuestion", s.ctx, questionID.Hex()).Return(nil, errors.New("delete question error"))
	ep := makeDeleteQuestionEndpoint(s.svc, s.logger)
	req := &DeleteQuestionRequest{
		QuestionID: questionID.Hex(),
	}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "delete question error")
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestCreateAnswer_Success() {
	questionID := primitive.NewObjectID()
	answerID := primitive.NewObjectID().Hex()
	var answer = &entities.Answer{
		Answer: "This is an answer",
	}
	s.svc.On("AddAnswer", s.ctx, questionID.Hex(), answer).Return(answerID, nil)
	ep := makeCreateAnswerEndpoint(s.svc, s.logger)
	req := &AddAnswerRequest{
		QuestionID: questionID.Hex(),
		Answer:     "This is an answer",
	}
	untyped, err := ep(s.ctx, req)
	assert.Nil(s.T(), err)
	resp, ok := untyped.(*CreateResourceResponse)
	assert.True(s.T(), ok)
	assert.NotNil(s.T(), resp)
	assert.IsType(s.T(), CreateResourceResponse{}, *resp)
	assert.Equal(s.T(), answerID, resp.CreatedID)
}

func (s *endpointsSuite) TestCreateAnswer_BadRequest() {
	questionID := primitive.NewObjectID()
	var answer = &entities.Answer{
		Answer: "This is an answer",
	}
	s.svc.On("AddAnswer", s.ctx, questionID.Hex(), answer).Return("", errors.New("bad request"))
	ep := makeCreateAnswerEndpoint(s.svc, s.logger)
	req := &CreateQuestionRequest{
		Title:       questionID.Hex(),
		Description: "This is an answer",
	}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "bad request")
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestCreateAnswer_InternalError() {
	questionID := primitive.NewObjectID()
	var answer = &entities.Answer{
		Answer: "This is an answer",
	}
	s.svc.On("AddAnswer", s.ctx, questionID.Hex(), answer).Return("", errors.New("error adding answer to question"))
	ep := makeCreateAnswerEndpoint(s.svc, s.logger)
	req := &AddAnswerRequest{
		QuestionID: questionID.Hex(),
		Answer:     "This is an answer",
	}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "error adding answer to question")
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestUpdateAnswer_Success() {
	questionID := primitive.NewObjectID()
	answerID := primitive.NewObjectID().Hex()
	var answer = &entities.Answer{
		Answer: "This is an answer",
	}
	s.svc.On("UpdateAnswer", s.ctx, questionID.Hex(), answerID, answer).Return("success", nil)
	ep := makeUpdateAnswerEndpoint(s.svc, s.logger)
	req := &UpdateAnswerRequest{
		QuestionID: questionID.Hex(),
		AnswerID:   answerID,
		Answer:     "This is an answer",
	}
	untyped, err := ep(s.ctx, req)
	assert.Nil(s.T(), err)
	resp, ok := untyped.(*HandleResourceResponse)
	assert.True(s.T(), ok)
	assert.NotNil(s.T(), resp)
	assert.IsType(s.T(), HandleResourceResponse{}, *resp)
	assert.Equal(s.T(), "success", resp.Message)
}

func (s *endpointsSuite) TestUpdateAnswer_BadRequest() {
	questionID := primitive.NewObjectID()
	answerID := primitive.NewObjectID().Hex()
	var answer = &entities.Answer{
		Answer: "This is an answer",
	}
	s.svc.On("UpdateAnswer", s.ctx, questionID.Hex(), answerID, answer).Return("", errors.New("bad request"))
	ep := makeUpdateAnswerEndpoint(s.svc, s.logger)
	req := &GetAllQuestionsRequest{ItemsPerPage: 0}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "bad request")
	assert.Nil(s.T(), untyped)
}

func (s *endpointsSuite) TestUpdateAnswer_InternalError() {
	questionID := primitive.NewObjectID()
	answerID := primitive.NewObjectID().Hex()
	var answer = &entities.Answer{
		Answer: "This is an answer",
	}
	s.svc.On("UpdateAnswer", s.ctx, questionID.Hex(), answerID, answer).Return("", errors.New("error updating answer"))
	ep := makeUpdateAnswerEndpoint(s.svc, s.logger)
	req := &UpdateAnswerRequest{
		QuestionID: questionID.Hex(),
		AnswerID:   answerID,
		Answer:     "This is an answer",
	}
	untyped, err := ep(s.ctx, req)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "error updating answer")
	assert.Nil(s.T(), untyped)
}
