package helper

import (
	"golang-database-demo/model/domain"
	"golang-database-demo/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
