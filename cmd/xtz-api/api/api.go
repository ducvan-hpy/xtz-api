package api

import (
	"github.com/gin-gonic/gin"

	"github.com/ducvan-hpy/xtz-api/internal/domain/repository"
)

type Api struct {
	repository *repository.Repository
}

func New(repo *repository.Repository) *Api {
	return &Api{
		repository: repo,
	}
}

func NewGinRouter(a *Api) *gin.Engine {
	r := gin.New()
	r.Use(logRequest())
	RegisterHandlers(r, a)
	return r
}
