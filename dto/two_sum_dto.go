package dto

type TwoSumRequest struct {
	Numbers []int `json:"numbers" binding:"required"`
	Target  int   `json:"target" binding:"required"`
}
