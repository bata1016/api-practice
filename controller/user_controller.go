package user

import (
	"fmt"

	user "github.com/bata1016/api-practice/service"
	"github.com/gin-gonic/gin"
)

// Controller is user controller
type Controller struct{}

// Index action: GET /users
func (pc Controller) Index(ctx *gin.Context) {
	var service user.Service
	pointer, err := service.GetAll()

	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, pointer)
	}
}

// Show action: GET /user/:id
func (pc Controller) Show(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var service user.Service
	pointer, err := service.GetByID(id)
	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, pointer)
	}
}

// Update action: PUT /users/:id
func (pc Controller) Update(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var service user.Service
	pointer, err := service.UpdateByID(id, ctx)
	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, pointer)
	}
}

// Delete action: DELETE /users/:id
func (pc Controller) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var service user.Service

	if err := service.DeleteByID(id); err != nil {
		ctx.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		ctx.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
