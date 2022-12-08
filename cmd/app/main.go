package main

import (
	"net/http"
	"test/openapi-generator/internal/bar"
	"test/openapi-generator/internal/foo"
	"test/openapi-generator/internal/repo"

	docs "test/openapi-generator/internal/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Test API"
	docs.SwaggerInfo.Description = "Testing code generation from swagger docs"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	router := gin.New()
	router.Use(gin.Recovery())

	doc := router.Group("/doc")

	doc.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/docs/index.html#")
	})

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API Docs"})
	})

	f := router.Group("/foo")
	f.GET("", FooHandler.ReadMany)
	f.GET("/:id", FooHandler.ReadOne)
	f.POST("", FooHandler.Create)
	f.PATCH("/:id", FooHandler.Update)
	f.DELETE("/:id", FooHandler.Delete)

	b := router.Group("/bar")
	b.GET("", BarHandler.ReadMany)
	b.GET("/:id", BarHandler.ReadOne)
	b.POST("", BarHandler.Create)
	b.PATCH("/:id", BarHandler.Update)
	b.DELETE("/:id", BarHandler.Delete)

	http.ListenAndServe(":4000", router)
}

var FooHandler = foo.NewFooHandler(foo.NewFooService(&repo.FooRepo{}))
var BarHandler = bar.NewBarHandler(bar.NewBarService(&repo.BarRepo{}))
