package bar

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BarHandler struct {
	BarService *BarService
}

func NewBarHandler(b *BarService) *BarHandler {
	return &BarHandler{BarService: b}
}

// ReadOne godoc
// @Summary Read a Bar
// @Description Read a Bar
// @Tags Bar
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path int true "Bar ID"
// @Success 200 {object} Bar
// @Failure 400 {object} any
// @Failure 404 {object} any
// @Failure 500 {object} any
// @Router /bar/{id} [get]
func (b *BarHandler) ReadOne(c *gin.Context) {
	// get ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(200, b.BarService.BarRepo.Find(id))
}

// ReadMany godoc
// @Summary Read Many Bar
// @Description Read Many Bar
// @Tags Bar
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param order[id] query string false "Order by id" enums(asc, desc)
// @Param pagination query boolean false "Pagination"
// @Param page query int false "Page"
// @Param rpp query int false "Results per page"
// @Success 200 {object} BarResponse
// @Failure 400 {object} any
// @Failure 500 {object} any
// @Router /bar [get]
func (b *BarHandler) ReadMany(c *gin.Context) {
	c.JSON(200, b.BarService.BarRepo.FindAll())
}

// Create godoc
// @Summary Create Bar
// @Description POST a bar, make a bar
// @Tags Bar
// @Accept application/json
// @Produce json
// @Param Bar body Bar true "Bar"
// @Success 201 {object} any
// @Failure 400 {object} any
// @Router /bar [post]
func (b *BarHandler) Create(c *gin.Context) {
	c.JSON(201, b.BarService.BarRepo.Create())
}

// Update godoc
// @Summary Update Bar
// @Description Patch a Bar
// @Tags Bar
// @Accept application/json
// @Produce json
// @Param id path int true "Bar ID"
// @Param Bar body Bar true "Bar"
// @Success 200 {object} any
// @Failure 400 {object} any
// @Failure 404 {object} any
// @Failure 500 {object} any
// @Router /bar/{id} [patch]
func (b *BarHandler) Update(c *gin.Context) {
	c.JSON(200, b.BarService.BarRepo.Update())

}

// Delete godoc
// @Summary Delete Bar
// @Description Delete Bar
// @Tags Bar
// @Accept application/json
// @Produce json
// @Param id path int true "Bar ID"
// @Success 204 {object} any
// @Failure 400 {object} any
// @Failure 404 {object} any
// @Failure 500 {object} any
// @Router /bar/{id} [delete]
func (b *BarHandler) Delete(c *gin.Context) {
	c.JSON(204, b.BarService.BarRepo.Delete())
}
