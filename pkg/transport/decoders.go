package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"bitbucket.org/aveaguilar/questions_and_answers/pkg/questions"
	"github.com/gorilla/mux"
)

func decodeGetAllQuestionsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	perPage, exists := r.URL.Query()["perPage"]
	if !exists {
		return &questions.GetAllQuestionsRequest{
			ItemsPerPage: 0,
		}, nil
	}
	pp, err := strconv.Atoi(perPage[0])
	if err != nil {
		return &questions.GetAllQuestionsRequest{
			ItemsPerPage: 0,
		}, nil
	}
	return &questions.GetAllQuestionsRequest{
		ItemsPerPage: pp,
	}, nil
}

func decodeGetQuestionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	questionID, exists := vars[QuestionID]
	if !exists {
		return nil, errors.New("missing questionID on request")
	}
	return &questions.GetQuestionRequest{
		QuestionID: questionID,
	}, nil
}

func decodeCreateQuestionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req questions.CreateQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.New("error decoding the request")
	}

	return &req, nil
}

func decodeUpdateQuestionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	questionID, exists := vars[QuestionID]
	if !exists {
		return nil, errors.New("missing questionID on request")
	}
	var req questions.UpdateQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.New("error decoding the request")
	}
	req.ID = questionID
	return &req, nil
}

func decodeDeleteQuestionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	questionID, exists := vars[QuestionID]
	if !exists {
		return nil, errors.New("missing questionID on request")
	}
	return &questions.DeleteQuestionRequest{
		QuestionID: questionID,
	}, nil
}

func decodeAddAnswerRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	questionID, exists := vars[QuestionID]
	if !exists {
		return nil, errors.New("missing questionID on request")
	}

	var req questions.AddAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.New("error decoding the request")
	}
	req.QuestionID = questionID
	return &req, nil
}

func decodeUpdateAnswerRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	questionID, exists := vars[QuestionID]
	if !exists {
		return nil, errors.New("missing questionID on request")
	}
	answerID, exists := vars[AnswerID]
	if !exists {
		return nil, errors.New("missing answerID on request")
	}
	var req questions.UpdateAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.New("error decoding the request")
	}
	req.QuestionID = questionID
	req.AnswerID = answerID
	return &req, nil
}
