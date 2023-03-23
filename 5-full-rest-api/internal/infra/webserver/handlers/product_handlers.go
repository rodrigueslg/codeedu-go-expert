package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rodrigueslg/codedu-goexpert/rest-api/internal/dto"
	"github.com/rodrigueslg/codedu-goexpert/rest-api/internal/entity"
	"github.com/rodrigueslg/codedu-goexpert/rest-api/internal/infra/database"

	entitypkg "github.com/rodrigueslg/codedu-goexpert/rest-api/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param request body dto.CreateProductInput true "Product request"
// @Success 201 {object} entity.Product
// @Failure 400 {object} dto.Error
// @Failure 500
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		err := dto.Error{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	p, err := entity.NewProduct(input.Name, input.Price)
	if err != nil {
		err := dto.Error{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetProduct godoc
// @Summary Get a product by id
// @Description Get a product by id
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" Format(uuid)
// @Success 200 {object} entity.Product
// @Failure 400 {object} dto.Error
// @Failure 404
// @Failure 500
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		error := dto.Error{Message: "invalid id"}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(error)
		return
	}

	p, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(p)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" Format(uuid)
// @Param request body dto.CreateProductInput true "Product request"
// @Success 200 {object} entity.Product
// @Failure 400 {object} dto.Error
// @Failure 404
// @Failure 500
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		error := dto.Error{Message: "invalid id"}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(error)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		error := dto.Error{Message: "invalid request body"}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(error)
		return
	}
	product.ID, err = entitypkg.ParseID(id)
	if err != nil {
		error := dto.Error{Message: "invalid id format"}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(error)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" Format(uuid)
// @Success 200
// @Failure 400 {object} dto.Error
// @Failure 404
// @Failure 500
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		error := dto.Error{Message: "invalid id"}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(error)
		return
	}

	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Param page query string false "Page number"
// @Param limit query string false "Limit of products per page"
// @Param sort query string false "Sort asc or desc"
// @Success 200 {array} entity.Product
// @Failure 500
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")
	sortParam := r.URL.Query().Get("sort")

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 0
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 0
	}

	products, err := h.ProductDB.FindAll(page, limit, sortParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(products)
}
