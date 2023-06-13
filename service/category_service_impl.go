package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"golang-database-demo/exception"
	"golang-database-demo/helper"
	"golang-database-demo/model/domain"
	"golang-database-demo/model/web"
	"golang-database-demo/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	//TODO implement me
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (service CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	//TODO implement me
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	update := service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(update)
}

func (service CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	var categoryReponse []web.CategoryResponse
	for _, category := range categories {
		categoryReponse = append(categoryReponse, helper.ToCategoryResponse(category))
	}

	return categoryReponse
}
