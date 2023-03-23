package repository

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/rodrigueslg/codedu-goexpert/rest-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateProductDBConnection(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	_ = db.AutoMigrate(&entity.Product{})
	return db
}

func TestCreateNewProduct(t *testing.T) {
	db := CreateProductDBConnection(t)

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)

	productRepo := NewProductRepository(db)
	err = productRepo.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID.String())
}

func TestFindAllProducts(t *testing.T) {
	db := CreateProductDBConnection(t)

	for i := 0; i < 10; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productRepo := NewProductRepository(db)
	products, err := productRepo.FindAll(1, 5, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "Product 0", products[0].Name)
	assert.Equal(t, "Product 4", products[4].Name)

	products, err = productRepo.FindAll(2, 5, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "Product 5", products[0].Name)
	assert.Equal(t, "Product 9", products[4].Name)
}

func TestFindProductByID(t *testing.T) {
	db := CreateProductDBConnection(t)

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)

	productRepo := NewProductRepository(db)
	product, err = productRepo.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", product.Name)
}

func TestUpdateProduct(t *testing.T) {
	db := CreateProductDBConnection(t)

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)

	productRepo := NewProductRepository(db)
	product.Name = "Product 2"
	err = productRepo.Update(product)
	assert.NoError(t, err)

	product, err = productRepo.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	db := CreateProductDBConnection(t)

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)

	productRepo := NewProductRepository(db)
	err = productRepo.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productRepo.FindByID(product.ID.String())
	assert.Error(t, err)
}
