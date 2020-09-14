package transport

import (
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/questions"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"testing"
)

type decodersSuite struct {
	ctx context.Context
	suite.Suite
}

func (s *decodersSuite) SetupTest() {
	s.ctx = context.Background()
}

func TestDecodersSuite(t *testing.T) {
	s := new(decodersSuite)
	suite.Run(t, s)
}

func (s *decodersSuite) TestDecodeGetAllQuestionsRequest_QueryParam() {
	b, err := json.Marshal(questions.GetAllQuestionsRequest{ItemsPerPage: 0})
	if err != nil {
		s.T().Fail()
	}
	buf := bytes.NewBuffer(b)
	rdr := ioutil.NopCloser(buf)
	request, _ := http.NewRequest(http.MethodGet, "/questions", rdr)
	q := request.URL.Query()
	q.Add("perPage", "100")
	request.URL.RawQuery = q.Encode()
	untyped, err := decodeGetAllQuestionsRequest(s.ctx, request)
	assert.Nil(s.T(), err)
	req := untyped.(*questions.GetAllQuestionsRequest)
	assert.IsType(s.T(), questions.GetAllQuestionsRequest{}, *req)
}

func (s *decodersSuite) TestDecodeGetAllQuestionsRequest_NoQueryParams() {
	b, err := json.Marshal(questions.GetAllQuestionsRequest{ItemsPerPage: 0})
	if err != nil {
		s.T().Fail()
	}
	buf := bytes.NewBuffer(b)
	rdr := ioutil.NopCloser(buf)
	request, _ := http.NewRequest(http.MethodGet, "/questions", rdr)
	untyped, err := decodeGetAllQuestionsRequest(s.ctx, request)
	assert.Nil(s.T(), err)
	req := untyped.(*questions.GetAllQuestionsRequest)
	assert.IsType(s.T(), questions.GetAllQuestionsRequest{}, *req)
}

func (s *decodersSuite) TestDecodeGetAllQuestionsRequest_InvalidQueryParam() {
	b, err := json.Marshal(questions.GetAllQuestionsRequest{ItemsPerPage: 0})
	if err != nil {
		s.T().Fail()
	}
	buf := bytes.NewBuffer(b)
	rdr := ioutil.NopCloser(buf)
	request, _ := http.NewRequest(http.MethodGet, "/questions", rdr)
	q := request.URL.Query()
	q.Add("perPage", "")
	request.URL.RawQuery = q.Encode()
	untyped, err := decodeGetAllQuestionsRequest(s.ctx, request)
	assert.Nil(s.T(), err)
	req := untyped.(*questions.GetAllQuestionsRequest)
	assert.IsType(s.T(), questions.GetAllQuestionsRequest{}, *req)
}

func (s *decodersSuite) TestDecodeGetQuestionRequest_Success() {
	qID := primitive.NewObjectID().Hex()
	b, err := json.Marshal(questions.GetQuestionRequest{QuestionID: qID})
	if err != nil {
		s.T().Fail()
	}
	buf := bytes.NewBuffer(b)
	rdr := ioutil.NopCloser(buf)
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/questions/%s", qID), rdr)
	vars := map[string]string{
		QuestionID: qID,
	}
	request = mux.SetURLVars(request, vars)
	untyped, err := decodeGetQuestionRequest(s.ctx, request)
	assert.Nil(s.T(), err)
	req := untyped.(*questions.GetQuestionRequest)
	assert.IsType(s.T(), questions.GetQuestionRequest{}, *req)
}

func (s *decodersSuite) TestDecodeGetQuestionRequest_MissingParams() {
	qID := primitive.NewObjectID().Hex()
	b, err := json.Marshal(questions.GetQuestionRequest{QuestionID: qID})
	if err != nil {
		s.T().Fail()
	}
	buf := bytes.NewBuffer(b)
	rdr := ioutil.NopCloser(buf)
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/questions/%s", qID), rdr)
	untyped, err := decodeGetQuestionRequest(s.ctx, request)
	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), untyped)
	assert.Equal(s.T(), err.Error(), "missing questionID on request")
}

