package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dreamsofcode-io/orders-api/handler"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type App struct {
	router *gin.Engine
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		router: loadRouter(),
		rdb:    redis.NewClient(&redis.Options{}),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	ordersRouter := a.router.Group(
		"/orders",
	)
	loadOrdresRouter(ordersRouter)

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect redis: %w", err)
	}

	fmt.Println("Starting server")

	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router.Handler(),
	}
	err = server.ListenAndServe()
	ch := make(chan error, 1)

	go func() {
		if err != nil {
			ch <- fmt.Errorf("failed to run http server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

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
