package maps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Point is a struct representing a point on a map with lattitude and longitude
type Point struct {
	Lat  string
	Long string
}

func (p Point) toParam() string {
	return fmt.Sprintf("%v,%v", p.Long, p.Lat)
}

type directions struct {
	Routes []struct {
		Distance float32 `json:"distance"`
		Duration float32 `json:"duration"`
	} `json:"routes"`
}

// CalculateTimeAndDistance takes
func CalculateTimeAndDistance(start, finish *Point) (distance int32, time int32) {
	url := buildURL(start, finish)

	resp, err := http.Get(url)
	if err != nil {
		//
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//
	}

	var d = new(directions)

	err = json.Unmarshal(body, &d)

	if err != nil {
		//
	}

	distance = int32(d.Routes[0].Distance)
	time = int32(d.Routes[0].Duration)

	return distance, time
}

func buildURL(start, finish *Point) string {
	return fmt.Sprintf(
		"http://osm:5000/route/v1/driving/%v;%v",
		start.toParam(),
		finish.toParam(),
	)

}
