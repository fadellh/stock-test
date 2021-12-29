package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"stock/api"
	movieController "stock/api/v1/movie"
	movieService "stock/business/movie"
	movieRepo "stock/modules/movie"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Start() {
	fmt.Println("sratr")
	mr := movieRepo.NewApiRepo()
	ms := movieService.NewService(mr)
	mc := movieController.NewController(ms)

	//create echo http
	e := echo.New()

	api.RegisterPath(e, mc)

	go func() {
		// address := fmt.Sprintf("localhost:%d", config.AppPort)
		address := fmt.Sprintf("localhost:6000")
		fmt.Println(address)
		if err := e.Start(address); err != nil {
			fmt.Println(err)
			log.Info("S")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
