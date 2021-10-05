package controller

import (
	"mini-project/database"
	"mini-project/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllTempatWisata(c echo.Context) error {
	developer := database.GetAllTempatWisata()

	return c.JSON(http.StatusOK, echo.Map{
		"message": "berhasil",
		"data":    developer,
	})
}

func GetTempatWisataByIDController(c echo.Context) error {
	id := c.Param("id")
	tempatWisata := database.GetTempatWisataByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetTempatWisataByIDController",
		"data":    tempatWisata,
	})
}

func DeleteTempatWisataByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteTempatWisataByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteTempatWisataByIDController",
	})
}

func UpdateTempatWisataByIDController(c echo.Context) error {
	id := c.Param("id")

	var tempatWisata model.TempatWisata
	if err := c.Bind(&tempatWisata); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateTempatWisataController",
			"error":   err.Error(),
		})
	}
	database.UpdateTempatWisataByID(id, tempatWisata)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetTempatWisataByIDController",
		"data":    tempatWisata,
	})
}

func CreateTempatWisataController(c echo.Context) error {
	var newTempatWisata model.TempatWisata
	if err := c.Bind(&newTempatWisata); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateTempatWisataController",
			"error":   err.Error(),
		})
	}

	newTempatWisata = database.CreateTempatWisata(newTempatWisata)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateTempatWisataController",
		"data":    newTempatWisata,
	})
}
