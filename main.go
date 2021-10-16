package main

import (
	"mini-project/config"
	"mini-project/route"
)

// import (
// 	"mini-project/config"
// 	"mini-project/route"
// )

// 	// isValid := database.IsValid(user.Email, user.Password)
// 	// if !isValid {
// 	// 	return c.String(http.StatusBadRequest, "email atau password salah")
// 	// }

// 	claims := jwt.MapClaims{}
// 	claims["userId"] = user.Email
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte("asdasdasd"))
// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.String(http.StatusOK, tokenString)
// }

func main() {

	// configDB := config.ConfigDB{
	// 	DB_Username: viper.GetString(`database.user`),
	// 	DB_Password: viper.GetString(`database.pass`),
	// 	DB_Host:     viper.GetString(`database.host`),
	// 	DB_Port:     viper.GetString(`database.port`),
	// 	DB_Name:     viper.GetString(`database.name`),
	// }
	config.InitDB()
	config.InitLog()
	e := route.New()
	// e.Use(middleware.LogMiddleware)
	// // middlewares.LogMiddlewares(e)
	// e.Use(middleware.Log)
	e.Start(":8080")

}
