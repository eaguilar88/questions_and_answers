package questions

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Endpoints struct {
	GetAll         endpoint.Endpoint
	CreateQuestion endpoint.Endpoint
	GetQuestion    endpoint.Endpoint
	UpdateQuestion endpoint.Endpoint
	DeleteQuestion endpoint.Endpoint
}

func MakeEndpoints(svc Service, logger log.Logger, middlewares []endpoint.Middleware) Endpoints {
	return Endpoints{
		GetAll:         wrapEndpoint(makeGetAllEndpoint(svc, logger), middlewares),
		CreateQuestion: wrapEndpoint(makeCreateQuestionEndpoint(svc, logger), middlewares),
		GetQuestion:    wrapEndpoint(makeGetQuestionEndpoint(svc, logger), middlewares),
		UpdateQuestion: wrapEndpoint(makeUpdateQuestionEndpoint(svc, logger), middlewares),
		DeleteQuestion: wrapEndpoint(makeDeleteQuestionEndpoint(svc, logger), middlewares),
	}
}

func makeGetAllEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		msg := svc.GetAll(ctx)
		return &IsPalResponse{
			Message: msg,
		}, nil
	}

}

func makeGetQuestionEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(ReverseRequest)
		if !ok {
			level.Error(logger).Log("message", "invalid request")
			return nil, errors.New("invalid request")
		}
		reverseString := svc.Reverse(ctx, req.Word)
		return &ReverseResponse{
			Word: reverseString,
		}, nil
	}
}

func wrapEndpoint(e endpoint.Endpoint, middlewares []endpoint.Middleware) endpoint.Endpoint {
	for _, m := range middlewares {
		e = m(e)
	}
	return e
}
