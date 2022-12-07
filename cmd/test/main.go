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

	router.GET("/foo/{id}")
	http.ListenAndServe(":4000", nil)
}

var FooHandler = foo.NewFooHandler(foo.NewFooService(&repo.FooRepo{}))
var BarHandler = bar.NewBarHandler(bar.NewBarService(&repo.BarRepo{}))
