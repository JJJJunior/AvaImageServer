package routers

import (
	"AvaImageServer/app"
	"AvaImageServer/middlewares"
	"AvaImageServer/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.POST("upload", app.UploadImage)
	return r
}
