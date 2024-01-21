package v1

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/ATtendev/share/internal/log"
	"github.com/ATtendev/share/internal/uuidx"
	"github.com/ATtendev/share/store/db"
	"github.com/ATtendev/share/store/geo"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

/*
	// TODO: validate data
	// TODO: sort points by timestamp
*/

type Point struct {
	X float64 `json:"x"` // optional latitude
	Y float64 `json:"y"` // optional longitude
	Z float64 `json:"z"` // optional elevation (if not provided, will be set to 0)
	T int64   `json:"t"` // required timestamp
}

type Session struct {
	ID          uuid.UUID  `json:"id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeleteAt    *time.Time `json:"delete_at,omitempty"`
	UserID      *uuid.UUID `json:"user_id,omitempty"`
	Description *string    `json:"description,omitempty"`
	Title       *string    `json:"title,omitempty"`
	IsShared    *bool      `json:"is_shared,omitempty"`
	Position    []Point    `json:"position,omitempty"`
	IsFinished  *bool      `json:"is_finished,omitempty"`
}

type CreateSession struct {
	UserID      uuid.UUID `json:"-"`
	Description *string   `json:"description"` // optional
	Title       string    `json:"title"`       // required
	Position    []Point   `json:"position"`    // optional
	IsFinished  bool      `json:"is_finished"`
	IsShared    bool      `json:"is_shared"`
}

type UpdateSession struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"-"`
	Description *string   `json:"description"` // optional
	Title       *string   `json:"title"`       // required
	IsShared    *bool     `json:"is_shared"`
}

type UpdateSessionPosition struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"-"`
	Position []Point   `json:"position"` // optional
}

type FinishSession struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"-"`
}

type CreateSessionResponse struct {
	Response
	Data Session `json:"data"`
}

type SearchSessionResponse struct {
	Response
	Data []Session `json:"data"`
}

func (s *APIV1Service) registerSessionRoutes(pub *echo.Group, priv *echo.Group) {
	priv.GET("/session/search/:sessionID", s.SearchSession)
	priv.POST("/session", s.CreateSession)
	priv.PUT("/session", s.UpdateSession)
	priv.PUT("/session/position", s.UpdateSessionPosition)
	priv.POST("/session/finish", s.FinishSession)
	priv.DELETE("/session/:sessionID", s.DeleteSession)
}

// CreateSession godoc
// @Summary  Create session to share
// @Tags     session
// @Accept   json
// @Produce  json
// @Security BearerAuth
// @Param    body body     CreateSession         true "Create session object"
// @Success  200  {object} CreateSessionResponse "session information"
// @Router   /api/v1/session [POST]
func (s *APIV1Service) CreateSession(c echo.Context) error {
	ctx := c.Request().Context()
	session := CreateSession{}
	if err := json.NewDecoder(c.Request().Body).Decode(&session); err != nil {
		log.Errorf("can't decode create session request %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Malformatted create session request",
			Code: http.StatusBadRequest,
		})
	}
	userID, ok := c.Get(userIDContextKey).(string)
	if !ok {
		log.Errorf("can't get user id from context")
		return c.JSON(http.StatusOK, Response{
			Msg:  "Missing auth session",
			Code: http.StatusUnauthorized,
		})
	}
	// TODO: sort point of positions by timestamp
	points := []db.Point{}
	for _, p := range session.Position {
		points = append(points, db.Point{
			X: p.X,
			Y: p.Y,
			Z: p.Z,
			T: p.T,
		})
	}
	result, err := s.storeDB.CreateSession(ctx, &db.Session{
		UserID:      uuidx.MustParse(userID),
		Title:       &session.Title,
		Description: session.Description,
		Position:    points,
		IsFinished:  session.IsFinished,
	})
	if err != nil {
		log.Errorf("can't create session %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Faild to create session",
			Code: http.StatusBadRequest,
		})
	}
	resp := CreateSessionResponse{
		Response: Response{
			Code: http.StatusOK,
			Msg:  "Successfully create session.",
		},
		Data: Session{
			ID: result.ID,
		},
	}
	return c.JSON(http.StatusOK, resp)
}

