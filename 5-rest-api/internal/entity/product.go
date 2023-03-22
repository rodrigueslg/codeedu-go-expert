package entity

import (
	"errors"
	"time"

	"github.com/rodrigueslg/codedu-goexpert/rest-api/pkg/entity"
)

var (
	ErrRequiredID = errors.New("id is required")
	ErrInvalidID  = errors.New("invalid id")

	ErrRequiredName = errors.New("name is required")

	ErrRequiredPrice = errors.New("price is required")
	ErrInvalidPrice  = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	p := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrRequiredID
	}
	if _, err := entity.ParseID(string(p.ID.String())); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrRequiredName
	}

	if p.Price == 0 {
		return ErrRequiredPrice
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
