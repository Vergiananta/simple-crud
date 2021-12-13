package manager

import (
	"database/sql"
	"learn-go/connect"
	"learn-go/repository"
)

type RepoManager interface {
	CustomerRepo() repository.ICustomerRepo
}

type repoManager struct {
	db      *sql.DB
	connect connect.Connect
}

func (r *repoManager) CustomerRepo() repository.ICustomerRepo {
	//TODO implement me
	return repository.NewCustomerRepostory(r.db)
}

func NewRepoManager(connect connect.Connect) RepoManager {
	return &repoManager{connect.SqlDb(), connect}
}
