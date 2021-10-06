package controller

import (
	"mini-project/database"
	"mini-project/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// func GetAllUser(c echo.Context) error {
// 	user := database.GetAllUser()

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "berhasil",
// 		"data":    user,
// 	})
// }

// func GetUserByIDController(c echo.Context) error {
// 	id := c.Param("id")
// 	user := database.GetUserByID(id)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "GetUserByIDController",
// 		"data":    user,
// 	})
// }

// func DeleteUserByIDController(c echo.Context) error {
// 	id := c.Param("id")
// 	database.DeleteUserByID(id)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "DeleteUserByIDController",
// 	})
// }

// func UpdateUserByIDController(c echo.Context) error {
// 	id := c.Param("id")

// 	var user model.User
// 	if err := c.Bind(&user); err != nil {
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"message": "CreateUserController",
// 			"error":   err.Error(),
// 		})
// 	}
// 	database.UpdateUserByID(id, user)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "GetUserByIDController",
// 		"data":    user,
// 	})
// }

// func CreateUserController(c echo.Context) error {
// 	var newUser model.User
// 	if err := c.Bind(&newUser); err != nil {
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"message": "CreateUserController",
// 			"error":   err.Error(),
// 		})
// 	}

// 	newUser = database.CreateUser(newUser)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "CreateUserController",
// 		"data":    newUser,
// 	})
// }

func GetAllUser(c echo.Context) error {
	user := database.GetAllUser()
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"error":   err.Error(),
		})
	}

	newUser = database.CreateUser(newUser)
	newUser.Password = ""
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateUserController",
		"data":    newUser,
	})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"error":   err.Error(),
		})
	}
	database.UpdateUser(id, user)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Update success",
		"data":    user,
	})
}
func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user := database.GetUserByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetUserByIDController",
		"data":    user,
	})
}

func UserLogin(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := database.IsValid(user.Email, user.Password)

	if err != nil {
		return c.String(http.StatusBadRequest, "email atau password salah")
	}

	claims := jwt.MapClaims{}
	claims["userId"] = user.ID
	claims["userEmail"] = user.Email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("asdasdasd"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "Login success",
		"data":   tokenString,
	})
}
