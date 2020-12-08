package main

import (
	"log"

	"github.com/nicobianchetti/DB-Go/pkg/invoiceHeader"
	"github.com/nicobianchetti/DB-Go/pkg/invoiceItem"
	"github.com/nicobianchetti/DB-Go/pkg/product"
	"github.com/nicobianchetti/DB-Go/storage"
)

func main() {
	//Creación de tablas de base de datos
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())

	//a newService le paso un psqlProduct pero podria pasarle tamb un MySqlProduct,etc
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	storageInvoice := storage.NewPsqlInoviceHeader(storage.Pool())

	serviceInovice := invoiceHeader.NewService(storageInvoice)

	if err := serviceInovice.Migrate(); err != nil {
		log.Fatalf("invoice.Migrate: %v", err)
	}

	storageInvoiceItem := storage.NewPsqlInoviceItem(storage.Pool())

	serviceInvoiceItem := invoiceItem.NewService(storageInvoiceItem)

	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceItem.Migrate: %v", err)
	}

}
