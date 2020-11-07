package user

import (
	"github.com/bata1016/api-practice/db"
	"github.com/bata1016/api-practice/entity"
	"github.com/gin-gonic/gin"
)

// Service procides user's behavior
type Service struct{}

// User is alias of entity.User struct
type User entity.User

// GetAll is get all User
func (service Service) GetAll() ([]User, error) {
	db := db.GetDB()
	var user []User

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (service Service) CreateModel(ctx *gin.Context) {
	db := db.GetDB()
	var user User

	if err := ctx.BindJSON(&user); err != nil {
		return user, err
	}

	if err := db.Create(&user).Error; err != nil {
		return user, nil
	}

	return user, nil

}

// GetByID is get a User
func (service Service) GetByID(id string) (User, error) {
	db := db.GetDB()
	var user User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// updateByID is update a User
func (service Service) UpdateByID(id string, ctx *gin.Context) {
	db := db.GetDB()
	var user User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	if err := ctx.BindJSON(&user); err != nil {
		return user, err
	}

	db.Save(&user)

	return user, nil
}

// DeleteByID is delete a User
func (service Service) DeleteByID(id string) error {
	db := db.GetDB()
	var user User

	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
