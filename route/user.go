package route

import (
	"mini-project/controller"
	"mini-project/middleware"

	"github.com/labstack/echo"
)

func NewUser(app *echo.Echo) {
	app.GET("/user", controller.GetAllUser, middleware.JWTAuthMiddleware)
	app.POST("/user", controller.CreateUserController, middleware.JWTAuthMiddleware)
	app.GET("/user/:id", controller.GetUserByIDController, middleware.JWTAuthMiddleware)
	app.DELETE("/user/:id", controller.DeleteUserByIDController, middleware.JWTAuthMiddleware)
	app.PUT("/user/:id", controller.UpdateUserByIDController, middleware.JWTAuthMiddleware)
}

func NewTempatWisata(app *echo.Echo) {
	app.GET("/tempatWisata", controller.GetAllTempatWisata)
	app.POST("/tempatWisata", controller.CreateTempatWisataController)
	app.GET("/tempatWisata/:id", controller.GetTempatWisataByIDController)
	app.DELETE("/tempatWisata/:id", controller.DeleteTempatWisataByIDController)
	app.PUT("/tempatWisata/:id", controller.UpdateTempatWisataByIDController)
}
func NewTransaction(app *echo.Echo) {
	app.GET("/transaction", controller.GetAllTransaction)
	app.POST("/transaction", controller.CreateTransactionController)
	app.GET("/transaction/:id", controller.GetTransactionByIDController)
	app.DELETE("/transaction/:id", controller.DeleteTransactionByIDController)
	app.PUT("/transaction/:id", controller.UpdateTransactionByIDController)
}
func NewCategory(app *echo.Echo) {
	app.GET("/category", controller.GetAllCategory)
	app.POST("/category", controller.CreateCategoryController)
	app.GET("/category/:id", controller.GetCategoryByIDController)
	app.DELETE("/category/:id", controller.DeleteCategoryByIDController)
	app.PUT("/category/:id", controller.UpdateCategoryByIDController)
}
