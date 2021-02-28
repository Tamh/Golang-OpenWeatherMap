package openweathermap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// OpenWeatherMap configuration variables for the module
type OpenWeatherMap struct {
	APIKey string
	Units  string
}

// City contains the city id and name info
type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Coord contains the location coordinates
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// Weather contains the weather general status
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Wind contains the wind info
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

// Clouds contains the cloud info
type Clouds struct {
	All int `json:"all"`
}

// Rain contains the rain forecast
type Rain struct {
	Threehr int `json:"3h"`
}

// Main contains the main data
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
}

// Sys contains other system or ephemerides data
type Sys struct {
	Type    int     `json:"type"`
	ID      int     `json:"ID"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int64   `json:"sunrise"`
	Sunset  int64   `json:"sunset"`
}

/*
Define API response objects (compose of the above fields)
*/

// CurrentWeatherResponse contains the response from the server
type CurrentWeatherResponse struct {
	Coord      `json:"coord"`
	Weather    []Weather `json:"weather"`
	Main       `json:"main"`
	Wind       `json:"wind"`
	Rain       `json:"rain"`
	Clouds     `json:"clouds"`
	Base       string `json:"base"`
	Visibility int    `json:"visibility"`
	Dt         int    `json:"dt"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Sys        `json:"sys"`
}

// ForecastResponse contains the forecast info
type ForecastResponse struct {
	City    `json:"city"`
	Coord   `json:"coord"`
	Country string `json:"country"`
	List    []struct {
		Dt      int `json:"dt"`
		Main    `json:"main"`
		Weather `json:"weather"`
		Clouds  `json:"clouds"`
		Wind    `json:"wind"`
	} `json:"list"`
}

const (
	// APIURL URL for the api
	APIURL string = "api.openweathermap.org/data/2.5"
)

func makeAPIRequest(url string) ([]byte, error) {
	// Build an http client so we can have control over timeout
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	res, getErr := client.Get(url)
	if getErr != nil {
		return nil, getErr
	}

	// defer the closing of the res body
	defer res.Body.Close()

	// read the http response body into a byte stream
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}

// CurrentWeatherFromCity returns the current weather in a freeform city
func (owm *OpenWeatherMap) CurrentWeatherFromCity(city string) (*CurrentWeatherResponse, error) {
	if owm.APIKey == "" {
		// No API keys present, return error
		return nil, errors.New("No API keys present")
	}
	var addToQuery = ""
	if owm.Units != "" {
		addToQuery = "&units=" + owm.Units
	}

	url := fmt.Sprintf("http://%s/weather?q=%s%s&APPID=%s", APIURL, city, addToQuery, owm.APIKey)

	body, err := makeAPIRequest(url)
	if err != nil {
		return nil, err
	}
	var cwr CurrentWeatherResponse

	// unmarshal the byte stream into a Go data type
	jsonErr := json.Unmarshal(body, &cwr)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &cwr, nil
}

// CurrentWeatherFromCoordinates returns the current weather in geographical coordinates
func (owm *OpenWeatherMap) CurrentWeatherFromCoordinates(lat, long float64) (*CurrentWeatherResponse, error) {
	if owm.APIKey == "" {
		// No API keys present, return error
		return nil, errors.New("No API keys present")
	}
	var addToQuery = ""
	if owm.Units != "" {
		addToQuery = "&units=" + owm.Units
	}

	url := fmt.Sprintf("http://%s/weather?lat=%f&lon=%f%s&APPID=%s", APIURL, lat, long, addToQuery, owm.APIKey)

	body, err := makeAPIRequest(url)
	if err != nil {
		return nil, err
	}

	var cwr CurrentWeatherResponse

	// unmarshal the byte stream into a Go data type
	jsonErr := json.Unmarshal(body, &cwr)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &cwr, nil
}

// CurrentWeatherFromZip returns the current weather in a zipcode
func (owm *OpenWeatherMap) CurrentWeatherFromZip(zip int) (*CurrentWeatherResponse, error) {
	if owm.APIKey == "" {
		// No API keys present, return error
		return nil, errors.New("No API keys present")
	}
	var addToQuery = ""
	if owm.Units != "" {
		addToQuery = "&units=" + owm.Units
	}

	url := fmt.Sprintf("http://%s/weather?zip=%d%s&APPID=%s", APIURL, zip, addToQuery, owm.APIKey)

	body, err := makeAPIRequest(url)
	if err != nil {
		return nil, err
	}
	var cwr CurrentWeatherResponse

	// unmarshal the byte stream into a Go data type
	jsonErr := json.Unmarshal(body, &cwr)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &cwr, nil
}

// CurrentWeatherFromCityID returns the current weather in a city id
func (owm *OpenWeatherMap) CurrentWeatherFromCityID(id int) (*CurrentWeatherResponse, error) {
	if owm.APIKey == "" {
		// No API keys present, return error
		return nil, errors.New("No API keys present")
	}
	var addToQuery = ""
	if owm.Units != "" {
		addToQuery = "&units=" + owm.Units
	}

	url := fmt.Sprintf("http://%s/weather?id=%d%s&APPID=%s", APIURL, id, addToQuery, owm.APIKey)

	body, err := makeAPIRequest(url)
	if err != nil {
		return nil, err
	}
	var cwr CurrentWeatherResponse

	// unmarshal the byte stream into a Go data type
	jsonErr := json.Unmarshal(body, &cwr)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &cwr, nil
}
