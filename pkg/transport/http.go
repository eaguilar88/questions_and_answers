package transport

import (
	"fmt"
	"net/http"

	"bitbucket.org/aveaguilar/questions_and_answers/pkg/questions"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func AddQuestionEndpoints(rtr *mux.Router, eps questions.Endpoints, options []kitHttp.ServerOption) {
	//All Questions
	rtr.Methods(http.MethodGet).Path("/questions").Handler(GetAllQuestionsHandler(eps.GetAll, options))

	{
		//Question by ID
		path := fmt.Sprintf("/questions/{%s}", QuestionID)
		rtr.Methods(http.MethodGet).Path(path).Handler(GetQuestionHandler(eps.GetQuestion, options))
	}

	//Create Question
	rtr.Methods(http.MethodPost).Path("/questions").Handler(CreateQuestionHandler(eps.CreateQuestion, options))

	{
		//Update question
		path := fmt.Sprintf("/questions/{%s}", QuestionID)
		rtr.Methods(http.MethodPatch).Path(path).Handler(UpdateQuestionHandler(eps.UpdateQuestion, options))
	}

	{
		//Delete question
		path := fmt.Sprintf("/questions/{%s}", QuestionID)
		rtr.Methods(http.MethodDelete).Path(path).Handler(DeleteQuestionHandler(eps.DeleteQuestion, options))
	}
	//Answers related endpoints
	{
		//Add answer to a question
		path := fmt.Sprintf("/questions/{%s}/answers", QuestionID)
		rtr.Methods(http.MethodPut).Path(path).Handler(AddAnswerHandler(eps.CreateAnswer, options))
	}
	{
		//Edit answer of a question
		path := fmt.Sprintf("/questions/{%s}/answers/{%s}", QuestionID, AnswerID)
		rtr.Methods(http.MethodPatch).Path(path).Handler(UpdateAnswerHandler(eps.UpdateAnswer, options))
	}
}
