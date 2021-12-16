package usecase

import (
	"crud-go/models"
	"crud-go/repository"
)

type ICustomerUsecase interface {
	CreateCustomer(newCustomer *models.Customer) (*models.Customer, error)
	UpdateCustomer(nama string, id string) (string, error)
	DeleteCustomer(id string) (string, error)
	GetAllCustomer() ([]models.Customer, error)
}

type customerUsecase struct {
	customerRepo repository.ICustomerRepo
}

func (c *customerUsecase) GetAllCustomer() ([]models.Customer, error) {
	//TODO implement me
	return c.customerRepo.GetAllCustomer()
}

func (c *customerUsecase) DeleteCustomer(id string) (string, error) {
	//TODO implement me
	return c.customerRepo.DeleteCustomer(id)
}

func (c *customerUsecase) UpdateCustomer(nama string, id string) (string, error) {
	//TODO implement me
	return c.customerRepo.UpdateName(nama, id)
}

func (c *customerUsecase) CreateCustomer(newCustomer *models.Customer) (*models.Customer, error) {
	return c.customerRepo.CreateCustomer(newCustomer)
}

func NewCustomerUsecase(custRepo repository.ICustomerRepo) ICustomerUsecase {
	return &customerUsecase{custRepo}
}
