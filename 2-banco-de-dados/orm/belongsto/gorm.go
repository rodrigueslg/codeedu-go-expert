package belongsto

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

	Product struct {
		ID         int `gorm:primary_key`
		Name       string
		Price      float64
		CategoryID int
		Category   Category
		gorm.Model
	}
)

func RunBelongsToSample() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

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

	fmt.Println("\n selecting all")
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}
}
