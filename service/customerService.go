package service

import (
	"github.com/dapinto8/banking/domain"
	"github.com/dapinto8/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *errs.AppError)
	GetCustomer(id string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	return s.repository.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repository.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
