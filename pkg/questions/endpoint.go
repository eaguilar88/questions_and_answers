package questions

import (
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

//Endpoints _
type Endpoints struct {
	GetAll         endpoint.Endpoint
	CreateQuestion endpoint.Endpoint
	GetQuestion    endpoint.Endpoint
	UpdateQuestion endpoint.Endpoint
	DeleteQuestion endpoint.Endpoint
	CreateAnswer   endpoint.Endpoint
	UpdateAnswer   endpoint.Endpoint
}

//MakeEndpoints _
func MakeEndpoints(svc Service, logger log.Logger, middlewares []endpoint.Middleware) Endpoints {
	return Endpoints{
		GetAll:         wrapEndpoint(makeGetAllEndpoint(svc, logger), middlewares),
		CreateQuestion: wrapEndpoint(makeCreateQuestionEndpoint(svc, logger), middlewares),
		GetQuestion:    wrapEndpoint(makeGetQuestionEndpoint(svc, logger), middlewares),
		UpdateQuestion: wrapEndpoint(makeUpdateQuestionEndpoint(svc, logger), middlewares),
		DeleteQuestion: wrapEndpoint(makeDeleteQuestionEndpoint(svc, logger), middlewares),
		CreateAnswer:   wrapEndpoint(makeCreateAnswerEndpoint(svc, logger), middlewares),
		UpdateAnswer:   wrapEndpoint(makeUpdateAnswerEndpoint(svc, logger), middlewares),
	}
}

func makeGetAllEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*GetAllQuestionsRequest)
		if !ok {
			level.Error(logger).Log("message", "bad request type")
			return nil, errors.New("bad request")
		}
		questions, err := svc.GetAll(ctx, req.ItemsPerPage)
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
		req, ok := request.(*GetQuestionRequest)
		if !ok {
			level.Error(logger).Log("message", "bad request type")
			return nil, errors.New("bad request")
		}
		question, err := svc.GetQuestionByID(ctx, req.QuestionID)
		if err != nil {
			_ = level.Error(logger).Log("endpoint", "get question by id error", err)
			return nil, err
		}
		return &GetQuestionResponse{
			Question: question,
		}, nil
	}
}

func makeCreateQuestionEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*CreateQuestionRequest)
		if !ok {
			level.Error(logger).Log("message", "bad request type")
			return nil, errors.New("bad request")
		}
		q := &entities.Question{
			Title:       req.Title,
			Description: req.Description,
		}
		createdID, err := svc.CreateQuestion(ctx, q)
		if err != nil {
			_ = level.Error(logger).Log("endpoint", "create question error", err)
			return nil, err
		}
		return &CreateResourceResponse{
			CreatedID: createdID,
		}, nil
	}
}

func makeUpdateQuestionEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*UpdateQuestionRequest)
		if !ok {
			level.Error(logger).Log("message", "bad request type")
			return nil, errors.New("bad request")
		}

		q := &entities.Question{
			Title:       req.Title,
			Description: req.Description,
		}
		message, err := svc.UpdateQuestion(ctx, req.ID, q)
		if err != nil {
			_ = level.Error(logger).Log("endpoint", "update question error", err)
			return nil, err
		}
		return &HandleResourceResponse{
			Message: message,
		}, nil
	}
}

func makeDeleteQuestionEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*DeleteQuestionRequest)
		if !ok {
			level.Error(logger).Log("message", "bad request type")
			return nil, errors.New("bad request")
		}

		message, err := svc.DeleteQuestion(ctx, req.QuestionID)
		if err != nil {
			_ = level.Error(logger).Log("endpoint", "update question error", err)
			return nil, err
		}
		return &HandleResourceResponse{
			Message: message,
		}, nil
	}
}

func makeCreateAnswerEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*AddAnswerRequest)
		if !ok {
			level.Error(logger).Log("message", "bad request type")
			return nil, errors.New("bad request")
		}

		a := &entities.Answer{
			Answer: req.Answer,
		}
		createdID, err := svc.AddAnswer(ctx, req.QuestionID, a)
		if err != nil {
			_ = level.Error(logger).Log("endpoint", "update question error", err)
			return nil, err
		}
		return &CreateResourceResponse{
			CreatedID: createdID,
		}, nil
	}
}

func makeUpdateAnswerEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*UpdateAnswerRequest)
		if !ok {
			level.Error(logger).Log("message", "bad request type")
			return nil, errors.New("bad request")
		}

		a := &entities.Answer{
			Answer: req.Answer,
		}
		message, err := svc.UpdateAnswer(ctx, req.QuestionID, req.AnswerID, a)
		if err != nil {
			_ = level.Error(logger).Log("endpoint", "update question error", err)
			return nil, err
		}
		return &HandleResourceResponse{
			Message: message,
		}, nil
	}
}

func wrapEndpoint(e endpoint.Endpoint, middlewares []endpoint.Middleware) endpoint.Endpoint {
	for _, m := range middlewares {
		e = m(e)
	}
	return e
}
