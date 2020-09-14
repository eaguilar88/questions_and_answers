package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"bitbucket.org/aveaguilar/questions_and_answers/pkg/questions"
)

func encodeGetAllQuestionsResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp, ok := response.(*questions.GetAllQuestionsResponse)
	if !ok {
		return errors.New("error encoding response")
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

func encodeGetQuestionResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp, ok := response.(*questions.GetQuestionResponse)
	if !ok {
		return errors.New("error encoding response")
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

func encodeCreateQuestionResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp, ok := response.(*questions.CreateResourceResponse)
	if !ok {
		return errors.New("error encoding response")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(resp)
}

func encodeUpdateQuestionResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp, ok := response.(*questions.HandleResourceResponse)
	if !ok {
		return errors.New("error encoding response")
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

func encodeDeleteQuestionResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp, ok := response.(*questions.HandleResourceResponse)
	if !ok {
		return errors.New("error encoding response")
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

func encodeAddAnswerResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp, ok := response.(*questions.CreateResourceResponse)
	if !ok {
		return errors.New("error encoding response")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(resp)
}

func encodeUpdateAnswerResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp, ok := response.(*questions.HandleResourceResponse)
	if !ok {
		return errors.New("error encoding response")
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}
