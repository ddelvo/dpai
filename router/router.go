package router

import (
	"dpai/handler"
	"github.com/gin-gonic/gin"
)

type Options struct {
	TwoSumHandler *handler.TwoSumHandler
}

func NewRouter(opts Options) *gin.Engine {
	r := gin.Default()
	r.POST("/two-sum", opts.TwoSumHandler.FindTwoSum)
	return r
}
