package api

import (
	_ "github.com/Yangiboev/golang-neo4j/api/docs"
	v1 "github.com/Yangiboev/golang-neo4j/api/handler/v1"
	"github.com/Yangiboev/golang-neo4j/config"
	"github.com/Yangiboev/golang-neo4j/pkg/logger"
	"github.com/Yangiboev/golang-neo4j/storage"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Config  config.Config
	Log     logger.Logger
	Storage storage.StorageI
}

func New(ro RouterOptions) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Storage: ro.Storage,
		Logger:  ro.Log,
		Cfg:     ro.Config,
	})
	// Category endpoints
	router.GET("/v1/responsible", handlerV1.GetAllResponsibles)
	router.GET("/v1/responsible/:responsible_id", handlerV1.GetResponsible)
	router.POST("/v1/responsible", handlerV1.CreateResponsible)
	router.PUT("/v1/responsible/:responsible_id", handlerV1.UpdateResponsible)
	router.DELETE("/v1/responsible/:responsible_id", handlerV1.DeleteResponsible)

	// Action endpoints
	router.GET("/v1/action", handlerV1.GetAllActions)
	router.GET("/v1/action/:action_id", handlerV1.GetAction)
	router.POST("/v1/action", handlerV1.CreateAction)
	router.PUT("/v1/action/:action_id", handlerV1.UpdateAction)
	router.DELETE("/v1/action/:action_id", handlerV1.DeleteAction)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router

}
