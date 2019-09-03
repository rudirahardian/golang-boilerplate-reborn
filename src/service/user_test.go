package services

import (
	"fmt"
	"testing"
	"sync"
)
import (
	modelDB "example_app/entity/db"
	// modelAPI "example_app/entity/api"
	modelHttp "example_app/entity/http"
)
import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// SETUP
type repositoryDBMock struct {
	mock.Mock
}

func (repository *repositoryDBMock) GetUserByID(id int, userData *modelDB.User, wg *sync.WaitGroup) error {
	repository.Called(id,userData,wg)
	userData.ID = uint(id)
	userData.Name = "Test Name"
	userData.IDCardNumber = "IDCARD123456789"
	userData.Address = "Street Test Number 69"
	wg.Done()
	return nil
}

func (repository *repositoryDBMock) CreateUser(userData *modelDB.User) ([]modelDB.User, error) {
	
	users := []modelDB.User{}
	const (
	    ID1 = iota + 1
	    ID2
	    ID3
	)
	users = append(users, modelDB.User{
		ID: uint(ID1),
		Name: "Test Name ONE",
		IDCardNumber: "IDCARD123456789 ONE",
		Address: "Street Test Number 69 ONE",
	})
	users = append(users, modelDB.User{
		ID: uint(ID2),
		Name: "Test Name TWO",
		IDCardNumber: "IDCARD123456789 TWO",
		Address: "Street Test Number 69 TWO",
	})
	users = append(users, modelDB.User{
		ID: uint(ID3),
		Name: "Test Name THREE",
		IDCardNumber: "IDCARD123456789 THREE",
		Address: "Street Test Number 69 THREE",
	})
	return users, nil
}

func (repository *repositoryDBMock) UpdateUserByID(id int, userData *modelDB.User) error {
	repository.Called(id,userData)
	userData.ID = uint(id)
	userData.Name = fmt.Sprintf("Updated - %s", userData.Name)
	userData.IDCardNumber = fmt.Sprintf("Updated - %s", userData.IDCardNumber)
	userData.Address = fmt.Sprintf("Updated - %s", userData.Address)
	return nil
}

func (repository *repositoryDBMock) DeleteUserByID(id int, userData *modelDB.User) error {
	repository.Called(id,userData)
	userData.ID = uint(id)
	userData.Name = fmt.Sprintf("Updated - %s", userData.Name)
	userData.IDCardNumber = fmt.Sprintf("Updated - %s", userData.IDCardNumber)
	userData.Address = fmt.Sprintf("Updated - %s", userData.Address)
	return nil
}

func (repository *repositoryDBMock) GetUsersList(limit int, offset int) ([]modelDB.User, error) {
	repository.Called(limit,offset)
	users := []modelDB.User{}
	const (
	    ID1 = iota + 1
	    ID2
	    ID3
	)
	users = append(users, modelDB.User{
		ID: uint(ID1),
		Name: "Test Name ONE",
		IDCardNumber: "IDCARD123456789 ONE",
		Address: "Street Test Number 69 ONE",
	})
	users = append(users, modelDB.User{
		ID: uint(ID2),
		Name: "Test Name TWO",
		IDCardNumber: "IDCARD123456789 TWO",
		Address: "Street Test Number 69 TWO",
	})
	users = append(users, modelDB.User{
		ID: uint(ID3),
		Name: "Test Name THREE",
		IDCardNumber: "IDCARD123456789 THREE",
		Address: "Street Test Number 69 THREE",
	})
	return users, nil
}

type repositoryAPIMock struct {
	mock.Mock
}

// func (apiMock *repositoryAPIMock) GetFriendID(id int, friendResponse *modelAPI.FriendResponse, wg *sync.WaitGroup) error{
// 	// apiMock.Called(id, friendResponse, wg)
// 	friendResponse.Data.ID = id
// 	friendResponse.Data.Email = "david.beckham@mail.com"
// 	friendResponse.Data.FirstName = "David"
// 	friendResponse.Data.LastName = "Beckham"
// 	friendResponse.Data.Avatar = "https://media.gq.com/photos/56e853e7161e63486d04d6c8/16:9/w_1280%2Cc_limit/david-beckham-gq-0416-2.jpg"
// 	wg.Done()
// 	return nil
// }

// TEST
// func TestUserServiceGetUserByIDMocked(t *testing.T) {
// 	t.Parallel()
// 	dbMockData := repositoryDBMock{}
// 	apiMockData := repositoryAPIMock{}
// 	wg := &sync.WaitGroup{}
	
// 	user := &modelDB.User{}
// 	friend := &modelAPI.FriendResponse{}
// 	var testId int = 1
// 	dbMockData.On("GetUserByID", testId, user, wg).Return(nil)
// 	apiMockData.On("GetFriendID", testId, friend, wg).Return(nil)
	
// 	userService := UserService{&dbMockData,&apiMockData}
// 	resultFuncService := userService.GetUserByID(testId, wg)
// 	// fmt.Println(user)
// 	// fmt.Println(friend)
// 	// fmt.Println(*resultFuncService.Avatar)
// 	assert.Equal(t, uint(testId), resultFuncService.ID, "It should be same ID")
// 	assert.Equal(t, "Test Name", resultFuncService.Name, "It should be same Name")
// 	assert.Equal(t, "IDCARD123456789", resultFuncService.IDCardNumber, "It should be same IDCardNumber")
// 	assert.Equal(t, "Street Test Number 69", resultFuncService.Address, "It should be same Address")
// 	assert.Equal(t, "https://media.gq.com/photos/56e853e7161e63486d04d6c8/16:9/w_1280%2Cc_limit/david-beckham-gq-0416-2.jpg", *resultFuncService.Avatar, "It should be same Avatar")
// }

func TestUserServiceGetAllUserMocked(t *testing.T) {
	t.Parallel()
	dbMockData := repositoryDBMock{}
	limit := 1
	offset := 3
	dbMockData.On("GetUsersList", limit, offset).Return([]modelDB.User{}, nil)
	userService := UserService{&dbMockData}
	resultFuncService := userService.GetAllUser(limit, offset)
	assert.Equal(t, len(resultFuncService), 3, "It should be same length as Mock Data")
	assert.Equal(t, resultFuncService[0].Name, "Test Name ONE", "It should be same NAME as Mock Data")
	assert.Equal(t, resultFuncService[1].Name, "Test Name TWO", "It should be same NAME as Mock Data")
	assert.Equal(t, resultFuncService[2].Name, "Test Name THREE", "It should be same NAME as Mock Data")
}

func TestUserServiceUpdateUserByIDMocked(t *testing.T) {
	t.Parallel()
	dbMockData := repositoryDBMock{}
	var testId int = 1
	dbMockData.On("UpdateUserByID", testId, &modelDB.User{
		Name: "Test Update",
		IDCardNumber: "IDCARDUPDATE1213243",
		Address: "Adress Update 96",
	}).Return(nil)
	userService := UserService{&dbMockData}
	resultFuncService := userService.UpdateUserByID(testId, modelHttp.UserRequest{
		Name: "Test Update",
		IDCardNumber: "IDCARDUPDATE1213243",
		Address: "Adress Update 96",
	})
	assert.Equal(t, resultFuncService, nil, "It should be true")
}