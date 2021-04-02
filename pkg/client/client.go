package client

import "github.com/ethanfrogers/beerapi/pkg/repository"

type BeerAPIClient interface {
	ListBeers() ([]repository.Beer, error)
	GetBeer(id string) (repository.Beer, error)
}

type defaultBeerAPIClient struct {
	baseURL string
}

func NewBeerAPIClient(baseURL string) (*defaultBeerAPIClient, error) {
	return &defaultBeerAPIClient{baseURL: baseURL}, nil
}
