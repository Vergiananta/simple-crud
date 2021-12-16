package handlers

import (
	"crud-go/controllers"
	"crud-go/models"
	"crud-go/usecase"
	"crud-go/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomerController struct {
	router   *gin.Engine
	response response.IResponder
	service  usecase.ICustomerUsecase
}

func (c *CustomerController) InitRoutes() {
	//TODO implement me
	u := c.router.Group("/customers")
	u.POST("", c.CreateCustomer)
	u.PATCH("/:id", c.UpdateNama)
	u.DELETE("/:id", c.DeleteCustomer)
	u.GET("", c.GetAllCustomer)
}

func NewCustomerController(router *gin.Engine, responder response.IResponder, service usecase.ICustomerUsecase) controllers.IDelivery {
	return &CustomerController{router, responder, service}
}

func (cu *CustomerController) CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		cu.response.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
		return
	}

	newCust, err := cu.service.CreateCustomer(&customer)
	if err != nil {
		cu.response.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
		return
	}
	cu.response.SetContext(c).SingleResponder(http.StatusCreated, newCust)

	fmt.Println(newCust)
}

func (cu *CustomerController) UpdateNama(c *gin.Context) {
	param := c.Param("id")
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		cu.response.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
		return
	}

	updateCust, err := cu.service.UpdateCustomer(customer.Nama, param)
	if err != nil {
		cu.response.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
		return
	}
	cu.response.SetContext(c).SingleResponder(http.StatusCreated, updateCust)
}

func (cu *CustomerController) DeleteCustomer(c *gin.Context) {
	param := c.Param("id")
	delCust, err := cu.service.DeleteCustomer(param)
	if err != nil {
		cu.response.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
		return
	}
	cu.response.SetContext(c).SingleResponder(http.StatusCreated, delCust)
}

func (cu *CustomerController) GetAllCustomer(c *gin.Context) {
	var cust []models.Customer
	var err error
	cust, err = cu.service.GetAllCustomer()
	if err != nil {
		cu.response.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
		return
	}
	cu.response.SetContext(c).SingleResponder(http.StatusCreated, cust)
}
