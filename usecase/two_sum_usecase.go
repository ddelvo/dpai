package usecase

import (
	"context"
	"dpai/apperror"
	"dpai/repository"
	"errors"
)

type TwoSumUseCase interface {
	FindTwoSum(ctx context.Context, numbers []int, target int) ([]int, error)
}

type twoSumUseCase struct {
	repository repository.TwoSumRepository
}

func (t twoSumUseCase) FindTwoSum(ctx context.Context, numbers []int, target int) ([]int, error) {
	_ = t.repository.ClearData(ctx)
	final := make([]int, 0)
	for i, number := range numbers {
		needed := target - number
		_, index, err := t.repository.FindByNumber(ctx, needed)
		if errors.Is(err, apperror.NumberNotInDict) {
			_, _, _ = t.repository.InsertNumber(ctx, number, i)
			continue
		}
		final = append(final, i, index)
		return final, nil
	}
	return final, nil
}

func NewTwoSumUseCase(b repository.TwoSumRepository) TwoSumUseCase {
	return &twoSumUseCase{
		repository: b,
	}
}
