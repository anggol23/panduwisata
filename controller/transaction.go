package controller

import (
	"mini-project/database"
	"mini-project/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllTransaction(c echo.Context) error {
	transaction := database.GetAllTransaction()

	return c.JSON(http.StatusOK, echo.Map{
		"message": "berhasil",
		"data":    transaction,
	})
}

func GetTransactionByIDController(c echo.Context) error {
	id := c.Param("id")
	transaction := database.GetTransactionByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetTransactionByIDController",
		"data":    transaction,
	})
}

func DeleteTransactionByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteTransactionByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteTransactionByIDController",
	})
}

func UpdateTransactionByIDController(c echo.Context) error {
	id := c.Param("id")

	var transaction model.Transaction
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatetransactionController",
			"error":   err.Error(),
		})
	}
	database.UpdateTransactionByID(id, transaction)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetTransactionByIDController",
		"data":    transaction,
	})
}

func CreateTransactionController(c echo.Context) error {
	var newTransaction model.Transaction
	if err := c.Bind(&newTransaction); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateTransactionController",
			"error":   err.Error(),
		})
	}

	newTransaction = database.CreateTransaction(newTransaction)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateTransactionController",
		"data":    newTransaction,
	})
}
