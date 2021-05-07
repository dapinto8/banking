package service

import (
	"github.com/dapinto8/banking/domain"
	"github.com/dapinto8/banking/dto"
	"github.com/dapinto8/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repository.FindAll(status)
	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repository.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
