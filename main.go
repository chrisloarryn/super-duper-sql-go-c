package main

import (
	"GO_MINE/crud_psql/storage"
	"fmt"
)

func main() {
	fmt.Println("EDteam")
	//storage.NewPostgresDB()
	storage.NewMySQLDB()
	storage.NewMySQLDB()
	storage.NewMySQLDB()
	storage.NewMySQLDB()

	//storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	//storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	//
	//storageInvoice := storage.NewPsqlInvoice(
	//	storage.Pool(),
	//	storageHeader,
	//	storageItems,
	//)
	//
	//serviceInvoice := invoice.NewService(storageInvoice)
	//
	//m := &invoice.Model{
	//	Header: &invoiceheader.Model{
	//		Client: "Alvaro Felipe",
	//	},
	//	Items: invoiceitem.Models{
	//		&invoiceitem.Model{
	//			ProductID: 1,
	//		},
	//		&invoiceitem.Model{
	//			ProductID: 4,
	//		},
	//	},
	//}
	//
	//if err := serviceInvoice.Create(m); err != nil {
	//	log.Fatalf("invoice.Create: %v", err)
	//}

	//storageProduct := storage.NewPsqlProduct(storage.Pool())
	//serviceProduct := product.NewService(storageProduct)
	//
	//m := &product.Model{
	//	Name:         "Curso DB con GO y Postgres",
	//	Observations: "",
	//	Price:        380,
	//}
	//
	//fmt.Printf("zcxzcx %+v\n", m)
	//
	//ms, err := serviceProduct.GetAll()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("zcxzcx %+v\n", ms)
	//
	//mm, err := serviceProduct.GetByID(4)
	//switch {
	//case errors.Is(err, sql.ErrNoRows):
	//	fmt.Println("No se encontro el producto")
	//case err != nil:
	//	log.Fatalf("product.GetByID: %v", err)
	//default:
	//	fmt.Printf("%+v\n", mm)
	//}
	//
	//// copy mm
	//toUpdate := *mm
	//toUpdate.ID = 90
	//toUpdate.Name = "Curso de Comer popo"
	//toUpdate.Price = 500
	//
	//err = serviceProduct.Update(&toUpdate)
	//if err != nil {
	//	fmt.Println("cannot update ", err)
	//}
	//
	//mm, err = serviceProduct.GetByID(4)
	//switch {
	//case errors.Is(err, sql.ErrNoRows):
	//	fmt.Println("No se encontro el producto")
	//case err != nil:
	//	fmt.Println("product.GetByID: %v", err)
	//default:
	//	fmt.Printf("%+v\n", mm)
	//}
	//
	//mm, err = serviceProduct.GetByID(5)
	//switch {
	//case errors.Is(err, sql.ErrNoRows):
	//	fmt.Println("No se encontro el producto")
	//case err != nil:
	//	fmt.Println("product.GetByID: %v", err)
	//default:
	//	fmt.Printf("%+v\n", mm)
	//}
	//
	//err = serviceProduct.Delete(5)
	//if err != nil {
	//	fmt.Println("could not delete ", err)
	//}

	//storageProduct := storage.NewPsqlProduct(storage.Pool())
	//serviceProduct := product.NewService(storageProduct)
	//
	//if err := serviceProduct.Migrate(); err != nil {
	//	log.Fatalf("product.Migrate: %v", err)
	//}
	//
	//storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	//serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
	//
	//if err := serviceInvoiceHeader.Migrate(); err != nil {
	//	log.Fatalf("invoiceHeader.Migrate: %v", err)
	//}
	//
	//storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	//serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	//
	//if err := serviceInvoiceItem.Migrate(); err != nil {
	//	log.Fatalf("invoiceItem.Migrate: %v", err)
	//}

}
