package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Order struct{}

func (o *Order) Create(ctx *gin.Context) {
	fmt.Println("order created")
}

func (o *Order) List(ctx *gin.Context) {
	fmt.Println("list order")
}

func (o *Order) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Printf("get a order xyz by id %s", id)
}

func (o *Order) UpdateByID(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Printf("update a order by id %s", id)
}

func (o *Order) DeleteByID(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Printf("delete a order by id %s", id)
}
