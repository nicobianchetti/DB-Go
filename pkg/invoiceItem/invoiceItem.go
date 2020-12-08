package invoiceItem

import (
	"time"
)

//Model of InvoiceItem
type InvoiceItem struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UodatedAt       time.Time
}

type Storage interface {
	Migrate() error
}

//Service of invoiceItem
type Service struct {
	storage Storage
}

//NewService return a ppointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

//Migrate is udes for migrate product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
