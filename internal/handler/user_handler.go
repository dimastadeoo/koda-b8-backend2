package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
			"token" : "hello",
			"fullname": user.Fullname,
			"email":    user.Email,
		},
	})
}

func (h *UserHandler) FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Id Not Valid",
		})
		return
	}
	user,err := h.svc.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "sukses get data",
			Results: user,
		})
	}

}

func (h *UserHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Id Not Valid",
		})
		return
	}
	var form models.UpdateUser
	err = ctx.ShouldBind(&form)

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
	updateUser, err := h.svc.Update(id, &form)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: fmt.Sprintf("user %s success updated", updateUser.Email),
			Results: updateUser,
		})
	}

}

func (h *UserHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Id Not Valid",
		})
		return
	}

	// fullname := ctx.PostForm("fullname")
	// email := ctx.PostForm("email")
	// password := ctx.PostForm("password")
	err = h.svc.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "User Success Deleted",
		})
	}

}