// UpdateSession godoc
// @Summary  Update session to share
// @Tags     session
// @Accept   json
// @Produce  json
// @Security BearerAuth
// @Param    body body     UpdateSession true "Update session object"
// @Success  200    {object} Response "session information"
// @Router   /api/v1/session [PUT]
func (s *APIV1Service) UpdateSession(c echo.Context) error {
	ctx := c.Request().Context()
	session := UpdateSession{}
	if err := json.NewDecoder(c.Request().Body).Decode(&session); err != nil {
		log.Errorf("malformatted update session request %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Malformatted update session request",
			Code: http.StatusBadRequest,
		})
	}
	userID, ok := c.Get(userIDContextKey).(string)
	if !ok {
		log.Errorf("can't get user id from context")
		return c.JSON(http.StatusOK, Response{
			Msg:  "Missing auth session",
			Code: http.StatusUnauthorized,
		})
	}
	if err := s.storeDB.UpdateSession(ctx, &db.UpdateSession{
		UserID:      uuidx.MustParse(userID),
		ID:          session.ID,
		Description: session.Description,
		Title:       session.Title,
		IsShared:    session.IsShared,
	}); err != nil {
		log.Errorf("can't update session %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Can't update session",
			Code: http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Msg:  "Successfully update session.",
		Code: http.StatusBadRequest,
	})
}

// UpdateSessionPosition godoc
// @Summary  Update session position to share
// @Tags     session
// @Accept   json
// @Produce  json
// @Security BearerAuth
// @Param    body body     UpdateSessionPosition true "Update session position object"
// @Success  200  {object} Response      "session information"
// @Router   /api/v1/session/position [PUT]
func (s *APIV1Service) UpdateSessionPosition(c echo.Context) error {
	ctx := c.Request().Context()
	// TODO: check if points eq to 0 return error
	session := UpdateSessionPosition{}
	if err := json.NewDecoder(c.Request().Body).Decode(&session); err != nil {
		log.Errorf("malformatted update session request %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Malformatted update session request",
			Code: http.StatusBadRequest,
		})
	}
	userID, ok := c.Get(userIDContextKey).(string)
	if !ok {
		log.Errorf("can't get user id from context")
		return c.JSON(http.StatusOK, Response{
			Msg:  "Missing auth session",
			Code: http.StatusUnauthorized,
		})
	}
	// TODO: sort point of positions by timestamp
	points := []db.Point{}
	for _, p := range session.Position {
		points = append(points, db.Point{
			X: p.X,
			Y: p.Y,
			Z: p.Z,
			T: p.T,
		})
	}
	// TODO: add last position to geoDB
	if err := s.geoDB.AddGeoPoint(ctx, "fleet", session.ID.String(), userID, &geo.Point{
		Lat: session.Position[len(session.Position)-1].X,
		Lon: session.Position[len(session.Position)-1].Y,
	}); err != nil {
		log.Errorf("can't insert position: %s", err)
	}
	//
	if err := s.storeDB.UpdateSessionPosition(ctx, &db.UpdatePosition{
		UserID:   uuidx.MustParse(userID),
		ID:       session.ID,
		Position: points,
	}); err != nil {
		log.Errorf("can't update session position %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Can't update session",
			Code: http.StatusBadRequest,
		})
	}
	return c.JSON(http.StatusOK, Response{
		Msg:  "Successfully update session.",
		Code: http.StatusBadRequest,
	})
}

// FinishSession godoc
// @Summary  Finish session position to share
// @Tags     session
// @Accept   json
// @Produce  json
// @Security BearerAuth
// @Param    body body     FinishSession true "Finish session object"
// @Success  200  {object} Response      "session information"
// @Router   /api/v1/session/finish [POST]
func (s *APIV1Service) FinishSession(c echo.Context) error {
	ctx := c.Request().Context()
	session := FinishSession{}
	if err := json.NewDecoder(c.Request().Body).Decode(&session); err != nil {
		log.Errorf("malformatted close session request %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Malformatted close session request",
			Code: http.StatusBadRequest,
		})
	}
	userID, ok := c.Get(userIDContextKey).(string)
	if !ok {
		log.Errorf("can't get user id from context")
		return c.JSON(http.StatusOK, Response{
			Msg:  "Missing auth session",
			Code: http.StatusUnauthorized,
		})
	}
	IsFinished := true
	if err := s.storeDB.UpdateSession(ctx, &db.UpdateSession{
		UserID:     uuidx.MustParse(userID),
		ID:         session.ID,
		IsFinished: &IsFinished,
	}); err != nil {
		log.Errorf("can't close session %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Can't close session",
			Code: http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Msg:  "Successfully close session.",
		Code: http.StatusBadRequest,
	})
}

