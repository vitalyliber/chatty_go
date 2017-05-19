package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Message struct {
	Name  string `json:"name"`
	Text string `json:"email"`
}


func main() {
	e := echo.New()
	e.GET("/messages", getMessages)
	e.Logger.Fatal(e.Start(":1323"))
}

func getMessages(c echo.Context) error {

	messages := []Message{
		{"Jon", "Hello!"},
		{"Mark", "Wow!"},
	}

	return c.JSON(http.StatusOK, messages)
}