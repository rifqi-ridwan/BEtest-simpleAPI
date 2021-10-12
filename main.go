package main

import (
	"BEtest-simpleAPI/app/db/mysql"
	h "BEtest-simpleAPI/app/handler"
	"BEtest-simpleAPI/app/routing"

	_ "BEtest-simpleAPI/app/lib/loadenv"

	"github.com/go-rel/rel"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	adapter := mysql.OpenMysqlConnection()
	defer adapter.Close()

	repo := rel.New(adapter)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", h.Login)

	config := middleware.JWTConfig{
		Claims:                  &jwt.StandardClaims{},
		SigningKey:              []byte("secret"),
		ErrorHandlerWithContext: h.JWTErrorChecker,
	}

	r := e.Group("/product", middleware.JWTWithConfig(config))
	routing.RegisterProduct(r, repo)

	e.Logger.Fatal(e.Start(":1323"))
}
