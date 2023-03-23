package hasmany

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Category struct {
		ID       int `gorm:primary_key`
		Name     string
		Products []Product
	}

	Product struct {
		ID           int `gorm:primary_key`
		Name         string
		Price        float64
		CategoryID   int
		Category     Category
		SerialNumber SerialNumber
		gorm.Model
	}

	SerialNumber struct {
		ID        int `gorm:primary_key`
		Number    string
		ProductID int
	}
)

func RunHasManySample() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	fmt.Println("\ncreating single category")
	cEletronicos := Category{Name: "Eletrônicos"}
	db.Create(&cEletronicos)

	fmt.Println("\ncreating single product")
	var pNotebook Product = Product{
		Name:       "Notebook",
		Price:      1000.0,
		CategoryID: cEletronicos.ID,
	}
	db.Create(&pNotebook)

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: pNotebook.ID,
	})

	fmt.Println("\ncreating single category")
	cCozinha := Category{Name: "Cozinha"}
	db.Create(&cCozinha)

	fmt.Println("\ncreating single product")
	var pPanela Product = Product{
		Name:       "Panela",
		Price:      48.0,
		CategoryID: cCozinha.ID,
	}
	db.Create(&pPanela)

	db.Create(&SerialNumber{
		Number:    "654321",
		ProductID: pPanela.ID,
	})

	fmt.Println("\n selecting products from a category")
	var categories []Category
	//err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("-", product.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}
}
