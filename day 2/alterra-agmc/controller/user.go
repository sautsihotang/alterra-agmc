package controller

import (
	"alterra-agmc/lib"
	"alterra-agmc/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := new(model.User)
	c.Bind(&user)

	createdUser, err := lib.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"data":    createdUser,
	})
}

func GetUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := lib.GetUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get a user",
		"data":    user,
	})
}

func UpdateUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	// Return http.StatusNotFound if user does not exist
	_, err := lib.GetUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	user := new(model.User)
	c.Bind(&user)

	updatedUser, err := lib.UpdateUser(userId, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update a user",
		"data":    updatedUser,
	})
}

func DeleteUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	if err := lib.DeleteUser(userId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete a user",
	})
}

func GetAllUser(c echo.Context) error {
	users, err := lib.GetAllUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    users,
	})
}
