package repository

import (
	"strings"

	"github.com/rodrigueslg/codedu-goexpert/rest-api/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (p *ProductRepository) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductRepository) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *ProductRepository) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product, "id = ?", id).Error
}

func (p *ProductRepository) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *ProductRepository) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var err error
	var products []entity.Product

	sort = strings.ToLower(sort)
	if sort == "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page > 0 && limit > 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}
