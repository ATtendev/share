package geo

import (
	"context"
)

// TODO : refactor this function with options pattern
func (g *Geo) AddGeoPoint(ctx context.Context, stauts, sessionID, userID string, c *Point) error {
	if err := g.rd.Do(ctx, "SET", stauts, sessionID, "FIELD", "user_id", userID, "FIELD", "age", 21, "POINT", c.Lat, c.Lon).Err(); err != nil {
		return err // handle error
	}
	return nil
}

func (g *Geo) GetGeoPoint(ctx context.Context, sessionID string) (*Point, error) {
	p := g.rd.Do(ctx, "GET", "fleet", sessionID, "POINT")
	if p.Err() != nil {
		return nil, p.Err() // handle error
	}
	return nil, nil
}

// TODO : refactor this function with options pattern
func (g *Geo) Nearby(ctx context.Context, status, avoidUserID string, c Point, radius int64) ([]Coordinates, error) {
	nearRes := g.rd.Do(ctx, "NEARBY", status, "WHERE", "user_id", "!=", avoidUserID, "POINT", c.Lat, c.Lon, radius)
	if nearRes.Err() != nil {
		return nil, nearRes.Err() // handle error
	}
	return ParsePoints(nearRes.Val())
}
