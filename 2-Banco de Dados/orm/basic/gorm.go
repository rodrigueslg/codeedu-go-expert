package basic

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RunBasicSample() {
	type Product struct {
		ID    int `gorm:primary_key`
		Name  string
		Price float64
	}

	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	fmt.Println("\ncreating single")
	db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.0,
	})

	fmt.Println("\n creating batch")
	products := []Product{
		{Name: "Notebook", Price: 1000.0},
		{Name: "Mouse", Price: 50.0},
		{Name: "Keyboard", Price: 100.0},
	}
	db.Create(&products)

	fmt.Println("\n selecting by id")
	var product Product
	db.First(&product, 2)
	fmt.Println(product)

	fmt.Println("\n selecting by name")
	var product2 Product
	db.First(&product2, "name = ?", "Mouse")
	fmt.Println(product2)

	fmt.Println("\n selecting all")
	var productsAll []Product
	db.Find(&productsAll)
	for _, product := range productsAll {
		fmt.Println(product)
	}

	fmt.Println("\n selecting with limit")
	var productsLimit []Product
	db.Limit(2).Find(&productsLimit)
	for _, product := range productsLimit {
		fmt.Println(product)
	}

	fmt.Println("\n selecting with limit and offset")
	var productsLimitOffset []Product
	db.Limit(2).Offset(2).Find(&productsLimitOffset)
	for _, product := range productsLimitOffset {
		fmt.Println(product)
	}

	fmt.Println("\n selecting where")
	var productsWhere []Product
	db.Where("name LIKE ?", "%ouse").Find(&productsWhere)
	for _, product := range productsWhere {
		fmt.Println(product)
	}

	fmt.Println("\n updating")
	var p Product
	db.First(&p, 1)
	p.Name = "New mouse"
	db.Save(&p)
	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)

	fmt.Println("\n deleting")
	db.Delete(&p2)
}

// using gorm.Model to add created_at, updated_at and deleted_at and soft delete
func RunModelSample() {
	type Product struct {
		ID    int `gorm:primary_key`
		Name  string
		Price float64
		gorm.Model
	}

	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	fmt.Println("\ncreating single")
	db.Create(&Product{
		Name:  "Notebook 1",
		Price: 1000.0,
	})

	fmt.Println("\ncreating single")
	db.Create(&Product{
		Name:  "Notebook 2",
		Price: 2000.0,
	})

	fmt.Println("\ncreating single")
	db.Create(&Product{
		Name:  "Notebook 3",
		Price: 3000.0,
	})

	fmt.Println("\n updating")
	var p Product
	db.First(&p, 2)
	p.Name = "New Notebook"
	db.Save(&p)
	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)

	fmt.Println("\n deleting")
	var pDelete Product
	db.First(&pDelete, 2)
	db.Delete(&pDelete)
}

func RunHasManySample() {
	type (
		Category struct {
			ID       int `gorm:primary_key`
			Name     string
			Products []Product
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
