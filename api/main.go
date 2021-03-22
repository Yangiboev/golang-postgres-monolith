package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Yangiboev/golang-postgres-monolith/api/docs"

	v1 "github.com/Yangiboev/golang-postgres-monolith/api/handlers/v1"
	"github.com/Yangiboev/golang-postgres-monolith/config"
	"github.com/Yangiboev/golang-postgres-monolith/pkg/logger"
	"github.com/Yangiboev/golang-postgres-monolith/storage"
)

type Config struct {
	Storage storage.StorageI
	Logger  logger.Logger
	Cfg     config.Config
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Storage: cnf.Storage,
		Logger:  cnf.Logger,
		Cfg:     cnf.Cfg,
	})

	// Category endpoints
	r.GET("/v1/category", handlerV1.GetAllCategories)
	r.GET("/v1/category/:category_id", handlerV1.GetCategory)
	r.POST("/v1/category", handlerV1.CreateCategory)
	r.PUT("/v1/category/:category_id", handlerV1.UpdateCategory)

	// Product endpoints
	r.GET("/v1/price", handlerV1.GetAllPrices)
	r.GET("/v1/price/:price_id", handlerV1.GetPrice)
	r.POST("/v1/price", handlerV1.CreatePrice)
	r.PUT("/v1/price/:price_id", handlerV1.UpdatePrice)

	// Product endpoints
	r.GET("/v1/product", handlerV1.GetAllProducts)
	r.GET("/v1/product/:product_id", handlerV1.GetProduct)
	r.POST("/v1/product", handlerV1.CreateProduct)
	r.PUT("/v1/product/:product_id", handlerV1.UpdateProduct)

	// Retailer endpoints
	r.GET("/v1/retailer", handlerV1.GetAllRetailers)
	r.GET("/v1/retailer/:retailer_id", handlerV1.GetRetailer)
	r.POST("/v1/retailer", handlerV1.CreateRetailer)
	r.PUT("/v1/retailer/:retailer_id", handlerV1.UpdateRetailer)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
