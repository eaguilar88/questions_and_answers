package transport

import (
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/questions"
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/testhelp"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
	"testing"
)

type encodersSuite struct {
	ctx context.Context
	rw  *httptest.ResponseRecorder
	suite.Suite
}

func (s *encodersSuite) SetupTest() {
	s.ctx = context.Background()
	s.rw = new(httptest.ResponseRecorder)
}

func TestEncodersSuite(t *testing.T) {
	s := new(encodersSuite)
	suite.Run(t, s)
}

func (s *encodersSuite) TestEncodeGetAllQuestionsResponse() {
	resp := questions.GetAllQuestionsResponse{Questions: testhelp.MockQuestionArray()}
	err := encodeGetAllQuestionsResponse(s.ctx, s.rw, &resp)
	assert.Nil(s.T(), err)
}

func (s *encodersSuite) TestEncodeCreateQuestionResponse() {
	resp := questions.CreateResourceResponse{CreatedID: primitive.NewObjectID().Hex()}
	err := encodeCreateQuestionResponse(s.ctx, s.rw, &resp)
	assert.Nil(s.T(), err)
}

func (s *encodersSuite) TestEncodeGetQuestionResponse() {
	qs := testhelp.MockQuestionArray()
	resp := questions.GetQuestionResponse{Question: &(*qs)[0]}
	err := encodeGetQuestionResponse(s.ctx, s.rw, &resp)
	assert.Nil(s.T(), err)
}

func (s *encodersSuite) TestEncodeUpdateQuestionResponse() {
	resp := questions.HandleResourceResponse{Message: "success"}
	err := encodeUpdateQuestionResponse(s.ctx, s.rw, &resp)
	assert.Nil(s.T(), err)
}

func (s *encodersSuite) TestEncodeDeleteQuestionResponse() {
	resp := questions.HandleResourceResponse{Message: "success"}
	err := encodeDeleteQuestionResponse(s.ctx, s.rw, &resp)
	assert.Nil(s.T(), err)
}

func (s *encodersSuite) TestEncodeCreateAnswerResponse() {
	resp := questions.CreateResourceResponse{CreatedID: primitive.NewObjectID().Hex()}
	err := encodeAddAnswerResponse(s.ctx, s.rw, &resp)
	assert.Nil(s.T(), err)
}

func (s *encodersSuite) TestEncodeUpdateAnswerResponse() {
	resp := questions.HandleResourceResponse{Message: "success"}
	err := encodeUpdateAnswerResponse(s.ctx, s.rw, &resp)
	assert.Nil(s.T(), err)
}
