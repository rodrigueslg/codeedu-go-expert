package hasone

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Category struct {
		ID   int `gorm:primary_key`
		Name string
	}

	SerialNumber struct {
		ID        int `gorm:primary_key`
		Number    string
		ProductID int
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
)

func RunHasOneSample() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	fmt.Println("\ncreating single category")
	category := Category{Name: "Eletrônicos"}
	db.Create(&category)

	fmt.Println("\ncreating single product")
	var p Product = Product{
		Name:       "Notebook",
		Price:      1000.0,
		CategoryID: category.ID,
	}
	db.Create(&p)

	fmt.Println("\ncreating serial number")
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: p.ID,
	})

	fmt.Println("\n selecting all")
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

	fmt.Println("\n selecting products from a category")
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			fmt.Println(product.Name)
		}
	}
}
