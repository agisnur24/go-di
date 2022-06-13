package main

import (
	"github.com/go-playground/validator/v10"
	"golang-rest-api/app"
	"golang-rest-api/controller"
	"golang-rest-api/helper"
	"golang-rest-api/middleware"
	"golang-rest-api/repository"
	"golang-rest-api/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	// Membuat object categoryRepository, dengan memanggil constructor NewCategoryRepository
	// Dependency : None
	categoryRepository := repository.NewCategoryRepository()

	// Membuat object categoryService, dengan memanggil constructor NewCategoryService
	// Dependency : categoryRepository, db, validate
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	// Membuat object categoryController, dengan memanggil constructor controller category
	// Dependency : categoryService
	categoryController := controller.NewCategoryController(categoryService)

	// Membuar object router, dengan memanggil constructor NewRouter
	// Dependency : categoryController
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
