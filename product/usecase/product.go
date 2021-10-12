package usecase

import (
	"BEtest-simpleAPI/domain"
	"context"
	"fmt"

	"github.com/go-rel/rel"
)

type productUsecase struct {
	repo rel.Repository
}

func New(repo rel.Repository) domain.ProductUsecase {
	return &productUsecase{repo}
}

func (u *productUsecase) Search(ctx context.Context, name string, categoryName string) ([]domain.Product, error) {
	var product []domain.Product
	var err error

	if name != "" {
		product, err = u.fetchByName(ctx, name)
		if err != nil {
			return product, err
		}
	} else if categoryName != "" {
		product, err = u.fetchByCategoryName(ctx, categoryName)
		if err != nil {
			return product, err
		}
	} else {
		product, err = u.fetch(ctx)
		if err != nil {
			return product, err
		}
	}

	return product, nil
}

func (u *productUsecase) FetchByID(ctx context.Context, id int) (domain.Product, error) {
	var product domain.Product
	err := u.repo.Find(ctx, &product, rel.Eq("id", id))
	return product, err
}

func (u *productUsecase) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	err := u.repo.Insert(ctx, &product)
	return product, err
}

func (u *productUsecase) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	updatedProduct, err := u.FetchByID(ctx, product.ID)
	if err != nil {
		return product, err
	}

	if product.Name != "" {
		updatedProduct.Name = product.Name
	}

	if product.Price != 0 {
		updatedProduct.Price = product.Price
	}

	if product.CategoryID != 0 {
		updatedProduct.CategoryID = product.CategoryID
	}

	if product.Image != "" {
		updatedProduct.Image = product.Image
	}

	product.Category = domain.Category{}
	err = u.repo.Update(ctx, &updatedProduct)
	return updatedProduct, err
}

func (u *productUsecase) Delete(ctx context.Context, id int) error {
	deletedProduct, err := u.FetchByID(ctx, id)
	if err != nil {
		return err
	}

	err = u.repo.Delete(ctx, &deletedProduct)
	return err
}

func (u *productUsecase) fetch(ctx context.Context) ([]domain.Product, error) {
	var product []domain.Product
	err := u.repo.FindAll(ctx, &product)
	return product, err
}

func (u *productUsecase) fetchByName(ctx context.Context, name string) ([]domain.Product, error) {
	var product []domain.Product
	pattern := fmt.Sprintf("%%%s%%", name)
	err := u.repo.FindAll(ctx, &product, rel.Like("name", pattern))
	return product, err
}

func (u *productUsecase) fetchByCategoryName(ctx context.Context, name string) ([]domain.Product, error) {
	var product []domain.Product
	var category domain.Category
	pattern := fmt.Sprintf("%%%s%%", name)
	err := u.repo.Find(ctx, &category, rel.Like("name", pattern))
	if err != nil {
		return product, err
	}

	err = u.repo.FindAll(ctx, &product, rel.Eq("category_id", category.ID))
	return product, err
}
