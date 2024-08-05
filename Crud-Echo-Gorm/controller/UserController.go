package controller

import (
	"Crud-Echo-Gorm/configuration"
	"Crud-Echo-Gorm/model"
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
	response := map[string]interface{}{
		"data": user,
	}

	return c.JSON(http.StatusOK, response)
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

	existing_user := new(model.User)
	if err := db.First(&existing_user, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existing_user.Name = u.Name
	existing_user.LastName = u.LastName
	existing_user.Email = u.Email
	existing_user.Age = u.Age

	if err := db.Save(&existing_user).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existing_user,
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
	return c.JSON(http.StatusOK, "deleted")
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
	users := []model.User{}
	if err := db.Find(&users).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, data)
	}
	return c.JSON(http.StatusOK, users)
}
