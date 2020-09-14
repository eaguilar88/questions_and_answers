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

func GetQuestionHandler(ep endpoint.Endpoint, options []kitHttp.ServerOption) *kitHttp.Server {
	return kitHttp.NewServer(
		ep,
		decodeGetQuestionRequest,
		encodeGetQuestionResponse,
		options...,
	)
}

func CreateQuestionHandler(ep endpoint.Endpoint, options []kitHttp.ServerOption) *kitHttp.Server {
	return kitHttp.NewServer(
		ep,
		decodeCreateQuestionRequest,
		encodeCreateQuestionResponse,
		options...,
	)
}

func UpdateQuestionHandler(ep endpoint.Endpoint, options []kitHttp.ServerOption) *kitHttp.Server {
	return kitHttp.NewServer(
		ep,
		decodeUpdateQuestionRequest,
		encodeUpdateQuestionResponse,
		options...,
	)
}

func DeleteQuestionHandler(ep endpoint.Endpoint, options []kitHttp.ServerOption) *kitHttp.Server {
	return kitHttp.NewServer(
		ep,
		decodeDeleteQuestionRequest,
		encodeDeleteQuestionResponse,
		options...,
	)
}

func AddAnswerHandler(ep endpoint.Endpoint, options []kitHttp.ServerOption) *kitHttp.Server {
	return kitHttp.NewServer(
		ep,
		decodeAddAnswerRequest,
		encodeAddAnswerResponse,
		options...,
	)
}

func UpdateAnswerHandler(ep endpoint.Endpoint, options []kitHttp.ServerOption) *kitHttp.Server {
	return kitHttp.NewServer(
		ep,
		decodeUpdateAnswerRequest,
		encodeUpdateAnswerResponse,
		options...,
	)
}
