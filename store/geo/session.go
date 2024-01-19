package geo

import (
	"context"
)

func (g *Geo) AddGeoPoint(ctx context.Context, id, objectiveID string, c Point) error {
	if err := g.rd.Do(ctx, "SET", "fleet", "truck", "POINT", c.Lat, c.Lon).Err(); err != nil {
		return err // handle error
	}
	return nil
}

func (g *Geo) Nearby(ctx context.Context, id string, c Point, radius float32) ([]Coordinates, error) {
	nearRes := g.rd.Do(ctx, "NEARBY", "fleet", "POINT", c.Lat, c.Lon, radius)
	if nearRes.Err() != nil {
		return nil, nearRes.Err() // handle error
	}
	return ParsePoint(nearRes.Val())
}
