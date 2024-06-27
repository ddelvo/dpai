package handler

import (
	"dpai/dto"
	"dpai/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TwoSumHandler struct {
	useCase usecase.TwoSumUseCase
}

func NewTwoSumHandler(serviceParam usecase.TwoSumUseCase) *TwoSumHandler {
	handler := TwoSumHandler{useCase: serviceParam}
	return &handler
}

func (tsh *TwoSumHandler) FindTwoSum(ctx *gin.Context) {
	var req dto.TwoSumRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := tsh.useCase.FindTwoSum(ctx, req.Numbers, req.Target)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Message: "Success", Data: res})
}
