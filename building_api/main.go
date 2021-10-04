package main

import (
	"net/http"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type WebPlusSH struct {
	web *echo.Echo
	sh  *shell.Shell
}

var sh_host string = "localhost:5001"

func NewEngine() *WebPlusSH {
	web := echo.New()
	sh := shell.NewShell(sh_host)
	println("Started ipfs shell on %s", sh_host)
	return &WebPlusSH{web, sh}
}

func main() {

	engine := NewEngine()

	web := *engine.web
	//sh := *engine.sh

	// Middleware
	web.Use(middleware.Logger())
	web.Use(middleware.Recover())

	// Routes
	web.GET("/get", engine.get)
	web.POST("/post", engine.post)

	// Start server
	web.Logger.Fatal(web.Start(":1323"))
}

// GET Handler
func (w *WebPlusSH) get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// POST Handler
func (w *WebPlusSH) post(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
