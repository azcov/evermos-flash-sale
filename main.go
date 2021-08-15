package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	appInit "github.com/azcov/evermos-flash-sale/init"
	"go.uber.org/zap"

	orderHandler "github.com/azcov/evermos-flash-sale/service/order/handler/api"
	orderRepository "github.com/azcov/evermos-flash-sale/service/order/repository"
	orderUsecase "github.com/azcov/evermos-flash-sale/service/order/usecase"

	_ "github.com/azcov/evermos-flash-sale/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var config *appInit.Config
var logger *zap.Logger

func init() {
	// Start pre-requisite app dependencies
	config, logger = appInit.StartAppInit()
}

// @title Evermos Online Store
// @version 1.0

// @BasePath /evermos
func main() {
	// Get PG Conn Instance
	pgDb, err := appInit.ConnectToPGServer(config)
	if err != nil {
		zap.S().Fatal(err)
	}

	// init router
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	router := e.Group("/evermos")

	orderRepo := orderRepository.NewRepository(pgDb)

	orderUc := orderUsecase.NewOrderUsecase(config, orderRepo)

	orderHandler.NewOrderHandler(router, orderUc)

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	go runHTTPHandler(e)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

func runHTTPHandler(e *echo.Echo) {
	if err := e.Start(config.API.Port); err != nil {
		fmt.Println("shutting down the server")
	}
}

