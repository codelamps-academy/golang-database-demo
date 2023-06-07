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
