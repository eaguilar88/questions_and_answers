package answers

import (
	"context"
	"strings"

	"github.com/go-kit/kit/log"
)

type Service interface {
	IsPal(context.Context, string) string
	Reverse(context.Context, string) string
}

type myStringService struct {
	repository Repository
	log        log.Logger
}

func NewService(repo Repository, logger log.Logger) *myStringService {
	return &myStringService{repository: repo, log: logger}
}

func (svc *myStringService) IsPal(ctx context.Context, s string) string {
	reverse := svc.Reverse(ctx, s)
	if strings.ToLower(s) != reverse {
		return "Is not palindrome"
	}
	return "Is palindrome"
}

func (svc *myStringService) Reverse(ctx context.Context, s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return strings.ToLower(string(rns))
}
