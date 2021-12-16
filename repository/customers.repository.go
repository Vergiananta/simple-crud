package repository

import (
	"crud-go/models"
	"database/sql"
	"errors"
)

const (
	INSERTNEW      = `insert into customers (nama,umur,address) values ($1, $2, $3)`
	UPDATECUSTOMER = `update customers set nama = $1 where id = $2`
	DELETECUSTOMER = `delete from customers where id = $1`
	GETALLCUSTOMER = `select * from customers`
)

type ICustomerRepo interface {
	CreateCustomer(newCustomer *models.Customer) (*models.Customer, error)
	UpdateName(nama string, id string) (string, error)
	DeleteCustomer(id string) (string, error)
	GetAllCustomer() ([]models.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) GetAllCustomer() ([]models.Customer, error) {
	//TODO implement me
	db, err := c.db.Begin()
	var result []models.Customer
	rows, err := db.Query("SELECT c.id,c.nama,c.umur,c.address FROM customers as c")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var cust = models.Customer{}
		err = rows.Scan(&cust.Id, &cust.Nama, &cust.Umur, &cust.Address)
		if err != nil {
			return nil, err
		}
		result = append(result, cust)
	}
	return result, nil
}

func (c *customerRepository) DeleteCustomer(id string) (string, error) {
	//TODO implement me
	db, err := c.db.Begin()
	_, err = db.Query(DELETECUSTOMER, id)
	if err != nil {
		db.Rollback()
		return "", errors.New("Data Tidak ditemukan")
	}
	err = db.Commit()
	if err != nil {
		db.Rollback()
		return "", err
	}
	return "Data has been deleted", nil
}

func (c *customerRepository) UpdateName(nama string, id string) (string, error) {
	//TODO implement me
	db, err := c.db.Begin()
	_, err = db.Exec(UPDATECUSTOMER, nama, id)
	if err != nil {
		db.Rollback()
		return "", err
	}
	err = db.Commit()
	if err != nil {
		return "", err
	}
	return "success", nil
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

func NewCustomerRepository(db *sql.DB) ICustomerRepo {
	return &customerRepository{db}
}
