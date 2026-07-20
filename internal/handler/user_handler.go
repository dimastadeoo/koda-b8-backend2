package handler

import (
	"fmt"
	"net/http"

	"github.com/dimastadeoo/backend1/internal/lib"
	"github.com/dimastadeoo/backend1/internal/models"
	"github.com/dimastadeoo/backend1/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *services.UserService
}

func NewHandlerUser(svc *services.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	fullname := ctx.PostForm("fullname")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	newUser, err := h.svc.Register(&models.RegisterUsers{
		Fullname: fullname,
		Email:    email,
		Password: password,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: fmt.Sprintf("user %s created", newUser.Email),
		})
	}

}

func (h *UserHandler) GetAll(ctx *gin.Context) {

	users := h.svc.GetAll()

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Results: users,
	})
}
