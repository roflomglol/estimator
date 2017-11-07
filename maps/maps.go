package maps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Point is a struct representing a point on a map with lattitude and longitude
type Point struct {
	Lat  string
	Long string
}

func (p Point) toParamForOSM() string {
	return fmt.Sprintf("%v,%v", p.Long, p.Lat)
}

func (p Point) toParamForGoogle() string {
	return fmt.Sprintf("%v,%v", p.Lat, p.Long)
}

type osmResponse struct {
	Routes []struct {
		Distance float32 `json:"distance"`
		Duration float32 `json:"duration"`
	} `json:"routes"`
}

type googleResponse struct {
	Routes []struct {
		Legs []struct {
			Distance struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"distance"`
			Duration struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"duration"`
		} `json:"legs"`
	} `json:"routes"`
}

// CalculateTimeAndDistance takes
func CalculateTimeAndDistance(origin, dest *Point) (distance, time int32, err error) {
	distance, time, err = fetchDataFromOSM(origin, dest)

	if err != nil {
		log.Println(err)
		distance, time, err = fetchDataFromGoogle(origin, dest)

		if err != nil {
			return
		}
	}

	return
}

func fetchDataFromOSM(origin, dest *Point) (d, t int32, err error) {
	url := buildURLforOSM(origin, dest)
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var j = new(osmResponse)

	err = json.Unmarshal(body, &j)

	if err != nil {
		return
	}

	d = int32(j.Routes[0].Distance)
	t = int32(j.Routes[0].Duration)

	return
}

func fetchDataFromGoogle(origin, dest *Point) (d, t int32, err error) {
	url := buildURLforGoogle(origin, dest)
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var j = new(googleResponse)

	err = json.Unmarshal(body, &j)

	if err != nil {
		return
	}

	d = int32(j.Routes[0].Legs[0].Distance.Value)
	t = int32(j.Routes[0].Legs[0].Duration.Value)

	return
}

func buildURLforOSM(origin, dest *Point) string {
	return fmt.Sprintf(
		"http://osm:5000/route/v1/driving/%v;%v",
		origin.toParamForOSM(),
		dest.toParamForOSM(),
	)
}

func buildURLforGoogle(origin, dest *Point) string {
	return fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/directions/json?mode=driving&origin=%v&destination=%v",
		origin.toParamForGoogle(),
		dest.toParamForGoogle(),
	)
}
