package handler

import (
	"fmt"
	"net/http"

	"github.com/dimastadeoo/backend1/internal/lib"
	"github.com/dimastadeoo/backend1/internal/models"
	"github.com/dimastadeoo/backend1/internal/services"
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
)

type UserHandler struct {
	svc *services.UserService
}

func NewHandlerUser(svc *services.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var form models.RegisterUsers

	err := ctx.ShouldBind(&form)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// fullname := ctx.PostForm("fullname")
	// email := ctx.PostForm("email")
	// password := ctx.PostForm("password")



	newUser, err := h.svc.Register(&form)

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

	users, err := h.svc.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
		Success: false,
		Message: err.Error(),
		Results: users,
	})
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Results: users,
	})
}

func (h *UserHandler) Login(ctx *gin.Context) {

	var form models.LoginUser

	err := ctx.ShouldBind(&form)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// email := ctx.PostForm("email")
	// password := ctx.PostForm("password")

	user, err := h.svc.Login(&form)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "login berhasil",
		Results: gin.H{
			"fullname": user.Fullname,
			"email":    user.Email,
		},
	})
}
