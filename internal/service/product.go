package service

import (
	"context"
	"fmt"

	"github.com/ysomad/go-project-structure/internal/domain"
	"github.com/ysomad/go-project-structure/internal/paging"
	"github.com/ysomad/go-project-structure/internal/storage/postgres"
)

type Product struct {
	storage *postgres.ProductStorage
}

func NewProduct(s *postgres.ProductStorage) *Product {
	return &Product{storage: s}
}

func (p *Product) List(ctx context.Context, f *domain.ProductFilters, s domain.ProductSort, tp paging.TokenParams) (domain.ProductList, error) {
	products, err := p.storage.GetAll(ctx, f, s, tp)
	if err != nil {
		return domain.ProductList{}, fmt.Errorf("storage: %w", err)
	}
	return products, nil
}
