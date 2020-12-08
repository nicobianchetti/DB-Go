package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items (
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_items_invoice_header_id_fx FOREIGN KEY
		(invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_invoice_product_id_fx FOREIGN KEY
		(invoice_header_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
)

//PsqlInoviceItem used for work with postgres - invoiceHeader
type PsqlInoviceItem struct {
	db *sql.DB
}

//NewPsqlInoviceItem return a new pointer of PsqlInoviceItem
func NewPsqlInoviceItem(db *sql.DB) *PsqlInoviceItem {
	return &PsqlInoviceItem{db}
}

//Migrate implement the interface invoiceHeader.Storage
func (p *PsqlInoviceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
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
