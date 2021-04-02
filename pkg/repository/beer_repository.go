package repository

// Beer models different beer types
type Beer struct {
	Alcohol     float64 `json:"alcohol"`
	Description string  `json:"description"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
}

// BeerRepository defines an interface for working with beer data
type BeerRepository interface {
	ListBeers() ([]Beer, error)
	GetBeer(id string) (Beer, error)
}

type defaultBeerRepository struct {
	data map[string]Beer
}

func NewBeerRepository() (*defaultBeerRepository, error) {
	// TODO: implement filling out of data
	return &defaultBeerRepository{}, nil
}
