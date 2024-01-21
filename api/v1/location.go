package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ATtendev/share/internal/log"
	"github.com/ATtendev/share/internal/uuidx"
	"github.com/ATtendev/share/store/geo"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Position struct {
	UserID   *uuid.UUID `json:"user_id,omitempty"`
	Position Point      `json:"position,omitempty"`
}

type UpdateCurrentPosition struct {
	UserID   uuid.UUID `json:"-"`
	Position Point     `json:"position"` // optional
}

type SearchCurrentPositionResponse struct {
	Response
	Data []Position `json:"data"`
}

func (s *APIV1Service) registerCurrentPositionRoutes(pub *echo.Group, priv *echo.Group) {
	priv.GET("/session/search", s.SearchCurrentPosition)
}

// UpdateCurrentPosition godoc
// @Summary  Update position to share
// @Tags     position
// @Accept   json
// @Produce  json
// @Security BearerAuth
// @Param    body body     UpdateCurrentPosition true "Update current position object"
// @Success  200  {object} Response              "Position information"
// @Router   /api/v1/position [PUT]
func (s *APIV1Service) UpdateCurrentPosition(c echo.Context) error {
	ctx := c.Request().Context()
	// TODO: check if points eq to 0 return error
	currentPosition := UpdateCurrentPosition{}
	if err := json.NewDecoder(c.Request().Body).Decode(&currentPosition); err != nil {
		log.Errorf("malformatted update position request %s", err)
		return c.JSON(http.StatusOK, Response{
			Msg:  "Malformatted update position request",
			Code: http.StatusBadRequest,
		})
	}
	userID, ok := c.Get(userIDContextKey).(string)
	if !ok {
		log.Errorf("can't get user id from context")
		return c.JSON(http.StatusOK, Response{
			Msg:  "Missing auth position",
			Code: http.StatusUnauthorized,
		})
	}
	// TODO: add last position to geoDB
	if err := s.geoDB.AddGeoPoint(ctx, "current", "", userID, &geo.Point{
		Lat: currentPosition.Position.X,
		Lon: currentPosition.Position.Y,
	}); err != nil {
		log.Errorf("can't insert position: %s", err)
	}
	return c.JSON(http.StatusOK, Response{
		Msg:  "Successfully update position.",
		Code: http.StatusBadRequest,
	})
}

// SearchCurrentPosition godoc
// @Summary  Search current position to share
// @Tags     position
// @Accept   json
// @Produce  json
// @Security BearerAuth
// @Param    lat    query    number                        true "Latitude for location-based search"
// @Param    lon    query    number                        true "Longitude for location-based search"
// @Param    radius query    number                        true "Radius for location-based search in meters"
// @Success  200    {object} SearchCurrentPositionResponse "position information"
// @Router   /api/v1/position/search [GET]
func (s *APIV1Service) SearchCurrentPosition(c echo.Context) error {
	ctx := c.Request().Context()
	userID, ok := c.Get(userIDContextKey).(string)
	if !ok {
		log.Errorf("can't get user id from context")
		return c.JSON(http.StatusOK, Response{
			Msg:  "Missing auth position",
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
	points, err := s.geoDB.Nearby(ctx, "current", userID, geo.Point{
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
	resp := SearchCurrentPositionResponse{
		Response: Response{
			Msg:  "Successfully search position.",
			Code: http.StatusOK,
		}}
	for _, p := range points {
		resp.Data = append(resp.Data, Position{
			UserID: uuidx.MustParsePointer(p.UserID),
			Position: Point{
				X: p.Point.Lat,
				Y: p.Point.Lon,
			},
		})
	}
	return c.JSON(http.StatusOK, resp)
}
