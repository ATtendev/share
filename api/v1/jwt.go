package v1

import (
	"net/http"
	"strings"

	"github.com/ATtendev/share/api/auth"
	"github.com/ATtendev/share/internal/uuidx"
	"github.com/pkg/errors"

	"github.com/ATtendev/share/store/db"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

const (
	// The key name used to store user id in the context
	// user id is extracted from the jwt token subject field.
	userIDContextKey = "user-id"
)

func extractTokenFromHeader(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", nil
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}

func findAccessToken(c echo.Context) string {
	accessToken, _ := extractTokenFromHeader(c)
	return accessToken
}

// JWTMiddleware validates the access token.
func JWTMiddleware(server *APIV1Service, next echo.HandlerFunc, secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		accessToken := findAccessToken(c)
		if accessToken == "" {
			return c.JSON(http.StatusUnauthorized, Response{
				Msg:  "Missing access token",
				Code: http.StatusUnauthorized,
			})
		}

		userID, err := getUserIDFromAccessToken(accessToken, secret)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, Response{
				Msg:  "Invalid or expired access token",
				Code: http.StatusUnauthorized,
			})
		}
		if !server.storeDB.IsTokenAvailable(ctx, &db.FindToken{
			Token:  &accessToken,
			UserID: uuidx.MustParsePointer(userID),
		}) {
			return c.JSON(http.StatusUnauthorized, Response{
				Msg:  "Invalid or expired access token",
				Code: http.StatusUnauthorized,
			})
		}
		// Stores userID into context.
		c.Set(userIDContextKey, userID)
		return next(c)
	}
}

func getUserIDFromAccessToken(accessToken, secret string) (string, error) {
	claims := &auth.ClaimsMessage{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, errors.Errorf("unexpected access token signing method=%v, expect %v", t.Header["alg"], jwt.SigningMethodHS256)
		}
		if kid, ok := t.Header["kid"].(string); ok {
			if kid == "v1" {
				return []byte(secret), nil
			}
		}
		return nil, errors.Errorf("unexpected access token kid=%v", t.Header["kid"])
	})
	if err != nil {
		return "", errors.Wrap(err, "Invalid or expired access token")
	}
	// We either have a valid access token or we will attempt to generate new access token.
	userID := claims.Subject
	return userID, nil
}
