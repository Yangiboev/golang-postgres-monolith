package v1

import (
	"net/http"

	"github.com/Yangiboev/golang-postgres-monolith/storage/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router /v1/price/{price_id} [get]
// @Summary Get Price
// @Description API for getting a price
// @Tags price
// @Accept  json
// @Produce  json
// @Param price_id path string true "price_id"
// @Success 200 {object} models.GetPrice
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetPrice(c *gin.Context) {
	priceID := c.Param("price_id")
	_, err := uuid.Parse(priceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "price_id format is invalid format!",
		})
		return
	}
	price, err := h.storage.Price().Get(priceID)

	if !handleError(h.log, c, err, "error while getting price by id") {
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    price,
		})

}

//@Router /v1/price [post]
//@Summary Create price
//@Description API for creating price
//@Tags price
//@Accept json
//@Produce json
//@Param Price body models.CreatePriceRequest  true "price"
// @Success 200 {object} models.CreateSuccessResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) CreatePrice(c *gin.Context) {
	var (
		price repo.Price
	)
	err := c.ShouldBindJSON(&price)

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
	price.Id = id.String()
	resp, err := h.storage.Price().Create(
		&price)

	if !handleError(h.log, c, err, "error while creating price") {
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
			"data":    resp,
		})
}

// @Router /v1/price [get]
// @Summary Get All Prices
// @Description API for getting all prices
// @Tags price
// @Accept  json
// @Produce  json
// @Param price path string false "price"
// @Success 200 {object} models.GetAllPricesResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetAllPrices(c *gin.Context) {
	price := c.Query("price")
	prices, err := h.storage.Price().GetAll(price)

	if !handleError(h.log, c, err, "error while getting all prices") {
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    prices,
		})

}

//@Router /v1/price/{price_id} [put]
//@Summary Update price
//@Description API for creating price
//@Tags price
//@Accept json
//@Produce json
// @Param price_id path string true "price_id"
//@Param Price body models.CreatePriceRequest  true "price"
// @Success 200 {object} models.CreateSuccessResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) UpdatePrice(c *gin.Context) {
	var (
		price   repo.Price
		priceID string
	)
	priceID = c.Param("price_id")
	_, err := uuid.Parse(priceID)
	if !handleError(h.log, c, err, "price_id is invalid format") {
		return
	}

	err = c.ShouldBindJSON(&price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	price.Id = priceID
	resp, err := h.storage.Price().Update(
		&price)

	if !handleError(h.log, c, err, "error while updating price") {
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
			"data":    resp,
		})
}
