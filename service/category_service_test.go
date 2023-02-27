package service

import (
	"go-init-test/entity"
	"go-init-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get_Not_Found(t *testing.T) {
	// given
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	// when
	category, err := categoryService.Get("1")

	// then
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategiryService_Get_Success(t *testing.T) {
	// given
	category := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	// when
	categoryResult, err := categoryService.Get("2")

	// then
	assert.NotNil(t, categoryResult)
	assert.Equal(t, category.Id, categoryResult.Id)
	assert.Equal(t, category.Name, categoryResult.Name)
	assert.Nil(t, err)
}
