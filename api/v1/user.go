package v1

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	ID           string `json:"id"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Nickname     string `json:"nickname"`
	PasswordHash string `json:"-"`
	AvatarURL    string `json:"avatar_url"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username  *string `json:"username"`
	Email     *string `json:"email"`
	Nickname  *string `json:"nickname"`
	Password  *string `json:"password"`
	AvatarURL *string `json:"avatar_url"`
}

func (s *APIV1Service) registerUserRoutes(pub *echo.Group, priv *echo.Group) {
}
