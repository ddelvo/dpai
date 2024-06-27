package dto

type ThreeSumRequest struct {
	Numbers []int `json:"numbers" binding:"required"`
}
