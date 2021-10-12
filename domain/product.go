package domain

import (
	"context"
	"time"
)

type Product struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	CategoryID int       `json:"category_id"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Category Category `json:"category" autoload:"true" ref:"category_id" fk:"id"`
}

type ProductUsecase interface {
	Search(ctx context.Context, name string, categoryName string) ([]Product, error)
	FetchByID(ctx context.Context, id int) (Product, error)
	Create(ctx context.Context, product Product) (Product, error)
	Update(ctx context.Context, product Product) (Product, error)
	Delete(ctx context.Context, id int) error
}
