package main

import (
	"github.com/gin-gonic/gin"
	productcontroller "github.com/luthfyhakim/go-restapi-gin/controllers/productController"
	"github.com/luthfyhakim/go-restapi-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabases()

	r.GET("api/products", productcontroller.Index)
	r.GET("api/products/:id", productcontroller.Show)
	r.POST("api/products", productcontroller.Create)
	r.PUT("api/products/:id", productcontroller.Update)
	r.DELETE("api/products", productcontroller.Delete)

	r.Run(":8888") // listen and serve on 0.0.0.0:8888
}
