package bar

import (
	"github.com/gin-gonic/gin"
)

type BarHandler struct {
	BarService *BarService
}

func NewBarHandler(b *BarService) *BarHandler {
	return &BarHandler{BarService: b}
}

func (b *BarHandler) GetBarById(c *gin.Context) {

}
