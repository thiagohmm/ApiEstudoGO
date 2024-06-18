package entity

import (
	"errors"
	"time"

	"github.com/thiagohmm/ApiGo/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrInvalidId       = errors.New("invalid ID")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid Price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {

	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := ValidateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil

}

func ValidateProduct(product *Product) error {
	if _, err := entity.ParseID(product.ID.String()); err != nil {
		return ErrInvalidId
	}

	if product.ID.String() == "" {
		return ErrIDIsRequired
	}

	if product.Name == "" {
		return ErrNameIsRequired
	}
	if product.Price == 0 {
		return ErrPriceIsRequired
	}
	if product.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
