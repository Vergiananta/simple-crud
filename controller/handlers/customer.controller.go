package handlers

import (
	"fmt"
	"learn-go/controller"
	"learn-go/models"
	"learn-go/usecase"
)

type CustomerController struct {
	service usecase.ICustomerUsecase
}

func (c *CustomerController) InitRoutes() {
	//TODO implement me
	c.CreateCustomer()
}

func NewCustomerController(service usecase.ICustomerUsecase) controller.IDelivery {
	return &CustomerController{service}
}

func (c *CustomerController) CreateCustomer() {
	customers := models.Customer{
		Nama:    "fakih",
		Address: "senayan",
		Umur:    21,
	}
	newCust, err := c.service.CreateCustomer(&customers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(newCust)
}
