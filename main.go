package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	// "os"
	"os/user"
	// "waiig/repl"
	"waiig/web/webrepl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s\n", u.Username)
	server := echo.New()

	server.GET("/", webrepl.HandleIndex)
	server.POST("/", webrepl.HandleEvaluate)

	server.Start(":8080")

	// repl.Start(os.Stdin, os.Stdout)
}
