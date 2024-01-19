package geo

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/tidwall/gjson"
)

type Geo struct {
	rd *redis.Client
}

type Point struct {
	Lat float64
	Lon float64
}

type Coordinates struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Point Point  `json:"coordinates"`
}

func NewConnection() *Geo {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:9851",
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return &Geo{
		rd: client,
	}
}

func (g *Geo) Close() error {
	return g.rd.Close()
}

// TODO optimize this function
// TODO add unit test
func ParsePoint(input interface{}) ([]Coordinates, error) {
	// [
	// 	0
	// 	[
	// 		[truck_34 {"type":"Point","coordinates":[116.423,34.32]}]
	// 		[truck_1 {"type":"Point","coordinates":[115.423,33.32]}]
	// 		[truck {"type":"Point","coordinates":[115.423,33.32]}]
	// 		[truck_2 {"type":"Point","coordinates":[0,0]}]
	// 	 ]
	// ]
	fmt.Println(input)
	v, ok := input.([]interface{})
	if !ok && len(v) < 2 {
		return nil, fmt.Errorf("cannot parse point")
	}

	p := make([]Coordinates, 0, len(input.([]interface{})[1].([]interface{})))
	for _, truckInfo := range input.([]interface{})[1].([]interface{}) {
		fmt.Println(truckInfo.([]interface{})[1].(string))
		point := gjson.Get(truckInfo.([]interface{})[1].(string), "coordinates")
		if !point.Exists() || len(point.Array()) < 2 {
			continue
		}
		p = append(p, Coordinates{
			ID: truckInfo.([]interface{})[0].(string),
			Point: Point{
				Lat: point.Array()[1].Float(),
				Lon: point.Array()[0].Float(),
			},
		})
	}
	return p, nil
}
