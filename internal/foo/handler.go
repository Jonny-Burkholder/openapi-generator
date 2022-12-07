package foo

import "github.com/gin-gonic/gin"

type FooHandler struct {
	FooService *FooService
}

func NewFooHandler(f *FooService) *FooHandler {
	return &FooHandler{FooService: f}
}

func (f *FooHandler) GetFooById(c gin.Context) {
	// get ID
	var id int
	c.JSON(200, f.FooService.Find(id))
}

func (f *FooHandler) GetFoo(c gin.Context) {
	c.JSON(200, f.FooService.FindAll())
}
