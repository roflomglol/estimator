package maps

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Point is a struct representing a point on a map with lattitude and longitude
type Point struct {
	Lat  string
	Long string
}

func (p Point) toParam() string {
	return fmt.Sprintf("%v,%v", p.Lat, p.Long)
}

type directions struct {
	Routes []struct {
		Legs []struct {
			Distance struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			}

			Duration struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			}
		} `json:"legs"`
	} `json:"routes"`
}

var apiKey string

func init() {
	apiKey = getAPIKey()
}

func getAPIKey() string {
	apiKey, exists := os.LookupEnv("API_KEY")

	if !exists {
		log.Fatal(errors.New("API_KEY env variable is not set"))
	}

	return apiKey
}

// CalculateTimeAndDistance takes
func CalculateTimeAndDistance(start, finish *Point) (distance int, duration int) {
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

	distance = d.Routes[0].Legs[0].Distance.Value
	duration = d.Routes[0].Legs[0].Duration.Value

	return distance, duration
}

func buildURL(start, finish *Point) string {
	return fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/directions/json?key=%v&origin=%v&destination=%v",
		apiKey,
		start.toParam(),
		finish.toParam(),
	)

}
