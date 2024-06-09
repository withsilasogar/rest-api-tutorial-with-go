package handler

import (
	"net/http"
	"test01/internals/services"
	"test01/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserHandler struct {
	userService services.UserService
	validate    *validator.Validate
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validate:    validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userRequest interfacesx.UserRegistrationRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	if err := h.validate.Struct(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	userData, err := h.userService.CreateUserAccount(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})

		return
	}

	c.JSON(http.StatusOK, interfacesx.UserResponse{
		Message: "User created successfully",
		Status:  interfacesx.StatusSuccess,
		Code:    http.StatusOK,
		Data:    *userData,
	})
}

// Retrieve a user
func (h *UserHandler) GetUser(c *gin.Context) {
	userEmail := c.Param("email")

	if userEmail == "" {
		c.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
			Message: "Email is required",
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	userData, err := h.userService.FetchUserAccount(userEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})

		return
	}

	c.JSON(http.StatusOK, interfacesx.UserResponse{
		Message: "User retrieved successfully",
		Status:  interfacesx.StatusSuccess,
		Code:    http.StatusOK,
		Data:    *userData,
	})
}
