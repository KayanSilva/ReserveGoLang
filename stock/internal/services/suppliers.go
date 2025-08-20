package services

import "fmt"

type InfoService interface {
	GetInfo() string
}

type AvailableService interface {
	CheckAvailability(quantityRequested int, quantityAvailable int) bool
}

type SupplierService interface {
	InfoService
	AvailableService
}

type Supplier struct {
	CNPJ    string
	Contact string
	City    string
}

func (s Supplier) GetInfo() string {
	return fmt.Sprintf("CNPJ: %s | Contact: %s | City: %s",
		s.CNPJ, s.Contact, s.City)
}

func (s Supplier) CheckAvailability(quantityRequested int, quantityAvailable int) bool {
	return quantityRequested <= quantityAvailable
}
