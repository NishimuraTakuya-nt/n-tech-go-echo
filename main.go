package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		logger.Info("log test", "message")
		return c.JSON(http.StatusOK, "Hello World")
	})

	e.GET("/echo", func(c echo.Context) error {
		logger.Info("log test2", "message")
		message := c.QueryParam("message")
		if message == "" {
			return c.JSON(http.StatusBadRequest, "Please provide a 'message' query parameter")
		}
		return c.JSON(http.StatusOK, map[string]string{"message": message})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
