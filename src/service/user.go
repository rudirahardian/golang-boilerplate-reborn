package services

import (
	"sync"
	"github.com/jinzhu/copier"
	httpEntity "example_app/entity/http"
	dbEntity "example_app/entity/db"
	repository "example_app/repository/db"
)

type UserService struct {
	userRepository repository.UserRepositoryInterface
}

func UserServiceHandler() *UserService {
	return &UserService{
		userRepository: repository.UserRepositoryHandler(),
	}
}

type UserServiceInterface interface {
	GetUserByID(id int, waitGroup *sync.WaitGroup) *httpEntity.UserDetailResponse
	GetAllUser(page int,count int) []httpEntity.UserResponse
	UpdateUserByID(id int, payload httpEntity.UserRequest) error
	DeleteUserByID(id int) bool
	CreateUser(payload httpEntity.UserRequestNew) ([]httpEntity.UserResponse, error)
}

func (service *UserService) GetUserByID(id int, waitGroup *sync.WaitGroup) *httpEntity.UserDetailResponse{
	waitGroup.Add(1)
	user := &dbEntity.User{}
	go service.userRepository.GetUserByID(id,user,waitGroup)
	
	waitGroup.Wait()

	result := &httpEntity.UserDetailResponse{}
	return result
}

func (service *UserService) GetAllUser(page int,count int) []httpEntity.UserResponse {
	users, _ := service.userRepository.GetUsersList(page,count)
	result := []httpEntity.UserResponse{}
	copier.Copy(&result, &users)
	return result
}

func (service *UserService) UpdateUserByID(id int, payload httpEntity.UserRequest) error {
	user := &dbEntity.User{}
	user.Name = payload.Name
	user.IDCardNumber = payload.IDCardNumber
	user.Address = payload.Address
	err := service.userRepository.UpdateUserByID(id, user)
	return err
}

func (service *UserService) DeleteUserByID(id int) bool {
	user := &dbEntity.User{}
	err := service.userRepository.DeleteUserByID(id, user)
	if nil != err {
		return false
	}
	return true
}

func (service *UserService) CreateUser(payload httpEntity.UserRequestNew) ([]httpEntity.UserResponse, error) {
	user := &dbEntity.User{}
	user.Name = payload.Name
	user.IDCardNumber = payload.IDCardNumber
	user.Address = payload.Address
	_, err := service.userRepository.CreateUser(user)
	result := []httpEntity.UserResponse{}
	copier.Copy(&result, *user)
	return result, err
}