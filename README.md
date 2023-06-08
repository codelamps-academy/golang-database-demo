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

```