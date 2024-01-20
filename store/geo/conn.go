package geo

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
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
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Type   string `json:"type"`
	Point  Point  `json:"coordinates"`
}

func NewConnection() *Geo {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:9851",
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	// SET
	setCmd := redis.NewStringCmd(context.Background(), "SET", "fleet", "truck", "POINT", 33.32, 115.423)
	if err := client.Process(context.Background(), setCmd); err != nil {
		log.Infof("SET command execution failed: %v", err)
	}
	setRes, err := setCmd.Result()
	if err != nil {
		log.Fatalf("Failed to retrieve SET result: %v", err)
	}
	log.Printf("SET: %s", setRes)
	// GET
	getCmd := redis.NewStringCmd(context.Background(), "GET", "fleet", "truck")
	if err := client.Process(context.Background(), getCmd); err != nil {
		log.Fatalf("GET command execution failed: %v", err)
	}
	getRes, err := getCmd.Result()
	if err != nil {
		log.Fatalf("Failed to retrieve GET result: %v", err)
	}
	ParsePoint(getRes)
	log.Printf("GET: %s", getRes)

	return &Geo{
		rd: client,
	}
}

func (g *Geo) Close() error {
	return g.rd.Close()
}

// TODO optimize this function
// TODO add unit test
func ParsePoints(input interface{}) ([]Coordinates, error) {
	// 	[0
	//   [
	//     [bbb5d139-c264-4a27-8258-c3f2b87ca4f6 {"type":"Point","coordinates":[116.423,33.32]}
	//     [userID 77508ecb-2701-4557-a49b-e4efcde1cae7 user_id 77508ecb-2701-4557-a49b-e4efcde1cae7]
	//     ]
	//   ]
	// ]
	v, ok := input.([]interface{})
	if !ok && len(v) < 3 {
		return nil, fmt.Errorf("cannot parse point")
	}

	p := make([]Coordinates, 0, len(input.([]interface{})[1].([]interface{})))
	for _, truckInfo := range input.([]interface{})[1].([]interface{}) {
		point := gjson.Get(truckInfo.([]interface{})[1].(string), "coordinates")
		fields := truckInfo.([]interface{})[2].([]interface{})
		properties := make(map[string]string)
		for i := 0; i < len(fields)-1; i += 2 {
			key := fields[i].(string)
			value := fields[i+1].(string)
			properties[key] = value
		}
		fmt.Println(properties)
		if !point.Exists() || len(point.Array()) < 2 {
			continue
		}
		p = append(p, Coordinates{
			ID:     truckInfo.([]interface{})[0].(string),
			UserID: properties["user_id"],
			Point: Point{
				Lat: point.Array()[1].Float(),
				Lon: point.Array()[0].Float(),
			},
		})
	}
	return p, nil
}

func ParsePoint(input string) (*Coordinates, error) {
	// {\"type\":\"Point\",\"coordinates\":[115.423,33.32]}"}
	fmt.Println(input)
	t := gjson.Get(input, "type")
	if !t.Exists() {
		return nil, fmt.Errorf("cannot parse point")
	}
	point := gjson.Get(input, "coordinates")
	if !point.Exists() || len(point.Array()) < 2 {
		return nil, fmt.Errorf("cannot parse coordinates")
	}
	return &Coordinates{
		Type: t.String(),
		Point: Point{
			Lat: point.Array()[1].Float(),
			Lon: point.Array()[0].Float(),
		},
	}, nil
}
