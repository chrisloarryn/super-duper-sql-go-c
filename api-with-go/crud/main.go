package main

import (
	"GO_MINE/crud_psql/pkg/product"
	"GO_MINE/crud_psql/storage"
	"fmt"
	"log"
)

func main() {
	fmt.Println("EDteam")
	//storage.NewPostgresDB()

	driver := storage.MySQL
	storage.New(driver)

	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %+v", err)
	}

	serviceProduct := product.NewService(myStorage)
	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("serviceProduct.GetAll: %+v", err)
	}

	fmt.Printf("%+v\n", ms)
	//
	//storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	//storageItems := storage.NewMySQLInvoiceItem(storage.Pool())
	//storageInvoice := storage.NewMySQLInvoice(
	//	storage.Pool(),
	//	storageHeader,
	//	storageItems,
	//)
	//
	//m := &invoice.Model{
	//	Header: &invoiceheader.Model{
	//		Client: "Alexys",
	//	},
	//	Items: invoiceitem.Models{
	//		&invoiceitem.Model{ProductID: 1},
	//		&invoiceitem.Model{ProductID: 3},
	//		&invoiceitem.Model{ProductID: 23},
	//	},
	//}
	//
	//serviceInvoice := invoice.NewService(storageInvoice)
	//if err := serviceInvoice.Create(m); err != nil {
	//	log.Fatalf("invoice.Create: %v", err)
	//}

	//storageProduct := storage.NewMySQLProduct(storage.Pool())
	//serviceProduct := product.NewService(storageProduct)
	//
	//err := serviceProduct.Delete(4)
	//if err != nil {
	//	log.Fatalf("product.Delete: %v", err)
	//}
	//
	//err = serviceProduct.Delete(5)
	//if err != nil {
	//	log.Fatalf("product.Delete: %v", err)
	//}
	//
	//err = serviceProduct.Delete(6)
	//if err != nil {
	//	log.Fatalf("product.Delete: %v", err)
	//}

	//storageProduct := storage.NewMySQLProduct(storage.Pool())
	//serviceProduct := product.NewService(storageProduct)
	//
	//m := &product.Model{
	//	ID:    2,
	//	Name:  "Curso de css",
	//	Price: 250,
	//}
	//err := serviceProduct.Update(m)
	//if err != nil {
	//	log.Fatalf("product.Update: %v", err)
	//}

	//storageProduct := storage.NewMySQLProduct(storage.Pool())
	//serviceProduct := product.NewService(storageProduct)
	//
	//m, err := serviceProduct.GetByID(3)
	//switch {
	//case errors.Is(err, sql.ErrNoRows):
	//	fmt.Println("no hay un producto con este id")
	//case err != nil:
	//	log.Fatalf("product.GetByID: %v", err)
	//default:
	//	fmt.Println(m)
	//}

	//storageProduct := storage.NewMySQLProduct(storage.Pool())
	//serviceProduct := product.NewService(storageProduct)
	//
	//ms, err := serviceProduct.GetAll()
	//if err != nil {
	//	log.Fatalf("product.GetAll: %v", err)
	//}
	//
	//fmt.Println(ms)

	//storageProduct := storage.NewMySQLProduct(storage.Pool())
	//serviceProduct := product.NewService(storageProduct)
	//
	//m := &product.Model{
	//	Name:  "Curso de testing con Go",
	//	Price: 70,
	//	//Observations: "on fire",
	//}
	//if err := serviceProduct.Create(m); err != nil {
	//	log.Fatalf("product.Create: %v", err)
	//}
	//
	//fmt.Printf("%+v\n", m)

	//storageProduct := storage.NewMySQLProduct(storage.Pool())
	//serviceProduct := product.NewService(storageProduct)
	//
	//if err := serviceProduct.Migrate(); err != nil {
	//	fmt.Println(err)
	//	log.Fatalf("Migration failed: product.Migrate %v", err)
	//}
	//
	//storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	//serviceHeader := invoiceheader.NewService(storageHeader)
	//
	//if err := serviceHeader.Migrate(); err != nil {
	//	fmt.Println(err)
	//	log.Fatalf("Migration failed: invoiceheader.Migrate %v", err)
	//}
	//
	//storageItem := storage.NewMySQLInvoiceItem(storage.Pool())
	//serviceItem := invoiceitem.NewService(storageItem)
	//
	//if err := serviceItem.Migrate(); err != nil {
	//	fmt.Println(err)
	//	log.Fatalf("Migration failed: invoiceitem.Migrate %v", err)
	//}

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
