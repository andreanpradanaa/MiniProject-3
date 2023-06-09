package admin

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"miniproject2/entities"
	"miniproject2/repositories/mocks"
	"testing"
	"time"
)

func TestCreateAdmin(t *testing.T) {
	adminRepo := new(mocks.AdminInterfaceRepo)
	adminUseCase := useCaseAdmin{
		adminRepo: adminRepo,
	}

	adminParam := AdminParam{
		Username: "testadmin",
		Password: "password",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(adminParam.Password), bcrypt.DefaultCost)
	expectedAdmin := entities.Admin{
		Username: adminParam.Username,
		Password: string(hashedPassword),
		RoleID:   2,
		Verified: "false",
		Active:   "false",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	// Mock the CreateAdmin method
	adminRepo.On("CreateAdmin", mock.AnythingOfType("*entities.Admin")).Return(&expectedAdmin, nil)

	createdAdmin, err := adminUseCase.CreateAdmin(adminParam)

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(createdAdmin.Password), []byte(adminParam.Password))
	assert.NoError(t, err)

	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin.Username, createdAdmin.Username)
	assert.Equal(t, expectedAdmin.RoleID, createdAdmin.RoleID)

}

func TestGetAdminById(t *testing.T) {
	adminRepo := new(mocks.AdminInterfaceRepo)
	useCase := useCaseAdmin{adminRepo}

	adminID := uint(1)
	admin := entities.Admin{
		ID:       adminID,
		Username: "testuser",
		Password: "testpassword",
		RoleID:   2,
		Verified: "false",
		Active:   "false",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	adminRepo.On("GetAdminsById", adminID).Return(admin, nil)

	retrievedAdmin, err := useCase.GetAdminById(adminID)
	assert.NoError(t, err)
	assert.Equal(t, admin, retrievedAdmin)

	adminRepo.AssertCalled(t, "GetAdminsById", adminID)
}

func TestLoginByUsername(t *testing.T) {
	adminRepo := new(mocks.AdminInterfaceRepo)
	useCase := useCaseAdmin{adminRepo}

	username := "testuser"
	admin := entities.Admin{
		Username: "testuser",
		Password: "testpassword",
		RoleID:   2,
		Verified: "false",
		Active:   "false",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	adminRepo.On("LoginByUsername", username).Return(admin, nil)

	retrievedAdmin, err := useCase.LoginByUsername(username)
	assert.NoError(t, err)
	assert.Equal(t, admin, retrievedAdmin)

	adminRepo.AssertCalled(t, "LoginByUsername", username)
}

func TestDeleteAdmin(t *testing.T) {
	adminRepo := new(mocks.AdminInterfaceRepo)
	useCase := useCaseAdmin{adminRepo}

	email := "test@example.com"

	adminRepo.On("DeleteAdmin", email).Return(nil, nil) // Ensure the method is called exactly once

	_, err := useCase.DeleteAdmin(email)
	assert.NoError(t, err)

	adminRepo.AssertCalled(t, "DeleteAdmin", email)
	adminRepo.AssertNumberOfCalls(t, "DeleteAdmin", 1) // Ensure exactly 1 call was made
}

func TestUpdateAdmin(t *testing.T) {
	adminRepo := new(mocks.AdminInterfaceRepo)
	useCase := useCaseAdmin{adminRepo}

	adminID := uint(1)
	adminParam := AdminParam{
		Username: "testuser",
		Password: "testpassword",
	}

	editAdmin := entities.Admin{
		ID:       adminID,
		Username: "testuser",
		Password: "testpassword",
		RoleID:   2,
		Verified: "false",
		Active:   "false",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	adminRepo.On("UpdateAdmin", mock.AnythingOfType("*entities.Admin")).Return(editAdmin, nil) // Ensure the method is called exactly once

	result, err := useCase.UpdateAdmin(adminParam, adminID)
	assert.NoError(t, err)
	assert.Equal(t, editAdmin, result)

}
