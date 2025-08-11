package main

///home/nia/my-pgx---echo-test/unitedbackendprgproject

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//Loger and Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//functions

	e.Static("/", "static")
	e.Logger.Fatal(e.Start(":1323"))
}
