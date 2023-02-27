package repository

import "go-init-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
