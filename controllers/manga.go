package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/jinzhu/gorm"
	"istudyatuni/kotync/config"
	"istudyatuni/kotync/models"
)

func ListManga(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	var manga []models.Manga
	if err := db.Find(&manga).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, manga)
}

func GetManga(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")

	var manga models.Manga
	if err := db.First(&manga, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Manga not found")
	}

	return c.JSON(http.StatusOK, manga)
}

func CreateManga(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	var manga models.Manga
	if err := c.Bind(&manga); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := db.Create(&manga).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, manga)
}

func UpdateManga(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")

	var manga models.Manga
	if err := db.First(&manga, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Manga not found")
	}

	if err := c.Bind(&manga); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := db.Save(&manga).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, manga)
}
