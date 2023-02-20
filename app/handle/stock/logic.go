package stock

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/base"
)

type StockService struct{}

func (s *StockService) GetList(ctx *gin.Context) *base.Response {
	return nil
}
