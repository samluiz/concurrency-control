package handlers

import "github.com/samluiz/concurrency-control/internal/db/repositories"

type Handler struct {
	repo *repositories.Repo
}

func NewHandler(repo *repositories.Repo) *Handler {
	return &Handler{repo}
}