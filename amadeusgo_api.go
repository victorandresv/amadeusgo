package amadeusgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"os"
	"time"
)

// HotelSearchOffers Search hotel availability and prices
func HotelSearchOffers(params HotelSearchOffersRequestModel) (resultsModel *HotelSearchOffersResultsModel, resultsString string, err error) {
	url := apiAmadeusPrepareCall(params, endpointHotelSearchOffers)
	result, status := apiAmadeusCall(url, nil)
	if status == 200 {
		err := json.Unmarshal(result, &resultsModel)
		if err != nil {
			return nil, "", errors.New(err.Error())
		}
		resultsString = string(result)
	}
	return
}

// HotelSearchByGeocode Search hotel list for autocomplete purpose by geocode coordinates
func HotelSearchByGeocode(params HotelSearchByGeocodeRequestModel) (resultsModel *HotelSearchByGeocodeResultsModel, err error) {
	url := apiAmadeusPrepareCall(params, endpointHotelSearchByGeocode)
	result, status := apiAmadeusCall(url, nil)
	if status == 200 {
		err := json.Unmarshal(result, &resultsModel)
		if err != nil {
			return nil, errors.New(err.Error())
		}
	}
	return
}

// HotelSearch Search hotel list for autocomplete purpose by keyword
func HotelSearch(params HotelSearchRequestModel) (resultsModel *HotelSearchResultsModel, err error) {
	url := apiAmadeusPrepareCall(params, endpointHotelSearch)
	result, status := apiAmadeusCall(url, nil)
	if status == 200 {
		err := json.Unmarshal(result, &resultsModel)
		if err != nil {
			return nil, errors.New(err.Error())
		}
	}
	return
}

// Make an Amadeus API call
func apiAmadeusCall(url string, params *string) (result []byte, status int) {
	accessToken := getAmadeusAccessToken()
	var headers = make(map[string]string)
	headers["Authorization"] = "Bearer " + accessToken.AccessToken
	result, status = apiRequest(url, "GET", params, headers)
	return
}

// Convert params object to url params and return the final url
func apiAmadeusPrepareCall(params interface{}, endpoint string) string {
	urlParams, err := query.Values(params)
	if err != nil {
		return ""
	}
	return getBaseUrl() + endpoint + "?" + urlParams.Encode()
}

// Get Amadeus Access Token
func getAmadeusAccessToken() Oauth2ResponseModel {

	expirationLeft := int(time.Now().Unix())

	if oauth2Data.ExpiresIn > 0 {
		expirationLeft = oauth2Data.ExpiresIn - int(time.Now().Unix())
	}

	if oauth2Data.ExpiresIn == 0 || expirationLeft < 60 {
		params := "grant_type=client_credentials"
		params += "&client_id=" + clientId
		params += "&client_secret=" + clientSecret

		var headers = make(map[string]string)
		headers["Content-Type"] = "application/x-www-form-urlencoded"
		data, status := apiRequest(getBaseUrl()+endpointOauth2, "POST", params, headers)

		if status == 200 {
			var oauth2Response Oauth2ResponseModel
			err := json.Unmarshal(data, &oauth2Response)
			if err == nil {
				oauth2Response.ExpiresIn = int(time.Now().Unix()) + oauth2Response.ExpiresIn
				oauth2Data = oauth2Response
			} else {
				fmt.Println(err.Error())
			}
		}
	}

	return oauth2Data
}

func getBaseUrl() string {
	return "https://test.api.amadeus.com"
}

func GetAmadeusFlightsSearch(params string) (result []byte, status int) {
	url := os.Getenv("AMADEUS_BASE_URL") + os.Getenv("AMADEUS_ENDPOINT_FLIGHTS_SEARCH") + "?" + params
	result, status = apiAmadeusCall(url, &params)
	return
}

func GetAmadeusAirportCitySearch(params string) (result []byte, status int) {
	url := os.Getenv("AMADEUS_BASE_URL") + os.Getenv("AMADEUS_ENDPOINT_AIRPORTS_AND_CITIES_SEARCH") + "?" + params
	result, status = apiAmadeusCall(url, &params)
	return
}

func GetAmadeusPoiSearch(params string) (result []byte, status int) {
	url := os.Getenv("AMADEUS_BASE_URL") + os.Getenv("AMADEUS_ENDPOINT_POI_SEARCH") + "?" + params
	result, status = apiAmadeusCall(url, &params)
	return
}