// DeleteSession godoc
// @Summary  Delete session to share
// @Tags     session
// @Accept   json
// @Produce  json
// @Security BearerAuth
// @Param    sessionID path     string   true "Session ID to be deleted"
// @Success  200       {object} Response "session information"
// @Router   /api/v1/session/{sessionID} [DELETE]
func (s *APIV1Service) DeleteSession(c echo.Context) error {
	ctx := c.Request().Context()
	sessionID := c.Param("sessionID")
	userID, ok := c.Get(userIDContextKey).(string)
	if !ok {
		log.Errorf("can't get user id from context")
		return c.JSON(http.StatusOK, Response{
			Msg:  "Missing auth session",
			Code: http.StatusUnauthorized,
		})
	}
	if err := s.storeDB.DeleteSession(ctx, &db.FinishSession{
		UserID: uuidx.MustParsePointer(userID),
		ID:     uuidx.MustParsePointer(sessionID),
	}); err != nil {
		log.Errorf("can't close session %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Can't close session",
			Code: http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Msg:  "Successfully close session.",
		Code: http.StatusBadRequest,
	})
}

// SearchSession godoc
// @Summary  Search session to share
// @Tags     session
// @Accept   json
// @Produce  json
// @Security BearerAuth
// @Param    lat    query    number   true "Latitude for location-based search"
// @Param    lon    query    number   true "Longitude for location-based search"
// @Param    radius query    number   true "Radius for location-based search in meters"
// @Success  200  {object} Response              "session information"
// @Router   /api/v1/session/search/{sessionID} [GET]
func (s *APIV1Service) SearchSession(c echo.Context) error {
	ctx := c.Request().Context()
	userID, ok := c.Get(userIDContextKey).(string)
	if !ok {
		log.Errorf("can't get user id from context")
		return c.JSON(http.StatusOK, Response{
			Msg:  "Missing auth session",
			Code: http.StatusUnauthorized,
		})
	}
	// TODO: check if points eq to 0 return error
	lat, err := strconv.ParseFloat(c.QueryParam("lat"), 64)
	if err != nil {
		log.Errorf("can't parse lat %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Can't parse lat",
			Code: http.StatusUnauthorized,
		})
	}
	lon, err := strconv.ParseFloat(c.QueryParam("lon"), 64)
	if err != nil {
		log.Errorf("can't parse lon %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Can't parse lon",
			Code: http.StatusUnauthorized,
		})
	}
	radius, err := strconv.ParseInt(c.QueryParam("radius"), 10, 64)
	if err != nil {
		log.Errorf("can't parse radius %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Can't parse radius",
			Code: http.StatusUnauthorized,
		})
	}
	sessionID := c.Param("sessionID")
	if !s.storeDB.IsSessionExist(ctx, &db.FinishSession{ID: uuidx.MustParsePointer(sessionID)}) {
		log.Errorf("can't get session")
		return c.JSON(http.StatusOK, Response{
			Msg:  "Can't get session",
			Code: http.StatusUnauthorized,
		})
	}
	points, err := s.geoDB.Nearby(ctx, "fleet", userID, geo.Point{
		Lat: lat,
		Lon: lon,
	}, radius)
	if err != nil {
		log.Errorf("can't get points %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Can't get points",
			Code: http.StatusUnauthorized,
		})
	}
	resp := SearchSessionResponse{
		Response: Response{
			Msg:  "Successfully search session.",
			Code: http.StatusOK,
		}}
	for _, p := range points {
		resp.Data = append(resp.Data, Session{
			ID:     uuidx.MustParse(p.ID),
			UserID: uuidx.MustParsePointer(p.UserID),
			Position: []Point{
				{
					X: p.Point.Lat,
					Y: p.Point.Lon,
				},
			},
		})
	}
	return c.JSON(http.StatusOK, resp)
}
