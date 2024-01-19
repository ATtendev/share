package v1

import (
	"net/http"
	"time"

	"github.com/ATtendev/share/store/db"
	"github.com/ATtendev/share/store/geo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type APIV1Service struct {
	Secret  string
	storeDB *db.Store
	geodb   *geo.Geo
}

// @title                      share API
// @version                    1.0
// @BasePath                   /
// @contact.email              s.vie4m@gmail.com
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @Security                   BearerAuth
func NewAPIV1Service(secret string, storeDB *db.Store, geoDB *geo.Geo) *APIV1Service {
	return &APIV1Service{
		Secret:  secret,
		storeDB: storeDB,
		geodb:   geoDB,
	}
}

func (s *APIV1Service) Register(rootGroup *echo.Group) {
	apiV1GroupPub := rootGroup.Group("/api/v1")
	apiV1GroupPriv := rootGroup.Group("/api/v1")
	rootGroup.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 30, Burst: 100, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}))
	apiV1GroupPriv.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return JWTMiddleware(s, next, s.Secret)
	})
	apiV1GroupPub.GET("/swagger/*", echoSwagger.WrapHandler)
	s.registerUserRoutes(apiV1GroupPub, apiV1GroupPriv)
	s.registerAuthRoutes(apiV1GroupPub, apiV1GroupPriv)
	s.registerSessionRoutes(apiV1GroupPub, apiV1GroupPriv)
}
