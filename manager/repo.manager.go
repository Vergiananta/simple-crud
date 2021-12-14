package manager

import (
	"crud-go/connect"
	"crud-go/repository"
	"database/sql"
)

type RepoManager interface {
	CustomerRepo() repository.ICustomerRepo
}

type repoManager struct {
	db      *sql.DB
	connect connect.Connect
}

func (r *repoManager) CustomerRepo() repository.ICustomerRepo {
	return repository.NewCustomerRepository(r.db)
}

func NewRepoManager(connect connect.Connect) RepoManager {
	return &repoManager{connect.SqlDb(), connect}
}
