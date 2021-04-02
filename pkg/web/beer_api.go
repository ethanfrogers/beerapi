package web

import (
	"github.com/ethanfrogers/beerapi/pkg/repository"
	"github.com/gorilla/mux"
)

type beerAPI struct {
	repo repository.BeerRepository
}

func NewBeerAPI(repo repository.BeerRepository) (*beerAPI, error) {
	api := &beerAPI{repo: repo}
	return api, nil
}

func (ba *beerAPI) RegisterHandlers(router *mux.Router) {
	// TODO: register handlers
}
