package transport

import (
	"context"
	"net/http"
)

func decodeGetAllQuestionsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}
