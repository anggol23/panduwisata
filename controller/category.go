package controller

import (
	"mini-project/database"
	"mini-project/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllCategory(c echo.Context) error {
	category := database.GetAllCategory()

	return c.JSON(http.StatusOK, echo.Map{
		"message": "berhasil",
		"data":    category,
	})
}

func GetCategoryByIDController(c echo.Context) error {
	id := c.Param("id")
	category := database.GetCategoryByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetCategoryByIDController",
		"data":    category,
	})
}

func DeleteCategoryByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteCategoryByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteCategoryByIDController",
	})
}

func UpdateCategoryByIDController(c echo.Context) error {
	id := c.Param("id")

	var category model.Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateCategoryController",
			"error":   err.Error(),
		})
	}
	database.UpdateCategoryByID(id, category)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetCategoryByIDController",
		"data":    category,
	})
}

func CreateCategoryController(c echo.Context) error {
	var newCategory model.Category
	if err := c.Bind(&newCategory); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateCategoryController",
			"error":   err.Error(),
		})
	}

	newCategory = database.CreateCategory(newCategory)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateCategoryController",
		"data":    newCategory,
	})
}
