package v1

import (
	"net/http"

	"github.com/Yangiboev/golang-postgres-monolith/storage/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router /v1/product/{product_id} [get]
// @Summary Get Product
// @Description API for getting a product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product_id path string true "product_id"
// @Success 200 {object} models.GetProduct
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetProduct(c *gin.Context) {
	productID := c.Param("product_id")
	_, err := uuid.Parse(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "product_id format is invalid format!",
		})
		return
	}
	product, err := h.storage.Product().Get(productID)

	if !handleError(h.log, c, err, "error while getting product by id") {
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    product,
		})

}

//@Router /v1/product [post]
//@Summary Create product
//@Description API for creating product
//@Tags product
//@Accept json
//@Produce json
//@Param Product body models.CreateProductRequest  true "product"
// @Success 200 {object} models.CreateSuccessResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) CreateProduct(c *gin.Context) {
	var (
		product repo.Product
	)
	err := c.ShouldBindJSON(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := uuid.NewRandom()
	if !handleError(h.log, c, err, "error while generating uuid") {
		return
	}
	product.Id = id.String()
	resp, err := h.storage.Product().Create(
		&product)

	if !handleError(h.log, c, err, "error while creating product") {
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
			"data":    resp,
		})
}

// @Router /v1/product [get]
// @Summary Get All Products
// @Description API for getting all Products
// @Tags product
// @Accept  json
// @Produce  json
// @Param name path string false "name"
// @Success 200 {object} models.GetAllProductsResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetAllProducts(c *gin.Context) {
	name := c.Query("name")
	products, err := h.storage.Product().GetAll(name)

	if !handleError(h.log, c, err, "error while getting all products") {
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    products,
		})

}

//@Router /v1/product/{product_id} [put]
//@Summary Update product
//@Description API for creating product
//@Tags product
//@Accept json
//@Produce json
// @Param product_id path string true "product_id"
//@Param Product body models.CreateProductRequest  true "product"
// @Success 200 {object} models.CreateSuccessResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) UpdateProduct(c *gin.Context) {
	var (
		product   repo.Product
		productID string
	)
	productID = c.Param("product_id")
	_, err := uuid.Parse(productID)
	if !handleError(h.log, c, err, "product_id is invalid format") {
		return
	}

	err = c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	product.Id = productID
	resp, err := h.storage.Product().Update(
		&product)

	if !handleError(h.log, c, err, "error while updating product") {
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
			"data":    resp,
		})
}
