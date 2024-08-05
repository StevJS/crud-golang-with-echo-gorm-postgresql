package controller

import (
	"Crud-Echo-Gorm/configuration"
	"Crud-Echo-Gorm/model"
	"Crud-Echo-Gorm/rest"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateUser(c echo.Context) error {
	u := new(model.User)
	db := configuration.DB()

	if err := c.Bind(u); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	user := &model.User{
		Name:     u.Name,
		LastName: u.LastName,
		Email:    u.Email,
		Age:      u.Age,
	}
	if err := db.Create(&user).Error; err != nil {
		data := map[string]interface{}{
			"message": user,
		}

		return c.JSON(http.StatusInternalServerError, data)

	}

	return c.JSON(http.StatusOK, rest.UserCreateResponse{Message: "user created successfully", User: *user})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	u := new(model.User)
	db := configuration.DB()

	if err := c.Bind(u); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingUser := new(model.User)
	if err := db.First(&existingUser, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existingUser.Name = u.Name
	existingUser.LastName = u.LastName
	existingUser.Email = u.Email
	existingUser.Age = u.Age

	if err := db.Save(&existingUser).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existingUser,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteUserById(c echo.Context) error {
	id := c.Param("id")
	db := configuration.DB()

	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	if err := db.Delete(&user).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	return c.JSON(http.StatusOK, rest.UserDeleteResponse{Message: "User delete"})
}
func GetUserById(c echo.Context) error {
	id := c.Param("id")
	db := configuration.DB()
	var user model.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		data := map[string]interface{}{
			"message": err,
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	return c.JSON(http.StatusOK, user)
}
func GetAllUsers(c echo.Context) error {
	db := configuration.DB()
	var users []model.User
	if err := db.Find(&users).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		err := c.JSON(http.StatusInternalServerError, data)
		if err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, users)
}
