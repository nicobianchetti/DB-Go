package invoiceHeader

import "time"

//Model of InvoiceHeader
type InvoiceHeader struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Storage interface {
	Migrate() error
}

//Service of InvoiceHeader
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
