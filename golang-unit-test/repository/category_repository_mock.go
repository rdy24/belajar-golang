package repository

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (mock *CategoryRepositoryMock) FindById(id string) *entity.Category {
	args := mock.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entity.Category)
}
