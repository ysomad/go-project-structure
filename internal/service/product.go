package service

import (
	"context"

	"github.com/ysomad/go-project-structure/internal/domain"
	"github.com/ysomad/go-project-structure/internal/paging"
)

type Product struct {
	repo productRepository
}

func NewProduct(r productRepository) *Product {
	return &Product{repo: r}
}

type productRepository interface {
	GetAll(context.Context, domain.ProductFilters, domain.ProductSort, paging.TokenParams) (domain.ProductList, error)
}

func (p *Product) List(ctx context.Context, f *domain.ProductFilters, s domain.ProductSort, tp paging.TokenParams) (domain.ProductList, error) {
	return domain.ProductList{}, nil
}
