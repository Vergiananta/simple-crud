package handlers

import (
	"crud-go/controllers"
	"crud-go/models"
	"crud-go/usecase"
	"fmt"
)

type CustomerController struct {
	service usecase.ICustomerUsecase
}

func (c *CustomerController) InitRoutes() {
	//TODO implement me
	//c.CreateCustomer()
	//c.UpdateNama()
	//c.DeleteCustomer()
	c.GetAllCustomer()
}

func NewCustomerController(service usecase.ICustomerUsecase) controllers.IDelivery {
	return &CustomerController{service}
}

func (c *CustomerController) CreateCustomer() {
	customers := models.Customer{
		Nama:    "alfika",
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

func (c *CustomerController) UpdateNama() {
	nama := "fika"
	id := 1
	updateCust, err := c.service.UpdateCustomer(nama, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(updateCust)
}

func (c *CustomerController) DeleteCustomer() {
	id := 8
	delCust, err := c.service.DeleteCustomer(id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(delCust)
}

func (c *CustomerController) GetAllCustomer() {
	var cust []models.Customer
	var err error
	cust, err = c.service.GetAllCustomer()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(&cust)
}
