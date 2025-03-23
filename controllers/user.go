package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/jinzhu/gorm"
	"istudyatuni/kotync/config"
	"istudyatuni/kotync/models"
)

func GetUser(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	userID := c.Get("user").(uint)

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	userID := c.Get("user").(uint)

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := db.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
