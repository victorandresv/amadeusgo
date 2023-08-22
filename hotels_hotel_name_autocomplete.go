package amadeusgo

import (
	"encoding/json"
	"errors"
)

const (
	endpointHotelNameAutocomplete string = "/v1/reference-data/locations/hotel"

	HotelsNameAutocompleteHotelLeisure string = "HOTEL_LEISURE"
	HotelsNameAutocompleteHotelGds     string = "HOTEL_GDS"
)

type HotelsNameAutocompleteRequest struct {
	Keyword     string `url:"keyword"`
	SubType     string `url:"subType"`
	CountryCode string `url:"countryCode,omitempty"`
	Lang        string `url:"lang,omitempty"`
	Max         string `url:"max,omitempty"`
}

type HotelsNameAutocompleteResults struct {
	Data []struct {
		Id        int      `json:"id"`
		Name      string   `json:"name"`
		IataCode  string   `json:"iataCode"`
		SubType   string   `json:"subType"`
		Relevance int      `json:"relevance"`
		Type      string   `json:"type"`
		HotelIds  []string `json:"hotelIds"`
		Address   struct {
			CityName    string `json:"cityName"`
			CountryCode string `json:"countryCode"`
		} `json:"address"`
		GeoCode struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"geoCode"`
	} `json:"data"`
	Meta struct {
		Count int `json:"count"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"meta"`
}

// HotelsHotelNameAutocomplete Search hotel list for autocomplete purpose by keyword
func HotelsHotelNameAutocomplete(params HotelsNameAutocompleteRequest) (resultsModel *HotelsNameAutocompleteResults, err error) {
	url := apiAmadeusPrepareCall(params, endpointHotelNameAutocomplete)
	result, status := apiAmadeusCall(url, nil)
	if status == 200 {
		err := json.Unmarshal(result, &resultsModel)
		if err != nil {
			return nil, errors.New(err.Error())
		}
	}
	return
}
