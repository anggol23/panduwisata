package middleware

import (
	"fmt"
	"mini-project/config"
	"mini-project/model"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	coll := config.DBLog.Database("pandu-wisata").Collection("logs")

	return func(c echo.Context) error {
		data := new(model.User)
		if c.Request().Method != http.MethodGet {
			if err := c.Bind(&data); err != nil {
				fmt.Println(err)
			}
		}

		log := bson.M{
			"time":    time.Now(),
			"method":  c.Request().Method,
			"path":    c.Path(),
			"request": data,
		}
		response := next(c)
		log["response"] = c.Response().Status
		coll.InsertOne(c.Request().Context(), log)
		fmt.Println(log)
		return response
	}
}
