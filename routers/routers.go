package routers

import (
	ap "gitlab.com/digiverse/gosseractcv/api"
	"gitlab.com/digiverse/gosseractcv/api/v1/ocr"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(ap.CORS)
	r.GET("/", ap.Ping)
	r.GET("/ping", ap.Ping)

	api := r.Group("/api")
	apiV1 := api.Group("/v1")

	apiV1.POST("/parse", ocr.Parse)

	return r
}
