package controller

import (
	"bytes"
	"fmt"
	"sync"
	"time"
	"net/http"
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	httpEntity "example_app/entity/http"
)

type UserServiceMock struct{}

func (service *UserServiceMock) GetUserByID(id int, wg *sync.WaitGroup) *httpEntity.UserDetailResponse {
	t, _ := time.Parse("2006-01-02", "2019-08-10")
	status := "Status Value"
	avatar := "Some Image Url"
	return &httpEntity.UserDetailResponse{
		ID: uint(id),
		Name: "Test Name Mock",
		IDCardNumber: "IDCARDTEST12345",
		Address: "Address Test Street",
		Status: &status,
		Avatar: &avatar,
		CreatedAt:&t,
		UpdatedAt:&t,
	}
}

func (service *UserServiceMock) DeleteUserByID(id int) bool {
	
	return true
}

func (service *UserServiceMock) CreateUser(payload httpEntity.UserRequestNew) ([]httpEntity.UserResponse, error){
	users := []httpEntity.UserResponse{}
	const (
	    ID1 = iota + 1
	    ID2
	    ID3
	)
	users = append(users, httpEntity.UserResponse{
		ID: uint(ID1),
		Name: "Test Name Mock",
		IDCardNumber: "IDCARDTEST12345",
		Address: "Address Test Street",
	})
	users = append(users, httpEntity.UserResponse{
		ID: uint(ID2),
		Name: "Test Name Mock",
		IDCardNumber: "IDCARDTEST12345",
		Address: "Address Test Street",
	})
	users = append(users, httpEntity.UserResponse{
		ID: uint(ID3),
		Name: "Test Name Mock",
		IDCardNumber: "IDCARDTEST12345",
		Address: "Address Test Street",
	})
	return users,nil
}

func (service *UserServiceMock) GetAllUser(page int,count int) []httpEntity.UserResponse{
	users := []httpEntity.UserResponse{}
	const (
	    ID1 = iota + 1
	    ID2
	    ID3
	)
	users = append(users, httpEntity.UserResponse{
		ID: uint(ID1),
		Name: "Test Name Mock",
		IDCardNumber: "IDCARDTEST12345",
		Address: "Address Test Street",
	})
	users = append(users, httpEntity.UserResponse{
		ID: uint(ID2),
		Name: "Test Name Mock",
		IDCardNumber: "IDCARDTEST12345",
		Address: "Address Test Street",
	})
	users = append(users, httpEntity.UserResponse{
		ID: uint(ID3),
		Name: "Test Name Mock",
		IDCardNumber: "IDCARDTEST12345",
		Address: "Address Test Street",
	})
	return users
}

func (service *UserServiceMock) UpdateUserByID(id int, payload httpEntity.UserRequest) error{
	return nil
}

func TestUserGetByIDMock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	c, r, resp := LoadRouterTestMock()

	var idTest uint = 19
	url := "/v1/users/" + fmt.Sprint(idTest)
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(resp, c.Request)
	assert.Equal(http.StatusOK, resp.Code, "Status should be 200")

	res := httpEntity.UserDetailResponse{}
	err := json.Unmarshal([]byte(resp.Body.String()), &res)
	
	assert.Equal(err, nil, "should have no error")
	assert.Equal(res.ID, idTest, "It should be same ID")
	assert.Equal(res.Name, "Test Name Mock", "It should be same Name")
	assert.Equal(res.IDCardNumber, "IDCARDTEST12345", "It should be same IDCardNumber")
	assert.Equal(res.Address, "Address Test Street", "It should be same Address")
	assert.Equal(*res.Status, "Status Value", "It should be same Status")
	assert.Equal(*res.Avatar, "Some Image Url", "It should be same Avatar")
}

func TestGetUserListMock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	c, r, resp := LoadRouterTestMock()

	limit := 1
	offset := 100

	url := "/v1/users?limit="+fmt.Sprint(limit)+"&offset="+fmt.Sprint(offset)
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(resp, c.Request)
	assert.Equal(http.StatusOK, resp.Code, "Status should be 200")

	res := []httpEntity.UserResponse{}
	err := json.Unmarshal([]byte(resp.Body.String()), &res)
	
	assert.Equal(err, nil, "should have no error")
	assert.Equal(len(res)>=0, true, "length must in minimum value")
	assert.Equal(len(res)==3, true, "length value must match with mock data")
}

func TestUpdateUserByIDMock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	c, r, resp := LoadRouterTestMock()

	var idTest uint = 1
	payload := &httpEntity.UserRequest{
		Name:  "Test Name",
		IDCardNumber: "TestNumber12345",
		Address: "Test Address",
	}
	jsonPayload, _ := json.Marshal(payload)
	url := "/v1/users/" + fmt.Sprint(idTest)
	c.Request, _ = http.NewRequest(http.MethodPut, url, bytes.NewBufferString(string(jsonPayload)))
	c.Request.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(resp, c.Request)
	assert.Equal(http.StatusOK, resp.Code, "Status should be 204")
}