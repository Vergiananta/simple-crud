package manager

import (
	"learn-go/connect"
	"learn-go/usecase"
)

type UsecaseManager interface {
	CustomerUsecase() usecase.ICustomerUsecase
}

type usecaseManager struct {
	repo RepoManager
}

func (u *usecaseManager) CustomerUsecase() usecase.ICustomerUsecase {
	//TODO implement me
	return usecase.NewCustomerUsecase(u.repo.CustomerRepo())
}

func NewUsecaseManager(connect connect.Connect) UsecaseManager {
	return &usecaseManager{NewRepoManager(connect)}
}
