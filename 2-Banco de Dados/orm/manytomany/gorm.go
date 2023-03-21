package manytomany

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Category struct {
		ID       int `gorm:primary_key`
		Name     string
		Products []Product `gorm:"many2many:product_categories;"`
	}

	Product struct {
		ID         int `gorm:primary_key`
		Name       string
		Price      float64
		Categories []Category `gorm:"many2many:product_categories;"`
		gorm.Model
	}
)

func RunManyToManySample() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// fmt.Println("\ncreating single category")
	// cEletronicos := Category{Name: "Eletrônicos"}
	// db.Create(&cEletronicos)

	// fmt.Println("\ncreating single category")
	// cCozinha := Category{Name: "Cozinha"}
	// db.Create(&cCozinha)

	// fmt.Println("\ncreating single product")
	// var pNotebook Product = Product{
	// 	Name:       "Notebook",
	// 	Price:      1000.0,
	// 	Categories: []Category{cEletronicos, cCozinha},
	// }
	// db.Create(&pNotebook)

	fmt.Println("\n selecting products from a category")
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("-", product.Name)
		}
	}
}
