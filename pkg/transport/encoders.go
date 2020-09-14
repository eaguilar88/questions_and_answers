package transport

import (
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/questions"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func encodeGetAllQuestionsResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp, ok := response.(*questions.GetAllQuestionsResponse)
	if !ok {
		return errors.New("error decoding")
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}
