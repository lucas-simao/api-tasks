package tasks

import (
	"github.com/lucas-simao/api-tasks/internal/repository"
)

type service struct {
	repository repository.Repository
}

func New(r repository.Repository) Service {
	return service{
		repository: r,
	}
}
