package handler

import (
	"dpai/dto"
	"dpai/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ThreeSumHandler struct {
	useCase usecase.ThreeSumUseCase
}

func NewThreeSumHandler(serviceParam usecase.ThreeSumUseCase) *ThreeSumHandler {
	handler := ThreeSumHandler{useCase: serviceParam}
	return &handler
}

func (tsh *ThreeSumHandler) FindThreeSum(ctx *gin.Context) {
	var req dto.ThreeSumRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := tsh.useCase.FindThreeSum(ctx, req.Numbers)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Message: "Success", Data: res})
}
