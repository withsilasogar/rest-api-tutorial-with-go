package handler

import (
	"net/http"
	"test01/internals/services"
	"test01/x/interfacesx"
	"test01/x/paseto"
	"time"

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

// @Summary Creating a new user
// @Description This endpoint is for creating new users
// @Tags auth
// @Accept json
// @Produce json
// @param user body interfacesx.UserRegistrationRequest true "User object"
// @Success 200 {object} interfacesx.UserResponse
// @Router /register [post]
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

	//Generate  A token
	//In a real world project, use a symetric key from your env
	tokenGenerator, err := paseto.NewPasetoGenerator("0123456789abcdef0123456789abcdef")
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})

		return
	}

	// In a real world project, this value should as well come from your environment variables
	duration, err := time.ParseDuration("24h")
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})

		return
	}

	token, payload, err := tokenGenerator.GenerateToken(*userData, duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})

		return
	}

	// c.JSON(http.StatusOK, interfacesx.UserResponse{
	// 	Message: "User created successfully",
	// 	Status:  interfacesx.StatusSuccess,
	// 	Code:    http.StatusOK,
	// 	Data:    *userData,
	// })

	c.JSON(http.StatusOK, interfacesx.LoginResponse{
		Token:     token,
		ExpiresAt: payload.ExpiresAt,
		User:      payload.User,
	})
}

// @Summary Gets a new user
// @Description This endpoint is for fetching a new users
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} interfacesx.UserResponse
// @Param email path string true "Users Email"
// @Router /{email} [get]
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
