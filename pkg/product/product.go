package product

import "time"

//Product is Model
type Product struct {
	ID          uint
	Name        string
	Observation string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Products []*Product

type Storage interface {
	Migrate() error
	Create(*Product) error
}

//Service of product
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
