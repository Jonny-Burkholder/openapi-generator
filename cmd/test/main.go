package main

import (
	"net/http"
	"test/openapi-generator/internal/bar"
	"test/openapi-generator/internal/foo"
	"test/openapi-generator/internal/repo"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	f := router.Group("/foo")
	f.GET("", FooHandler.ReadMany)
	f.GET("/{id}", FooHandler.ReadOne)
	f.POST("", FooHandler.Create)
	f.PATCH("/{id}", FooHandler.Update)
	f.DELETE("/{id}", FooHandler.Delete)

	b := router.Group("/bar")
	b.GET("", BarHandler.ReadMany)
	b.GET("/{id}", BarHandler.ReadOne)
	b.POST("", BarHandler.Create)
	b.PATCH("/{id}", BarHandler.Update)
	b.DELETE("/{id}", BarHandler.Delete)

	http.ListenAndServe(":4000", router)
}

var FooHandler = foo.NewFooHandler(foo.NewFooService(&repo.FooRepo{}))
var BarHandler = bar.NewBarHandler(bar.NewBarService(&repo.BarRepo{}))
