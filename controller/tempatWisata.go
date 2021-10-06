package controller

import (
	"mini-project/database"
	"mini-project/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// func GetAllTempatWisata(c echo.Context) error {
// 	developer := database.GetAllTempatWisata()

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "berhasil",
// 		"data":    developer,
// 	})
// }

// func GetTempatWisataByIDController(c echo.Context) error {
// 	id := c.Param("id")
// 	tempatWisata := database.GetTempatWisataByID(id)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "GetTempatWisataByIDController",
// 		"data":    tempatWisata,
// 	})
// }

// func DeleteTempatWisataByIDController(c echo.Context) error {
// 	id := c.Param("id")
// 	database.DeleteTempatWisataByID(id)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "DeleteTempatWisataByIDController",
// 	})
// }

// func UpdateTempatWisataByIDController(c echo.Context) error {
// 	id := c.Param("id")

// 	var tempatWisata model.TempatWisata
// 	if err := c.Bind(&tempatWisata); err != nil {
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"message": "CreateTempatWisataController",
// 			"error":   err.Error(),
// 		})
// 	}
// 	database.UpdateTempatWisataByID(id, tempatWisata)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "GetTempatWisataByIDController",
// 		"data":    tempatWisata,
// 	})
// }

// func CreateTempatWisataController(c echo.Context) error {
// 	var newTempatWisata model.TempatWisata
// 	if err := c.Bind(&newTempatWisata); err != nil {
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"message": "CreateTempatWisataController",
// 			"error":   err.Error(),
// 		})
// 	}

// 	newTempatWisata = database.CreateTempatWisata(newTempatWisata)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "CreateTempatWisataController",
// 		"data":    newTempatWisata,
// 	})
// }

func GetAllTempatWisata(c echo.Context) error {
	tempatWisata := database.GetAllTempatWisata()
	return c.JSON(http.StatusOK, tempatWisata)
}

func CreateTempatWisata(c echo.Context) error {
	var NewTempatWisata model.TempatWisata
	u := c.Get("user")
	claims := u.(jwt.MapClaims)
	userID := claims["userId"].(float64)
	if err := c.Bind(&NewTempatWisata); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateTempatWisataController",
			"error":   err.Error(),
		})
	}
	NewTempatWisata.UserID = uint(userID)
	NewTempatWisata = database.CreateTempatWisata(NewTempatWisata)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateTempatWisataController",
		"data":    NewTempatWisata,
	})
}

func UpdateTempatWisata(c echo.Context) error {
	id := c.Param("id")

	var TempatWisata model.TempatWisata
	if err := c.Bind(&TempatWisata); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateTempatWisataController",
			"error":   err.Error(),
		})
	}
	database.UpdateTempatWisata(id, TempatWisata)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Update success",
		"data":    TempatWisata,
	})
}
func GetTempatWisataByID(c echo.Context) error {
	id := c.Param("id")
	TempatWisata := database.GetTempatWisataByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetTempatWisataByIDController",
		"data":    TempatWisata,
	})
}

func DeleteTempatWisataByID(c echo.Context) error {
	id := c.Param("id")
	database.DeleteTempatWisataByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteTempatWisataByID",
	})
}
