package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"golang-database-demo/controller"
	"golang-database-demo/exception"
	"golang-database-demo/helper"
	"golang-database-demo/repository"
	"golang-database-demo/service"
	"net/http"
)

func main() {

	db := helper.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	// ROUTER UNTUK FIND ALL
	router.GET("/api/categories", categoryController.FindAll)

	// ROUTER UNTUK FIND BY ID
	router.GET("/api/categories/:categoryId", categoryController.FindById)

	// ROUTER UNTUK CREATE
	router.POST("/api/categories", categoryController.Create)

	// ROUTER UNTUK UPDATE
	router.PUT("/api/categories/:categoryId", categoryController.Update)

	// ROUTER UNTUK DELETE
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
