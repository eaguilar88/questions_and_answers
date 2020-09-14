package testhelp

import (
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
)

func MockQuestionArray() *entities.Questions {
	questions := make(entities.Questions, 0, 2)
	q := entities.Question{
		ID:          primitive.NewObjectID(),
		Title:       "Generic question?",
		Description: "Generic explanation",
		Answers:     MockAnswers(rand.Intn(3)),
	}
	questions = append(questions, q)
	q = entities.Question{
		ID:          primitive.NewObjectID(),
		Title:       "Generic question?",
		Description: "Generic explanation",
		Answers:     MockAnswers(rand.Intn(3)),
	}
	questions = append(questions, q)
	return &questions
}

func MockAnswers(quantity int) []*entities.Answer {
	if quantity == 0 {
		return nil
	}
	answers := make([]*entities.Answer, 0, 2)
	a := &entities.Answer{
		ID:     primitive.NewObjectID(),
		Answer: "Right answer",
	}
	answers = append(answers, a)
	a = &entities.Answer{
		ID:     primitive.NewObjectID(),
		Answer: "Not so good answer",
	}
	answers = append(answers, a)
	return answers
}
