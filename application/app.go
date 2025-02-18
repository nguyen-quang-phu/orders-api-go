package application

import (
	"context"
	"fmt"

	"github.com/dreamsofcode-io/orders-api/handler"
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func New() *App {
	app := &App{
		router: loadRouter(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	ordersRouter := a.router.Group(
		"/orders",
	)
	loadOrdresRouter(ordersRouter)
	err := a.router.Run(":3000")
	if err != nil {
		return fmt.Errorf("failed to run http server: %w", err)
	}
	return nil
}

func loadOrdresRouter(routerGroup *gin.RouterGroup) {
	orderHandler := &handler.Order{}
	routerGroup.POST("/", orderHandler.Create)
	routerGroup.GET("/", orderHandler.List)
	routerGroup.GET("/:id", orderHandler.GetByID)
	routerGroup.PUT("/:id", orderHandler.UpdateByID)
	routerGroup.DELETE("/:id", orderHandler.DeleteByID)
}
