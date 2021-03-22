package v1

import (
	"net/http"

	"github.com/Yangiboev/golang-postgres-monolith/storage/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router /v1/retailer/{retailer_id} [get]
// @Summary Get Retailer
// @Description API for getting a retailer
// @Tags retailer
// @Accept  json
// @Produce  json
// @Param retailer_id path string true "retailer_id"
// @Success 200 {object} models.GetRetailer
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetRetailer(c *gin.Context) {
	retailerID := c.Param("retailer_id")
	_, err := uuid.Parse(retailerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "retailer_id format is invalid format!",
		})
		return
	}
	retailer, err := h.storage.Retailer().Get(retailerID)

	if !handleError(h.log, c, err, "error while getting retailer by id") {
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    retailer,
		})

}

//@Router /v1/retailer [post]
//@Summary Create retailer
//@Description API for creating retailer
//@Tags retailer
//@Accept json
//@Produce json
//@Param Retailer body models.CreateRetailerRequest  true "retailer"
// @Success 200 {object} models.CreateSuccessResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) CreateRetailer(c *gin.Context) {
	var (
		retailer repo.Retailer
	)
	err := c.ShouldBindJSON(&retailer)

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
	retailer.Id = id.String()
	resp, err := h.storage.Retailer().Create(
		&retailer)

	if !handleError(h.log, c, err, "error while creating retailer") {
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
			"data":    resp,
		})
}

// @Router /v1/retailer [get]
// @Summary Get All Retailers
// @Description API for getting all Retailers
// @Tags retailer
// @Accept  json
// @Produce  json
// @Param name path string false "name"
// @Success 200 {object} models.GetAllRetailersResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetAllRetailers(c *gin.Context) {
	name := c.Query("name")
	retailers, err := h.storage.Retailer().GetAll(name)

	if !handleError(h.log, c, err, "error while getting all retailers") {
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    retailers,
		})

}

//@Router /v1/retailer/{retailer_id} [put]
//@Summary Update retailer
//@Description API for creating retailer
//@Tags retailer
//@Accept json
//@Produce json
// @Param retailer_id path string true "retailer_id"
//@Param Retailer body models.CreateRetailerRequest  true "retailer"
// @Success 200 {object} models.CreateSuccessResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) UpdateRetailer(c *gin.Context) {
	var (
		retailer   repo.Retailer
		retailerID string
	)
	retailerID = c.Param("retailer_id")
	_, err := uuid.Parse(retailerID)
	if !handleError(h.log, c, err, "retailer_id is invalid format") {
		return
	}

	err = c.ShouldBindJSON(&retailer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	retailer.Id = retailerID
	resp, err := h.storage.Retailer().Update(
		&retailer)

	if !handleError(h.log, c, err, "error while updating retailer") {
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
			"data":    resp,
		})
}
