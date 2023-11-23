package main

import (
	"golang-third-assignment/config"
	"golang-third-assignment/controllers"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const PORT = ":8088"

func main() {
	config.ConnectGorm()
	config.StartingApps()

	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				controllers.UpdateData()
			}
		}
	}()
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Aplikasi Buat Baca Status Udara dan Status Air"
		return ctx.String(http.StatusOK, data)
	})
	r.POST("/updateenvirontment/:id", controllers.UpdateEnvirontmentByGorm)

	r.Logger.Fatal(r.Start(PORT))
}
