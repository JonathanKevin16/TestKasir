package repository

import (
	"TestKasir/infras"
)

type ProductRepository interface {
	ProductManagementRepository
}

type ProductRepositoryPostgres struct {
	DB *infras.PostgreSQLConn
}

func ProvideProductRepositoryPostgres(db *infras.PostgreSQLConn) *ProductRepositoryPostgres {
	return &ProductRepositoryPostgres{
		DB: db,
	}
}
