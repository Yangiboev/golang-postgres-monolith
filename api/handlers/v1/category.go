package v1

import (
	"net/http"

	"github.com/Yangiboev/golang-postgres-monolith/storage/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router /v1/category/{category_id} [get]
// @Summary Get Category
// @Description API for getting a category
// @Tags category
// @Accept  json
// @Produce  json
// @Param category_id path string true "category_id"
// @Success 200 {object} models.GetCategory
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetCategory(c *gin.Context) {

	categoryID := c.Param("category_id")

	_, err := uuid.Parse(categoryID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "category_id format is invalid format!",
		})
		return
	}
	category, err := h.storage.Category().Get(categoryID)

	if err != nil {
		HandleInternalErrWithMessage(c, h.log, err, "Error while getting category by id")
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    category,
		})

}

// @Router /v1/category [post]
// @Summary Create category
// @Description API for creating category
// @Tags category
// @Accept json
// @Produce json
// @Param Category body models.CreateCategoryRequest  true "category"
// @Success 200 {object} models.CreateSuccessResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) CreateCategory(c *gin.Context) {

	var (
		category repo.Category
	)

	err := c.ShouldBindJSON(&category)

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
	category.Id = id.String()
	resp, err := h.storage.Category().Create(
		&category)

	if !handleError(h.log, c, err, "error while creating category") {
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
			"data":    resp,
		})
}

// @Router /v1/category [get]
// @Summary Get All Categories
// @Description API for getting all categories
// @Tags category
// @Accept  json
// @Produce  json
// @Param name path string false "name"
// @Success 200 {object} models.GetAllCategoriesResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetAllCategories(c *gin.Context) {
	name := c.Query("name")
	categories, err := h.storage.Category().GetAll(name)

	if !handleError(h.log, c, err, "error while getting all categories") {
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    categories,
		})

}

// @Router /v1/category/{category_id} [put]
// @Summary Update category
// @Description API for creating category
// @Tags category
// @Accept json
// @Produce json
// @Param category_id path string true "category_id"
// @Param Category body models.CreateCategoryRequest  true "category"
// @Success 200 {object} models.CreateSuccessResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) UpdateCategory(c *gin.Context) {
	var (
		category   repo.Category
		categoryID string
	)
	categoryID = c.Param("category_id")
	_, err := uuid.Parse(categoryID)
	if !handleError(h.log, c, err, "category_id is invalid format") {
		return
	}

	err = c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	category.Id = categoryID
	resp, err := h.storage.Category().Update(
		&category)

	if !handleError(h.log, c, err, "error while updating category") {
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
			"data":    resp,
		})
}
