package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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

// Register godoc
// @Summary      Register a new user
// @Description  Creates a new user account with fullname, email, and password.
// @Tags         auth
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        fullname formData string true "User full name"
// @Param        email    formData string true "User email" Format(email)
// @Param        password formData string true "User password" Format(password)
// @Success      201  {object}  map[string]interface{}  "Returns user data or success message"
// @Failure      400  {object}  map[string]interface{}  "Invalid request payload (e.g., missing fields, invalid email format)"
// @Failure      409  {object}  map[string]interface{}  "Email already exists"
// @Router       /auth/register [POST]
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
	curentUser, _ := ctx.Get("userId")
	if curentUser != nil {
		form.CreatedBy = curentUser.(*int)
	}

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

// GetAll godoc
// @Summary      List All Users
// @Description  requires a valid JWT token to fetch metadata
// @Tags         users
// @Produce      json
// @Accept       x-www-form-urlencoded
// @Success      200  {array}   models.Users
// @Security     Bearer
// @Router       /users [get]
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

// Login godoc
// @Summary      User login
// @Description  Authenticates a user using email and password, returns a JWT token.
// @Tags         auth
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        email    formData string true "User email" Format(email)
// @Param        password formData string true "User password" Format(password)
// @Success      200  {object}  map[string]interface{}  "Returns token and user data"
// @Failure      400  {object}  map[string]interface{}  "Invalid request payload"
// @Failure      401  {object}  map[string]interface{}  "Invalid email or password"
// @Router       /auth/login [POST]
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
			"token":    lib.GeneratedToken(user.Id),
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
	user, err := h.svc.FindById(id)

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

func (h *UserHandler) UpdatePicture(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Id Not Valid",
		})
		return
	}

	file, err := ctx.FormFile("picture")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Validasi extensions

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowed := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}
	if !allowed[ext] {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "upload file only jpg, jpeg, png, webpg extension",
		})
		return
	}

	// validasi size
	const maxSize = 2 << 20
	if file.Size > maxSize {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Maximum file size is 2 MB",
		})
		return
	}

	// Get User
	user, err := h.svc.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User Not Found",
		})
		return
	}

	// save file new img
	filename := fmt.Sprintf("image-picture-%d%d%s", id, time.Now().Unix(), ext)
	path := "uploads/" + filename
	err = ctx.SaveUploadedFile(file, path)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// save path to database
	UpdatePicture, err := h.svc.UpdatePicture(id, &models.UpdatePicture{
		Picture: path,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	// remove file old img before update
	if user.Picture != nil {
		os.Remove(*user.Picture)
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: fmt.Sprintf("user %s success updated picture", UpdatePicture.Email),
	})

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
