package usecase

import (
	"learn-go/models"
	"learn-go/repository"
)

type ICustomerUsecase interface {
	CreateCustomer(newCustomer *models.Customer) (*models.Customer, error)
}

type customerUsecase struct {
	customerRepo repository.ICustomerRepo
}

func (c *customerUsecase) CreateCustomer(newCustomer *models.Customer) (*models.Customer, error) {
	//TODO implement me
	return c.customerRepo.CreateCustomer(newCustomer)
}

func NewCustomerUsecase(custRepo repository.ICustomerRepo) ICustomerUsecase {
	return &customerUsecase{custRepo}
}
