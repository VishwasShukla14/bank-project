package domain

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepository {
	customers := []Customer{
		{"101", "VIshwas", "Ahmedabd", "380014", "14/05/2002", "Active"},
	}
	return CustomerRepositoryStub{customers}
}
