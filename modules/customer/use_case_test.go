package customer

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"miniproject2/entities"
	"miniproject2/repositories/mocks"
	"testing"
	"time"
)

func TestCreateCustomer(t *testing.T) {
	// Create a mock repository
	customerRepo := new(mocks.CustomerInterfaceRepo)
	useCase := useCaseCustomer{
		customerRepo: customerRepo,
	}

	// Set up test data
	customerParam := CustomerParam{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Avatar:    "avatar.png",
	}

	expectedCustomer := &entities.Customer{
		Firstname: customerParam.FirstName,
		Lastname:  customerParam.LastName,
		Email:     customerParam.Email,
		Avatar:    customerParam.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Mock the repository method
	customerRepo.On("CreateCustomer", mock.AnythingOfType("*entities.Customer")).Return(expectedCustomer, nil)

	// Call the method being tested
	createdCustomer, err := useCase.CreateCustomer(customerParam)

	// Assert the expectations
	//customerRepo.AssertCalled(t, "CreateCustomer", mock.AnythingOfType("*entities.Customer"))
	assert.NoError(t, err)
	assert.Equal(t, *expectedCustomer, createdCustomer)
}

func TestDeleteCustomer(t *testing.T) {
	// Create a mock repository
	repo := new(mocks.CustomerInterfaceRepo)
	useCase := useCaseCustomer{
		customerRepo: repo,
	}

	// Set up test data
	email := "john.doe@example.com"

	// Mock the repository method
	repo.On("DeleteCustomer", email).Return(nil, nil)

	// Call the method being tested
	result, err := useCase.DeleteCustomer(email)

	// Assert the expectations
	repo.AssertCalled(t, "DeleteCustomer", email)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestUpdateCustomer(t *testing.T) {
	// Create a mock repository
	repo := new(mocks.CustomerInterfaceRepo)
	useCase := useCaseCustomer{
		customerRepo: repo,
	}

	// Set up test data
	id := uint(1)
	customerParam := CustomerParam{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Avatar:    "avatar.png",
	}

	expectedCustomer := &entities.Customer{
		ID:        id,
		Firstname: customerParam.FirstName,
		Lastname:  customerParam.LastName,
		Email:     customerParam.Email,
		Avatar:    customerParam.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Mock the repository method
	repo.On("UpdateCustomer", mock.AnythingOfType("*entities.Customer")).Return(expectedCustomer, nil)

	// Call the method being tested
	result, err := useCase.UpdateCustomer(customerParam, id)

	// Assert the expectations
	assert.NoError(t, err)
	assert.Equal(t, *expectedCustomer, result)
}

func TestGetCustomerById(t *testing.T) {
	// Create a mock repository
	repo := new(mocks.CustomerInterfaceRepo)
	useCase := useCaseCustomer{
		customerRepo: repo,
	}

	// Set up test data
	id := uint(1)
	expectedCustomer := entities.Customer{
		ID:        1,
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		Avatar:    "avatar.png",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Mock the repository method
	repo.On("GetCustomerById", id).Return(expectedCustomer, nil)

	// Call the method being tested
	result, err := useCase.GetCustomerById(id)

	// Assert the expectations
	repo.AssertCalled(t, "GetCustomerById", id)
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, result)
}
