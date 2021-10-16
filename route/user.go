package route

import (
	"mini-project/controller"
	"mini-project/middleware"

	"github.com/labstack/echo"
	// "github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	// func NewUser(app *echo.Echo) {
	// app.GET("/user", controller.GetAllUser, middleware.JWTAuthMiddleware)
	// app.POST("/user", controller.CreateUserController, middleware.JWTAuthMiddleware)
	// app.GET("/user/:id", controller.GetUserByIDController, middleware.JWTAuthMiddleware)
	// app.DELETE("/user/:id", controller.DeleteUserByIDController, middleware.JWTAuthMiddleware)
	// app.PUT("/user/:id", controller.UpdateUserByIDController, middleware.JWTAuthMiddleware)
	// app.POST("/register", controller.CreateUser)
	// app.PUT("/user/:id", controller.UpdateUser, middleware.JWTAuth)
	// app.GET("/user/:id", controller.GetUserByID, middleware.JWTAuth)
	// app.POST("/login", controller.UserLogin)
	e.POST("/register", controller.CreateUser)
	e.PUT("/user/:id", controller.UpdateUser, middleware.JWTAuth)
	e.GET("/user/:id", controller.GetUserByID, middleware.JWTAuth)
	e.POST("/login", controller.UserLogin)
	// }

	// func NewTempatWisata(app *echo.Echo) {
	// app.GET("/tempatWisata", controller.GetAllTempatWisata)
	// app.POST("/tempatWisata", controller.CreateTempatWisataController)
	// app.GET("/tempatWisata/:id", controller.GetTempatWisataByIDController)
	// app.DELETE("/tempatWisata/:id", controller.DeleteTempatWisataByIDController)
	// app.PUT("/tempatWisata/:id", controller.UpdateTempatWisataByIDController)
	e.GET("/tempatWisata", controller.GetAllTempatWisata)
	e.POST("/tempatWisata", controller.CreateTempatWisata, middleware.JWTAuth)
	e.GET("/tempatWisata/:id", controller.GetTempatWisataByID)
	e.DELETE("/tempatWisata/:id", controller.DeleteTempatWisataByID, middleware.JWTAuth)
	e.PUT("/tempatWisata/:id", controller.UpdateTempatWisata, middleware.JWTAuth)
	// }
	// func NewTransaction(app *echo.Echo) {
	// app.GET("/transaction", controller.GetAllTransaction)
	// app.POST("/transaction", controller.CreateTransactionController)
	// app.GET("/transaction/:id", controller.GetTransactionByIDController)
	// app.DELETE("/transaction/:id", controller.DeleteTransactionByIDController)
	// app.PUT("/transaction/:id", controller.UpdateTransactionByIDController)
	e.GET("/transaction", controller.GetAllTransaction, middleware.JWTAuth)
	e.GET("/transaction/:id", controller.GetTransactionByID, middleware.JWTAuth)
	e.POST("/transaction", controller.CreateTransaction, middleware.JWTAuth)
	e.DELETE("transaction", controller.DeleteTransactionByID, middleware.JWTAuth)
	e.PUT("transaction/:id", controller.UpdateTransaction, middleware.JWTAuth)
	// }
	// func NewCategory(app *echo.Echo) {
	// app.GET("/category", controller.GetAllCategory)
	// app.POST("/category", controller.CreateCategoryController)
	// app.GET("/category/:id", controller.GetCategoryByIDController)
	// app.DELETE("/category/:id", controller.DeleteCategoryByIDController)
	// app.PUT("/category/:id", controller.UpdateCategoryByIDController)
	e.GET("/category", controller.GetAllCategory)
	e.POST("/category", controller.CreateCategory)
	e.PUT("/category", controller.CreateCategory)

	return e
}
