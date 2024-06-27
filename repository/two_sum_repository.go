package repository

import (
	"context"
	"dpai/apperror"
)

type TwoSumRepository interface {
	InsertNumber(ctx context.Context, number int, index int) (int, int, error)
	FindByNumber(ctx context.Context, number int) (int, int, error)
	ClearData(ctx context.Context) error
}

type twoSumRepository struct {
	storedNumberIndex *map[int]int
}

func (t *twoSumRepository) ClearData(ctx context.Context) error {
	initMap := make(map[int]int)
	t.storedNumberIndex = &initMap
	return nil
}

func (t *twoSumRepository) FindByNumber(ctx context.Context, number int) (int, int, error) {
	index, exists := (*t.storedNumberIndex)[number]
	if exists == false {
		return 0, 0, apperror.NumberNotInDict
	}
	return number, index, nil
}

func NewTwoSumRepository() TwoSumRepository {
	initMap := make(map[int]int)
	repo := twoSumRepository{storedNumberIndex: &initMap}
	return &repo
}

func (t *twoSumRepository) InsertNumber(ctx context.Context, number int, index int) (int, int, error) {
	(*t.storedNumberIndex)[number] = index
	return number, index, nil
}
