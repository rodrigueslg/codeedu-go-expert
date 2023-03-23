package repository

import (
	"testing"

	"github.com/rodrigueslg/codedu-goexpert/rest-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateUserDBConnection(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	_ = db.AutoMigrate(&entity.User{})
	return db
}

func TestCreateUser(t *testing.T) {
	db := CreateUserDBConnection(t)

	user, err := entity.NewUser("John", "j@j.com", "123456")
	userRepo := NewUserRepository(db)

	err = userRepo.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db := CreateUserDBConnection(t)

	user, err := entity.NewUser("John", "j@j.com", "123456")
	userRepo := NewUserRepository(db)

	err = userRepo.Create(user)
	assert.Nil(t, err)

	userFound, err := userRepo.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}
