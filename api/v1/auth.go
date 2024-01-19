package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ATtendev/share/api/auth"
	"github.com/ATtendev/share/internal/log"
	"github.com/ATtendev/share/internal/uuidx"
	"github.com/ATtendev/share/store/db"
	"github.com/ATtendev/share/store/db/ent"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Response
	Data Token `json:"data"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type SignUp struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *APIV1Service) registerAuthRoutes(pub *echo.Group, priv *echo.Group) {
	pub.POST("/auth/signin", s.SignIn)
	priv.POST("/auth/signout", s.SignOut)
	pub.POST("/auth/signup", s.SignUp)
}

// SignIn godoc
// @Summary Sign-in to share .
// @Tags    auth
// @Accept  json
// @Produce  json
// @Param   body body     SignIn         true "Sign-in object"
// @Success 200  {object} SignInResponse "signin information"
// @Router  /api/v1/auth/signin [POST]
func (s *APIV1Service) SignIn(c echo.Context) error {
	ctx := c.Request().Context()
	signin := SignIn{}
	if err := json.NewDecoder(c.Request().Body).Decode(&signin); err != nil {
		return c.JSON(http.StatusOK, Response{
			Msg:  "Malformatted signin request",
			Code: http.StatusBadRequest,
		})
	}
	user, err := s.storeDB.GetUser(ctx, &db.FindUser{UserName: &signin.Username})
	if err != nil {
		log.Errorf("can't get user %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Incorrect login credentials, please try again",
			Code: http.StatusBadRequest,
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(signin.Password)); err != nil {
		return c.JSON(http.StatusOK, Response{
			Msg:  "Incorrect login credentials, please try again",
			Code: http.StatusBadRequest,
		})
	}
	expireAt := time.Now().Add(auth.AccessTokenDuration)
	accessToken, err := auth.GenerateAccessToken(user.UserName, user.ID.String(), expireAt, []byte(s.Secret))
	if err != nil {
		return c.JSON(http.StatusOK, Response{
			Msg:  fmt.Sprintf("Failed to generate tokens, err: %s", err),
			Code: http.StatusBadRequest,
		})
	}
	if err := s.storeDB.CreateToken(ctx, &db.Token{
		UserID: user.ID,
		Token:  accessToken,
	}); err != nil {
		log.Errorf("Can't save token %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Incorrect login credentials, please try again",
			Code: http.StatusBadRequest,
		})
	}
	resp := SignInResponse{
		Response: Response{
			Code: http.StatusOK,
			Msg:  "Successfully signed in.",
		},
		Data: Token{
			AccessToken: accessToken,
		},
	}
	return c.JSON(http.StatusOK, resp)
}

// SignOut godoc
// @Summary  Sign-out from share.
// @Tags     auth
// @Produce json
// @Security BearerAuth
// @Success  200 {object} Response "Sign-out success"
// @Router   /api/v1/auth/signout [POST]
func (s *APIV1Service) SignOut(c echo.Context) error {
	ctx := c.Request().Context()
	accessToken := findAccessToken(c)
	userID, err := getUserIDFromAccessToken(accessToken, s.Secret)
	if err != nil {
		log.Errorf("Can't parser token %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Can't get token",
			Code: http.StatusBadRequest,
		})
	}
	if err := s.storeDB.DeleteToken(ctx, db.FindToken{
		UserID: uuidx.MustParsePointer(userID),
		Token:  &accessToken,
	}); err != nil {
		log.Errorf("Can't delete token %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "can't get token",
			Code: http.StatusBadRequest,
		})
	}
	return c.JSON(http.StatusOK, Response{
		Msg:  "Successfully signed out.",
		Code: http.StatusOK,
	})
}

// SignUp godoc
// @Summary Sign-up to share.
// @Tags    auth
// @Accept  json
// @Produce json
// @Param   body body     SignUp   true "Sign-up object"
// @Success 200  {object} Response "response information"
// @Router  /api/v1/auth/signup [POST]
func (s *APIV1Service) SignUp(c echo.Context) error {
	ctx := c.Request().Context()
	signup := SignUp{}
	if err := json.NewDecoder(c.Request().Body).Decode(&signup); err != nil {
		return c.JSON(http.StatusOK, Response{
			Msg:  "Malformatted signup request",
			Code: http.StatusBadRequest,
		})
	}
	_, err := s.storeDB.GetUser(ctx, &db.FindUser{UserName: &signup.Username})
	if !ent.IsNotFound(err) {
		log.Errorf("User already exists %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "User already exists",
			Code: http.StatusBadRequest,
		})
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("failed to generate password hash %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "failed to generate password hash",
			Code: http.StatusBadRequest,
		})
	}

	if _, err := s.storeDB.CreateUser(ctx, &db.User{
		UserName:     signup.Username,
		Email:        signup.Email,
		PasswordHash: string(passwordHash),
	}); err != nil {
		log.Errorf("failed to create user %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "failed to create user",
			Code: http.StatusBadRequest,
		})
	}
	return c.JSON(http.StatusOK, Response{
		Msg:  "Successfully signed up.",
		Code: http.StatusOK,
	})
}
