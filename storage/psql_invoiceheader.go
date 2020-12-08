package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers (
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
	)`
)

//PsqlInoviceHeader used for work with postgres - invoiceHeader
type PsqlInoviceHeader struct {
	db *sql.DB
}

//NewPsqlInoviceHeader return a new pointer of PsqlInoviceHeader
func NewPsqlInoviceHeader(db *sql.DB) *PsqlInoviceHeader {
	return &PsqlInoviceHeader{db}
}

//Migrate implement the interface invoiceHeader.Storage
func (p *PsqlInoviceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de invoiceHeader ejecutada correctamente")
	return nil
}
