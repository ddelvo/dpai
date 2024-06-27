package usecase

import (
	"context"
	"dpai/repository"
	"sort"
)

type ThreeSumUseCase interface {
	FindThreeSum(ctx context.Context, numbers []int) ([][]int, error)
	Populate(ctx context.Context, numbers []int) error
}

type threeSumUseCase struct {
	repository repository.TwoSumRepository
}

func (t threeSumUseCase) FindThreeSum(ctx context.Context, numbers []int) ([][]int, error) {
	_ = t.repository.ClearData(ctx)
	err := t.Populate(ctx, numbers)
	if err != nil {
		return nil, err
	}

	final := make(map[[3]int]struct{})
	var result [][]int

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			needed := -numbers[i] - numbers[j]

			_, index, err := t.repository.FindByNumber(ctx, needed)
			if err == nil && index != i && index != j {

				triplet := []int{numbers[i], numbers[j], needed}
				sort.Ints(triplet[:])
				final[[3]int{triplet[0], triplet[1], triplet[2]}] = struct{}{}
			}

		}
	}

	for item := range final {
		result = append(result, []int{item[0], item[1], item[2]})
	}
	_ = t.repository.ClearData(ctx)

	return result, nil
}

func (t threeSumUseCase) Populate(ctx context.Context, numbers []int) error {
	for index, num := range numbers {
		_, _, _ = t.repository.InsertNumber(ctx, num, index)
	}
	return nil
}

func NewThreeSumUseCase(b repository.TwoSumRepository) ThreeSumUseCase {
	return &threeSumUseCase{
		repository: b,
	}
}
