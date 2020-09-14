package questions

import (
	"context"
	"github.com/go-kit/kit/log/level"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

//Endpoints _
type Endpoints struct {
	GetAll         endpoint.Endpoint
	CreateQuestion endpoint.Endpoint
	GetQuestion    endpoint.Endpoint
	UpdateQuestion endpoint.Endpoint
	DeleteQuestion endpoint.Endpoint
}

//MakeEndpoints _
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
		level.Debug(logger).Log("message", "endpoint: Get All")
		questions, err := svc.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return &GetAllQuestionsResponse{
			Questions: questions,
		}, nil
	}
}

func makeGetQuestionEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return nil, nil
	}
}

func makeCreateQuestionEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return nil, nil
	}
}

func makeUpdateQuestionEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return nil, nil
	}
}

func makeDeleteQuestionEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return nil, nil
	}
}

func wrapEndpoint(e endpoint.Endpoint, middlewares []endpoint.Middleware) endpoint.Endpoint {
	for _, m := range middlewares {
		e = m(e)
	}
	return e
}
