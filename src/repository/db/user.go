package db

import (
	"github.com/jinzhu/gorm"
	connection "example_app/util/helper/mysqlconnection"
	dbEntity "example_app/entity/db"
	"sync"
	"errors"
)

type UserRepository struct {
	DB gorm.DB
}

func UserRepositoryHandler() *UserRepository {
	return &UserRepository{DB: *connection.GetConnection()}
}

type UserRepositoryInterface interface {
	GetUserByID(id int, userData *dbEntity.User, wg *sync.WaitGroup) error
	UpdateUserByID(id int, userData *dbEntity.User) error
	GetUsersList(limit int, offset int) ([]dbEntity.User, error)
	DeleteUserByID(id int, userData *dbEntity.User) error
	CreateUser(userData *dbEntity.User) ([]dbEntity.User, error)
}

func (repository *UserRepository) GetUserByID(id int, userData *dbEntity.User, wg *sync.WaitGroup) error {
	query := repository.DB.Preload("UserStatus")
	query = query.Where("id=?", id)
	query = query.First(userData)
	wg.Done()
	return query.Error
}

func (repository *UserRepository) UpdateUserByID(id int, userData *dbEntity.User) error {
	var count int
	query := repository.DB.Table("users")
	query = query.Where("id=?", id)
	test := query
	test.Count(&count)
	query = query.Updates(userData)
	query.Scan(&userData)
	if count == 0{
		query.Error = errors.New("id not found")
	}
	
	return query.Error
}

func (repository *UserRepository) GetUsersList(limit int, offset int) ([]dbEntity.User, error) {
	users := []dbEntity.User{}
	// user := &dbEntity.User{}
	query := repository.DB.Table("users")
	query = query.Find(&users)
	return users, query.Error
}

func (repository *UserRepository) DeleteUserByID(id int, userData *dbEntity.User) error {
	var count int
	query := repository.DB.Table("users")
	query = query.Where("id=?", id)
	test := query
	test.Count(&count)

	query = query.Find(&userData)
	query = query.Delete(&userData)
	if count == 0{
		query.Error = errors.New("error brook")
	}
	
	return query.Error
}

func (repository *UserRepository) CreateUser(userData *dbEntity.User) ([]dbEntity.User, error) {
	users := []dbEntity.User{}
	// user := &dbEntity.User{}
	query := repository.DB.Table("users")
	query = query.Create(&userData)
	return users, query.Error
}