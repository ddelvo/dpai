package router

import (
	"dpai/handler"
	"github.com/gin-gonic/gin"
)

type Options struct {
	TwoSumHandler   *handler.TwoSumHandler
	ThreeSumHandler *handler.ThreeSumHandler
}

func NewRouter(opts Options) *gin.Engine {
	r := gin.Default()
	r.POST("/two-sum", opts.TwoSumHandler.FindTwoSum)
	r.POST("/three-sum", opts.ThreeSumHandler.FindThreeSum)
	return r
}
