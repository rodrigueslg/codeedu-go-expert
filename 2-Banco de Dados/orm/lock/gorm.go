package lock

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func RunLockSample() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	fmt.Println("\ncreating single category")
	cEletronicos := Category{Name: "Eletrônicos"}
	db.Create(&cEletronicos)

	fmt.Println("\ncreating single category")
	cCozinha := Category{Name: "Cozinha"}
	db.Create(&cCozinha)

	fmt.Println("\ncreating single product")
	var pNotebook Product = Product{
		Name:       "Notebook",
		Price:      1000.0,
		Categories: []Category{cEletronicos, cCozinha},
	}
	db.Create(&pNotebook)

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

	// locking the row for concurrent updates
	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error // SELECT * FROM categories WHERE id = 1 FOR UPDATE;
	if err != nil {
		panic(err)
	}
	c.Name = "Eletrônicos 2"
	tx.Debug().Save(&c)
	tx.Commit()
}
