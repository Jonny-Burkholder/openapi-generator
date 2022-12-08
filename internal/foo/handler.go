package foo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FooHandler struct {
	FooService *FooService
}

func NewFooHandler(f *FooService) *FooHandler {
	return &FooHandler{FooService: f}
}

// ReadOne godoc
// @Summary Read a Foo
// @Description Read a Foo
// @Tags Foo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path int true "Foo ID"
// @Success 200 {object} Foo
// @Failure 400 {object} any
// @Failure 404 {object} any
// @Failure 500 {object} any
// @Router /foo/{id} [get]
func (f *FooHandler) ReadOne(c *gin.Context) {
	// get ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(200, f.FooService.Find(id))
}

// ReadMany godoc
// @Summary Read Many Foo
// @Description Read May Foo
// @Tags Foo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param order[id] query string false "Order by id" enums(asc, desc)
// @Param pagination query boolean false "Pagination"
// @Param page query int false "Page"
// @Param rpp query int false "Results per page"
// @Success 200 {object} FooResponse
// @Failure 400 {object} any
// @Failure 500 {object} any
// @Router /foo [get]
func (f *FooHandler) ReadMany(c *gin.Context) {
	c.JSON(200, f.FooService.FindAll())
}

// Create godoc
// @Summary Create Foo
// @Description POST a foo, make a foo
// @Tags Foo
// @Accept application/json
// @Produce json
// @Param Foo body Foo true "Foo"
// @Success 201 {object} any
// @Failure 400 {object} any
// @Router /foo [post]
func (f *FooHandler) Create(c *gin.Context) {
	c.JSON(201, f.FooService.FooRepo.Create())
}

// Update godoc
// @Summary Update Foo
// @Description Patch a Foo
// @Tags Foo
// @Accept application/json
// @Produce json
// @Param id path int true "Foo ID"
// @Param Foo body Foo true "Foo"
// @Success 200 {object} any
// @Failure 400 {object} any
// @Failure 404 {object} any
// @Failure 500 {object} any
// @Router /foo/{id} [patch]
func (f *FooHandler) Update(c *gin.Context) {
	c.JSON(200, f.FooService.FooRepo.Update())

}

// Delete godoc
// @Summary Delete Foo
// @Description Delete Foo
// @Tags Foo
// @Accept application/json
// @Produce json
// @Param id path int true "Foo ID"
// @Success 204 {object} any
// @Failure 400 {object} any
// @Failure 404 {object} any
// @Failure 500 {object} any
// @Router /foo/{id} [delete]
func (f *FooHandler) Delete(c *gin.Context) {
	c.JSON(204, f.FooService.FooRepo.Delete())
}
