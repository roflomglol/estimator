package maps

import (
	"encoding/json"
	"errors"
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

type result struct {
	distance int32
	time     int32
	err      error
}

// CalculateTimeAndDistance takes
func CalculateTimeAndDistance(origin, dest *Point) (distance, time int32, err error) {
	ch1 := make(chan result)

	go asyncFetchData(fetchDataFromOSM, origin, dest, ch1)

	r := <-ch1

	if r.err != nil {
		log.Println(r.err)
		ch2 := make(chan result)

		go asyncFetchData(fetchDataFromGoogle, origin, dest, ch2)

		r = <-ch2

		if r.err != nil {
			return 0, 0, r.err
		}
	}

	return r.distance, r.time, nil
}

func asyncFetchData(f func(*Point, *Point) (int32, int32, error), o, d *Point, c chan result) {
	distance, time, err := f(o, d)
	r := result{distance: distance, time: time, err: err}

	c <- r
	close(c)
}

func fetchDataFromOSM(origin, dest *Point) (d, t int32, err error) {
	url := buildURLforOSM(origin, dest)
	resp, err := http.Get(url)

	if err != nil {
		return d, t, errors.New("Problems reaching OSM API")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return d, t, errors.New("Problems reading OSM API response body")
	}

	var j = new(osmResponse)

	err = json.Unmarshal(body, &j)

	if err != nil {
		return d, t, errors.New("Problems parsing OSM API response body")
	}

	d = int32(j.Routes[0].Distance)
	t = int32(j.Routes[0].Duration)

	return d, t, nil
}

func fetchDataFromGoogle(origin, dest *Point) (d, t int32, err error) {
	url := buildURLforGoogle(origin, dest)
	resp, err := http.Get(url)

	if err != nil {
		return d, t, errors.New("Problems reaching Google API")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return d, t, errors.New("Problems reading Google API response body")
	}

	var j = new(googleResponse)

	err = json.Unmarshal(body, &j)

	if err != nil {
		return d, t, errors.New("Problems parsing Google API response body")
	}

	d = int32(j.Routes[0].Legs[0].Distance.Value)
	t = int32(j.Routes[0].Legs[0].Duration.Value)

	return d, t, nil
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
