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
// 
// @Param        fullname formData string true "User full name"
// @Param        email    formData string true "User email" Format(email)
// @Param        password formData string true "User password" Format(password)
// 
// @Success      201  {object}  lib.Response  "Returns user data or success message"
// @Failure      400  {object}  lib.Response  "Invalid request payload (e.g., missing fields, invalid email format)"
// @Failure      409  {object}  lib.Response  "Email already exists"
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

// CreateUser untuk admin (endpoint /users)
// @Summary      Create a new user (require auth login)
// @Description  Admin creates a new user account. Requires authentication.
// @Tags         users
// @Accept       x-www-form-urlencoded
// @Produce      json
// 
// @Param        fullname formData string true "User full name"
// @Param        email    formData string true "User email" Format(email)
// @Param        password formData string true "User password" Format(password)
// 
// @Success      201  {object}  lib.Response
// @Failure      400  {object}  lib.Response
// @Failure      401  {object}  lib.Response
// @Failure      403  {object}  lib.Response
// @Security     Bearer
// @Router       /users [POST]
func (h *UserHandler) RegisterAdmin(ctx *gin.Context) {
	h.Register(ctx)
}

// GetAll godoc
// @Summary      List All Users
// @Description  requires a valid JWT token to fetch metadata
// @Tags         users
// @Produce      json
// @Accept       x-www-form-urlencoded
// 
// @Param		 search[fullname]  query string false "Search by fullname"
// @Param		 search[email] 	   query string false "Search by email"
// @Param        page             query int    false "Page number (default: 1 if limit is specified)" minimum(1)
// @Param        limit            query int    false "Number of items per page (default: 10 if page is specified)" minimum(1)
// 
// @Success      200  {array}   models.Users
// @Failure		 500  {object}  lib.Response
// @Failure		 400  {object}  lib.Response
// @Security     Bearer
// @Router       /users [get]
func (h *UserHandler) GetAll(ctx *gin.Context) {
	search := map[string]string{}

	if fullname := ctx.Query("search[fullname]"); fullname != "" {
		search["fullname"] = fullname
	}

	if email := ctx.Query("search[email]"); email != "" {
		search["email"] = email
	}

	pageStr := ctx.Query("page")
	limitStr := ctx.Query("limit")

	page := 0
	limit := 0
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, lib.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, lib.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
	}

	users, err := h.svc.GetAll(search, page, limit)

	for _, user := range users {
		user.CreatedAt = lib.TimeToWIB(user.CreatedAt)
		user.UpdatedAt = lib.TimeToWIB(user.UpdatedAt)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	if len(users) == 0 {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Data Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Success Get Data",
		Results: users,
	})
}

// Login godoc
// @Summary      User login
// @Description  Authenticates a user using email and password, returns a JWT token.
// @Tags         auth
// @Accept       x-www-form-urlencoded
// @Produce      json
// 
// @Param        email    formData string true "User email" Format(email)
// @Param        password formData string true "User password" Format(password)
// 
// @Success      200  {object}  lib.Response  "Returns token and user data"
// @Failure      400  {object}  lib.Response  "Invalid request payload"
// @Failure      401  {object}  lib.Response  "Invalid email or password"
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

// FindById godoc
// @Summary      Get user by ID
// @Description  Retrieves detailed information of a specific user by their unique ID. Requires authentication.
// @Tags         users
// @Accept       x-www-form-urlencoded
// @Produce      json
// 
// @Param        id   path      string  true  "User ID"
// 
// @Success      200  {object}  models.Users "Returns the user data (password excluded)"
// @Failure      400  {object}  lib.Response  "Invalid ID format"
// @Failure      401  {object}  lib.Response  "Unauthorized (missing or invalid JWT)"
// @Failure      404  {object}  lib.Response  "User not found"
// @Failure      500  {object}  lib.Response  "Internal server error"
// @Security     Bearer
// @Router       /users/{id} [GET]
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
	user.CreatedAt = lib.TimeToWIB(user.CreatedAt)
	user.UpdatedAt = lib.TimeToWIB(user.UpdatedAt)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "sukses get data",
		Results: user,
	})

}

// Update godoc
// @Summary      Update user's fullname and email
// @Description  Updates the fullname and email of an existing user. Requires authentication (JWT) and the user must own the account or have admin privileges.
// @Tags         users
// @Accept       x-www-form-urlencoded
// @Produce      json
// 
// @Param        id        path      string  true  "User ID"
// @Param        fullname  formData  string  true  "New full name"
// @Param        email     formData  string  true  "New email" Format(email)
// 
// @Success      200       {object}  models.Users  "Returns updated user data"
// @Failure      400       {object}  lib.Response  "Invalid request payload"
// @Failure      401       {object}  lib.Response  "Unauthorized"
// @Failure      403       {object}  lib.Response  "Forbidden"
// @Failure      404       {object}  lib.Response  "User not found"
// @Failure      500       {object}  lib.Response  "Internal server error"
// @Security     Bearer
// @Router       /users/{id} [PATCH]
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

// UpdatePicture godoc
// @Summary      Update user profile picture
// @Description  Uploads a new profile picture for the user. The file must be in JPG, JPEG, PNG, or WEBP format with a maximum size of 2 MB. The old picture file will be automatically removed.
// @Tags         users
// @Accept       multipart/form-data
// @Produce      json
// 
// @Param        id      path      string  true  "User ID (numeric)"
// @Param        picture formData  file    true  "Profile picture file (jpg, jpeg, png, webp, max 2MB)"
// 
// @Success      200     {object}  lib.Response  "Returns success message: 'user <email> success updated picture'"
// @Failure      400     {object}  lib.Response  "Invalid ID format, invalid file extension, file too large, or other bad request"
// @Failure      404     {object}  lib.Response  "User not found"
// @Failure      500     {object}  lib.Response  "Internal server error (e.g., file save failed)"
// @Security     Bearer
// @Router       /users/{id}/picture [PATCH]
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

// Delete godoc
// @Summary      Delete user by ID
// @Description  Permanently deletes a user from the system. Requires authentication and proper authorization (user owns the account or admin).
// @Tags         users
// @Accept       x-www-form-urlencoded
// @Produce      json
// 
// @Param        id   path      string  true  "User ID (numeric)"
// @Success      200  {object}  lib.Response  "User Success Deleted"
// 
// @Failure      400  {object}  lib.Response  "Invalid ID format or service error (e.g., user not found)"
// @Failure      401  {object}  lib.Response  "Unauthorized (missing or invalid JWT)"
// @Failure      403  {object}  lib.Response  "Forbidden (not authorized to delete this user)"
// @Failure      404  {object}  lib.Response  "User not found"
// @Failure      500  {object}  lib.Response  "Internal server error"
// @Security     Bearer
// @Router       /users/{id} [DELETE]
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
