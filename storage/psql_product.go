package storage

import (
	"database/sql"
	"fmt"

	"github.com/nicobianchetti/DB-Go/pkg/product"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products (
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`

	psqlCreateProduct = `INSERT INTO products(name,observations,price,created_at) VALUES($1,$2,$3,$4) RETURNING id`
)

//PsqlProduct used for work with postgres - product
type PsqlProduct struct {
	db *sql.DB
}

//NewPsqlProduct return a new pointer of PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

//Migrate implement the interface produc.Storage
func (pr *PsqlProduct) Migrate() error {
	stmt, err := pr.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de producto ejecutada correctamente")
	return nil
}

//Create implement the interface product.Storage
func (pr *PsqlProduct) Create(p *product.Product) error {
	return nil
}
