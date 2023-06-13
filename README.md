# APLIKASI CRUD SEDERHANA
    - Kita akan membuat aplikasi CRUD sederhana
    - Tujuannya untuk belajar RESTful API, bukan untuk membuat aplikasi
    - Kita akan membuat aplikasi CRUD untuk data category
    - Dimana data category memiliki attribut id (number) dan name (string)
    - Kita akan buat API untuk semua operasi nya, Create Category, Get Category, List Category, Update Category dan Delete Category
    - Semua API akan kita tambahkan Authentication berupa API-Key

# MENAMBAHKAN DEPENDENCY
    - Driver MySQL : https://github.com/go-sql-driver/mysql 
    - HTTP Router : https://github.com/julienschmidt/httprouter
    - Validation : https://github.com/go-playground/validator

# MEMBUAT OPEN API
### - API SPEC LIST CATEGORY
### - API SPEC CREATE CATEGORY
### - API SPEC GET CATEGORY
### - API SPEC PUT CATEGORY
### - API SPEC DELETE CATEGORY


# API SPEC SECURITY


# CREATE DATABASE
```mysql
CREATE TABLE category
(
    id integer primary key auto_increment,
    name varchar(255) not null
) engine = InnoDB;

select * from category;

desc category;
```


# CATEGORY DOMAIN
```go
type Category struct {
	Id   int
	Name string
}
```

# CATEGORY REPOSITORY
```go
type CategoryRepository interface {
	Save(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx sql.Tx, categoryId int) domain.Category
	FindAll(ctx context.Context, tx sql.Tx) []domain.Category
}
```


# CATEGORY REPOSITORY IMPLEMENTATION
```go
//SAVE
func (repository CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	//TODO implement me
	SQL := "INSERT INTO customer(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

//UPDATE
func (repository CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	//TODO implement me
	SQL := "UPDATE category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category

}

//DELETE
func (repository CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	//TODO implement me
	SQL := "DELETE FROM category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

//FIND BY ID
func (repository CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	//TODO implement me
	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

//FIND ALL
func (repository CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	//TODO implement me
	SQL := "SELECT * FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var categories []domain.Category
	if rows.Next() {
		category := domain.Category{}
		err := rows.Scan(category.Id, category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
```


# CATEGORY SERVICE
```go
type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
```

# CATEGORY SERVICE IMPLEMENTATION
```go
type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository

	DB *sql.DB
}

func (service CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	//TODO implement me
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
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

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
```


# CATEGORY VALIDATION
```go
===
type CategoryCreateRequest struct {
	Name string `validate:"required"`
}
===

===
type CategoryUpdateRequest struct {
Id   int    `validate:"required"`
Name string `validate:"required,min=1"`
}
===

===
type CategoryServiceImpl struct {
CategoryRepository repository.CategoryRepository
DB                 *sql.DB
Validate           *validator.Validate
}
===

===
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
```


# CATEGORY CONTROLLER
```go
type CategoryController interface {
	Create(writer http.ResponseWriter, request http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request http.Request, params httprouter.Params)
}
```

# CATEGORY CONTROLLER IMPLEMENTATION
```go
type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (controller CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)

}

func (controller CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	decoder := json.NewDecoder(request.Body)
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	helper.PanicIfError(err)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   "",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   categoryResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
```


# CATEGORY SERVICE
```go
type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
```


# CATEGORY SERVICE IMPLEMENTATION
```go
type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
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
	helper.PanicIfError(err)

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
```


# HTTP ROUTER
```go
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
}

```


# HTTP SERVER