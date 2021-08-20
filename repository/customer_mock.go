package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{
			CustomerID:  1001,
			Name:        "Ashish",
			City:        "New Delhi",
			ZipCode:     "110011",
			DateOfBirth: "2000",
		},
		{
			CustomerID:  1002,
			Name:        "Rob",
			City:        "New Delhi",
			ZipCode:     "110011",
			DateOfBirth: "2000",
		},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}