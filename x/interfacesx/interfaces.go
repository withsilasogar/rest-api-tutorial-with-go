package interfacesx

import (
	"test01/internals/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type ResponseStatus string

const (
	StatusSuccess ResponseStatus = "success"
	StatusError   ResponseStatus = "error"
)

type UserRegistrationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"fullName" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type UserData struct {
	ID        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	FullName  string     `json:"fullName"`
	UserRole  model.Role `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
}

type ErrorMessage struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Status  ResponseStatus `json:"status"`
}

type UserResponse struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Status  ResponseStatus `json:"status"`
	Data    UserData       `json:"data"`
}

type RouteDefinition struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}
