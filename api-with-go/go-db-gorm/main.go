package main

import (
	"fmt"
	"go-db-gorm/model"
	"go-db-gorm/storage"
)

func main() {
	driver := storage.MySQL

	storage.New(driver)

	//invoice := model.InvoiceHeader{
	//	Client: "Alvaro Felipe",
	//	InvoiceItems: []model.InvoiceItem{
	//		{ProductID: 3},
	//	},
	//}
	//
	//storage.DB().Create(&invoice)

	// AutoMigrate models

	//productOne := model.Product{
	//	Name:  "Curso de Go",
	//	Price: 120,
	//}
	//productTwo := model.Product{
	//	Name:         "Curso de Testing",
	//	Price:        150,
	//	Observations: StringPointer("testing con go"),
	//}
	//productThree := model.Product{
	//	Name:  "Curso de Python",
	//	Price: 200,
	//}
	//storage.DB().Create(&productOne)
	//storage.DB().Create(&productTwo)
	//storage.DB().Create(&productThree)

	//storageProduct := storage.NewPsqlProduct(storage.Pool())
	//serviceProduct := product.NewService(storageProduct)
	//
	//if err := serviceProduct.Migrate(); err != nil {
	//	log.Fatalf("product.Migrate: %v", err)
	//}

	//storageProduct := storage.New(storage.MySQL)
	//_ = storageProduct.AutoMigrate(
	//	//&model.Product{},
	//	&model.InvoiceHeader{},
	//	&model.InvoiceItem{},
	//)

	products := make([]model.InvoiceHeader, 0)
	storage.DB().Find(&products)

	for _, product := range products {
		fmt.Printf("%d - %s\n", product.ID, product.Client)
	}

	//myProduct := model.Product{}
	//
	//storage.DB().First(&myProduct, 3)
	//
	//storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
	//
	//storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("invoice_header_id", "invoice_headers(id)", "RESTRICT", "RESTRICT")

	invoice := model.InvoiceHeader{
		Client: "Alvaro Felipe",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 3},
		},
	}

	storage.DB().Create(&invoice)

	//myProduct = model.Product{
	//	Name: "prod 2",
	//}
	//myProduct.ID = 2
	//
	//storage.DB().Model(&myProduct).Updates(
	//	model.Product{Name: "Curso de Java", Price: 120},
	//)

	//ivh := model.InvoiceHeader{
	//	Client: "Aracely Paz",
	//	InvoiceItems: []model.InvoiceItem{
	//		{
	//			InvoiceHeaderID: 0,
	//			ProductID:       0,
	//		},
	//	},
	//}
	//
	//storage.DB().Create(&ivh)
	//
	//prod1 := model.Product{
	//	Name:         "prod 1",
	//	Observations: StringPointer("asdda"),
	//	Price:        1000,
	//	InvoiceItems: []model.InvoiceItem{
	//		{
	//			InvoiceHeaderID: 1,
	//			ProductID:       2,
	//		},
	//	},
	//}
	//storage.DB().Create(&prod1)
	//
	//prod2 := model.Product{
	//	Name:         "prod 2",
	//	Observations: StringPointer("pppp"),
	//	Price:        1500,
	//	InvoiceItems: nil,
	//}
	//storage.DB().Create(&prod2)
	//
	//invoice := model.InvoiceHeader{
	//	Client: "Alvaro Felipe",
	//	InvoiceItems: []model.InvoiceItem{
	//		{ProductID: 4},
	//	},
	//}
	//
	//storage.DB().Create(&invoice)

	//serviceProduct.

	//storage.DB().AutoMigrate(
	//	&model.Product{},
	//	&model.InvoiceHeader{},
	//	&model.InvoiceItem{},
	//)
}

// StringPointer func string to *string
func StringPointer(s string) *string {
	return &s
}
