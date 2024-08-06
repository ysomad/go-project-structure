package postgres

import (
	"context"
	"errors"

	"github.com/ysomad/go-project-structure/internal/domain"
	"github.com/ysomad/go-project-structure/internal/paging"
	"github.com/ysomad/go-project-structure/internal/storage/postgres/pgclient"
)

type ProductStorage struct {
	cli *pgclient.Client
}

func NewProductStorage(c *pgclient.Client) *ProductStorage {
	return &ProductStorage{cli: c}
}

func (s *ProductStorage) GetAll(context.Context, *domain.ProductFilters, domain.ProductSort, paging.TokenParams) (domain.ProductList, error) {
	return domain.ProductList{}, errors.New("not implemented")
}
