package main

import (
  // "trout-analyzer-back/database"
  "net/http"
  "github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

func main() {
  e := echo.New()
  e.GET("/", hello)
  e.Logger.Fatal(e.Start(":3000"))
}