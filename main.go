package main

import (
	"fmt"
	"mini-project/config"
	"mini-project/database"
	"mini-project/middleware"
	"mini-project/route"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func HandlerTest(c echo.Context) error {
	return c.String(http.StatusOK, "berhasil")
}

func HandlerLogin(c echo.Context) error {
	user := struct {
		Email    string
		Password string
	}{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	isValid := database.IsValid(user.Email, user.Password)
	if !isValid {
		return c.String(http.StatusBadRequest, "email atau password salah")
	}

	claims := jwt.MapClaims{}
	claims["userId"] = user.Email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("asdasdasd"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, tokenString)
}

func main() {
	config.InitDB()
	config.InitLog()

	// if os.Getenv("MIGRATE") == "true" {
	// 	fmt.Println("migrating ...")
	// 	config.InitMigration()
	// }

	app := echo.New()
	app.Use(middleware.Log)
	// app.Use(middleware.BasicAuthMiddleware)
	// Recover, Logger, CORS, JWT Auth

	app.GET("/test", HandlerTest)

	app.GET(
		"/",
		func(c echo.Context) error {
			email := c.Get("email")
			return c.String(http.StatusOK, fmt.Sprintf("selamat datang %s", email))
		},
		middleware.JWTAuthMiddleware,
	)

	app.POST("/login", HandlerLogin)
	route.NewUser(app)
	route.NewTempatWisata(app)
	route.NewTransaction(app)
	route.NewCategory(app)
	app.Start(":8080")

}
