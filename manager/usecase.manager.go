package manager

import (
	"crud-go/connect"
	"crud-go/usecase"
)

type UsecaseManager interface {
	CustomerUsecase() usecase.ICustomerUsecase
}

type usecaseManager struct {
	repo RepoManager
}

func (u *usecaseManager) CustomerUsecase() usecase.ICustomerUsecase {
	return usecase.NewCustomerUsecase(u.repo.CustomerRepo())
}

func NewUsecasManager(connect connect.Connect) UsecaseManager {
	return &usecaseManager{NewRepoManager(connect)}
}
