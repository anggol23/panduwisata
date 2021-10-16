package controller

import (
	"mini-project/database"
	"mini-project/model"
	"net/http"

	// "net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// func GetAllTransaction(c echo.Context) error {
// 	transaction := database.GetAllTransaction()

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "berhasil",
// 		"data":    transaction,
// 	})
// }

// func GetTransactionByIDController(c echo.Context) error {
// 	id := c.Param("id")
// 	transaction := database.GetTransactionByID(id)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "GetTransactionByIDController",
// 		"data":    transaction,
// 	})
// }

// func DeleteTransactionByIDController(c echo.Context) error {
// 	id := c.Param("id")
// 	database.DeleteTransactionByID(id)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "DeleteTransactionByIDController",
// 	})
// }

// func UpdateTransactionByIDController(c echo.Context) error {
// 	id := c.Param("id")

// 	var transaction model.Transaction
// 	if err := c.Bind(&transaction); err != nil {
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"message": "CreatetransactionController",
// 			"error":   err.Error(),
// 		})
// 	}
// 	database.UpdateTransactionByID(id, transaction)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "GetTransactionByIDController",
// 		"data":    transaction,
// 	})
// }

// func CreateTransactionController(c echo.Context) error {
// 	var newTransaction model.Transaction
// 	if err := c.Bind(&newTransaction); err != nil {
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"message": "CreateTransactionController",
// 			"error":   err.Error(),
// 		})
// 	}

// 	newTransaction = database.CreateTransaction(newTransaction)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "CreateTransactionController",
// 		"data":    newTransaction,
// 	})
// }

func GetAllTransaction(c echo.Context) error {
	transactions := database.GetAllTempatWisata()
	return c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c echo.Context) error {
	var newTransaction model.Transaction
	u := c.Get("user")
	claims := u.(jwt.MapClaims)
	userID := claims["userId"].(float64)
	if err := c.Bind(&newTransaction); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Transaction succes",
			"error":   err.Error(),
		})
	}
	newTransaction.UserID = uint(userID)
	newTransaction = database.CreateTransaction(newTransaction)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateTransactionController",
		"data":    newTransaction,
	})
}

func UpdateTransaction(c echo.Context) error {
	id := c.Param("id")

	var transaction model.Transaction
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateTransactionController",
			"error":   err.Error(),
		})
	}
	database.UpdateTransaction(id, transaction)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Update success",
		"data":    transaction,
	})
}
func GetTransactionByID(c echo.Context) error {
	id := c.Param("id")
	transaction := database.GetTransactionByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Transaction Succes",
		"data":    transaction,
	})
}

func DeleteTransactionByID(c echo.Context) error {
	id := c.Param("id")
	database.DeleteTransactionByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteTransactionByID",
	})
}