func (s *decodersSuite) TestDecodeCreateQuestionRequest() {
	b, err := json.Marshal(questions.CreateQuestionRequest{Title: "Title", Description: "Description"})
	if err != nil {
		s.T().Fail()
	}
	buf := bytes.NewBuffer(b)
	rdr := ioutil.NopCloser(buf)
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/questions"), rdr)
	untyped, err := decodeCreateQuestionRequest(s.ctx, request)
	assert.Nil(s.T(), err)
	req := untyped.(*questions.CreateQuestionRequest)
	assert.IsType(s.T(), questions.CreateQuestionRequest{}, *req)
}

func (s *decodersSuite) TestDecodeUpdateQuestionRequest() {
	qID := primitive.NewObjectID().Hex()
	b, err := json.Marshal(questions.UpdateQuestionRequest{ID: qID, Title: "Title", Description: "Description"})
	if err != nil {
		s.T().Fail()
	}
	buf := bytes.NewBuffer(b)
	rdr := ioutil.NopCloser(buf)
	request, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("/questions/%s", qID), rdr)
	vars := map[string]string{
		QuestionID: qID,
	}
	request = mux.SetURLVars(request, vars)
	untyped, err := decodeUpdateQuestionRequest(s.ctx, request)
	assert.Nil(s.T(), err)
	req := untyped.(*questions.UpdateQuestionRequest)
	assert.IsType(s.T(), questions.UpdateQuestionRequest{}, *req)
}

func (s *decodersSuite) TestDecodeDeleteQuestionRequest() {
	qID := primitive.NewObjectID().Hex()
	b, err := json.Marshal(questions.DeleteQuestionRequest{QuestionID: qID})
	if err != nil {
		s.T().Fail()
	}
	buf := bytes.NewBuffer(b)
	rdr := ioutil.NopCloser(buf)
	request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/questions/%s", qID), rdr)
	vars := map[string]string{
		QuestionID: qID,
	}
	request = mux.SetURLVars(request, vars)
	untyped, err := decodeDeleteQuestionRequest(s.ctx, request)
	assert.Nil(s.T(), err)
	req := untyped.(*questions.DeleteQuestionRequest)
	assert.IsType(s.T(), questions.DeleteQuestionRequest{}, *req)
}

func (s *decodersSuite) TestDecodeCreateAnswerRequest() {
	qID := primitive.NewObjectID().Hex()
	b, err := json.Marshal(questions.AddAnswerRequest{QuestionID: qID, Answer: "Answer"})
	if err != nil {
		s.T().Fail()
	}
	buf := bytes.NewBuffer(b)
	rdr := ioutil.NopCloser(buf)
	request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/questions/%s/answers", qID), rdr)
	vars := map[string]string{
		QuestionID: qID,
	}
	request = mux.SetURLVars(request, vars)
	untyped, err := decodeAddAnswerRequest(s.ctx, request)
	assert.Nil(s.T(), err)
	req := untyped.(*questions.AddAnswerRequest)
	assert.IsType(s.T(), questions.AddAnswerRequest{}, *req)
}

func (s *decodersSuite) TestDecodeUpdateAnswerRequest() {
	qID := primitive.NewObjectID().Hex()
	aID := primitive.NewObjectID().Hex()
	b, err := json.Marshal(questions.UpdateAnswerRequest{
		QuestionID: qID,
		AnswerID:   aID,
		Answer:     "Answer",
	})
	if err != nil {
		s.T().Fail()
	}
	buf := bytes.NewBuffer(b)
	rdr := ioutil.NopCloser(buf)
	request, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("/questions"), rdr)
	vars := map[string]string{
		QuestionID: qID,
		AnswerID:   aID,
	}
	request = mux.SetURLVars(request, vars)
	untyped, err := decodeUpdateAnswerRequest(s.ctx, request)
	assert.Nil(s.T(), err)
	req := untyped.(*questions.UpdateAnswerRequest)
	assert.IsType(s.T(), questions.UpdateAnswerRequest{}, *req)
}
