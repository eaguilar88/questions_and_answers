package transport

import (
	"github.com/go-kit/kit/endpoint"
	kitHttp "github.com/go-kit/kit/transport/http"
)

func GetAllQuestionsHandler(ep endpoint.Endpoint, options []kitHttp.ServerOption) *kitHttp.Server {
	return kitHttp.NewServer(
		ep,
		decodeGetAllQuestionsRequest,
		encodeGetAllQuestionsResponse,
		options...,
	)
}
