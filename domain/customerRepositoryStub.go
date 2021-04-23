package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Mark Prince", "Miami", "33199", "2000-04-19", "1"},
		{"1002", "Rob Ayish", "New Delhi", "1100011", "2000-01-03", "1"},
	}

	return CustomerRepositoryStub{customers}
}
