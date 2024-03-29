package repository

import (
	"TestKasir/infras"
)

type AuthRepository interface {
	AuthManagementRepository
}

type AuthRepositoryPostgres struct {
	DB *infras.PostgreSQLConn
}

func ProvideAuthRepositoryPostgres(db *infras.PostgreSQLConn) *AuthRepositoryPostgres {
	return &AuthRepositoryPostgres{
		DB: db,
	}
}
