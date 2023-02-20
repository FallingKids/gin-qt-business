package news

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/base"
)

type NewsController struct{}

func (n *NewsController) GetList(ctx *gin.Context) *base.Response {
	return nil
}
