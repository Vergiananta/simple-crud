package repository

import (
	"database/sql"
	"learn-go/models"
)

const (
	INSERTNEW = `insert into customers (nama,umur,address) values ($1, $2, $3)`
)

type ICustomerRepo interface {
	CreateCustomer(newCustomer *models.Customer) (*models.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) CreateCustomer(newCustomer *models.Customer) (*models.Customer, error) {
	db, err := c.db.Begin()
	_, err = db.Exec(INSERTNEW, newCustomer.Nama, newCustomer.Umur, newCustomer.Address)
	if err != nil {
		return nil, err
	}
	err = db.Commit()
	if err != nil {
		return nil, err
	}
	return newCustomer, nil
}

func NewCustomerRepostory(db *sql.DB) ICustomerRepo {
	return &customerRepository{db}
}
