package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	apiv1 "github.com/ATtendev/share/api/v1"
	"github.com/ATtendev/share/config"
	"github.com/ATtendev/share/internal/log"
	"github.com/ATtendev/share/store/db"
	"github.com/ATtendev/share/store/geo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	ctx     context.Context
	e       *echo.Echo
	storeDB *db.Store
	geodb   *geo.Geo
	cfgs    config.Config
}

func NewServer(ctx context.Context, cfgs *config.Config, storeDB *db.Store, geoDB *geo.Geo) (*Server, error) {
	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.HidePort = true

	s := &Server{
		ctx:     ctx,
		e:       e,
		cfgs:    *cfgs,
		storeDB: storeDB,
		geodb:   geoDB,
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","latency":"${latency_human}",` +
			`"method":"${method}","uri":"${uri}",` +
			`"status":${status},"error":"${error}"}` + "\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// Register healthz endpoint.
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service ready.")
	})

	// Register API v1 endpoints.
	rootGroup := e.Group("")
	apiV1Service := apiv1.NewAPIV1Service(s.cfgs.ServerConf.Secret, storeDB, geoDB)
	apiV1Service.Register(rootGroup)

	return s, nil
}

func (s *Server) Start() error {
	log.Infof("Staring share server: %s:%d \n", s.cfgs.ServerConf.Host, s.cfgs.ServerConf.Port)
	return s.e.Start(fmt.Sprintf("%s:%d", s.cfgs.ServerConf.Host, s.cfgs.ServerConf.Port))
}

func (s *Server) Shutdown(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Shutdown echo server
	if err := s.e.Shutdown(ctx); err != nil {
		fmt.Printf("failed to shutdown server, error: %v\n", err)
	}

	// Close database connection
	if err := s.storeDB.Close(); err != nil {
		fmt.Printf("failed to close database, error: %v\n", err)
	}

	fmt.Printf("share stopped properly\n")
}
