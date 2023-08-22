package amadeusgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"os"
	"time"
)

func HotelSearchOffers(params HotelSearchOffersRequestModel) (resultsModel *HotelSearchOffersResultsModel, err error) {
	url := apiAmadeusPrepareCall(params, amadeusEndpointHotelSearchOffers)
	fmt.Println(url)
	result, status := apiAmadeusCall(url, nil)
	if status == 200 {
		err := json.Unmarshal(result, &resultsModel)
		if err != nil {
			return nil, errors.New(err.Error())
		}
	} else {
		fmt.Println("status")
		fmt.Println(status)
	}
	return
}

func HotelSearchByGeocode(params HotelSearchByGeocodeRequestModel) (resultsModel *HotelSearchByGeocodeResultsModel, err error) {
	url := apiAmadeusPrepareCall(params, amadeusEndpointHotelSearchByGeocode)
	result, status := apiAmadeusCall(url, nil)
	if status == 200 {
		err := json.Unmarshal(result, &resultsModel)
		if err != nil {
			return nil, errors.New(err.Error())
		}
	}
	return
}

func HotelSearch(params HotelSearchRequestModel) (resultsModel *HotelSearchResultsModel, err error) {
	url := apiAmadeusPrepareCall(params, amadeusEndpointHotelSearch)
	result, status := apiAmadeusCall(url, nil)
	if status == 200 {
		err := json.Unmarshal(result, &resultsModel)
		if err != nil {
			return nil, errors.New(err.Error())
		}
	}
	return
}

func CitySearch(params CitySearchRequestModel) (resultsModel *CitySearchResultsModel, err error) {
	url := apiAmadeusPrepareCall(params, amadeusEndpointCitySearch)
	result, status := apiAmadeusCall(url, nil)
	if status == 200 {
		err := json.Unmarshal(result, &resultsModel)
		if err != nil {
			return nil, err
		}
	}
	return
}

func apiAmadeusCall(url string, params *string) (result []byte, status int) {
	accessToken := getAmadeusAccessToken()

	var headers = make(map[string]string)
	headers["Authorization"] = "Bearer " + accessToken.AccessToken
	result, status = apiRequest(url, "GET", params, headers)
	return
}

func apiAmadeusPrepareCall(params interface{}, endpoint string) string {
	urlParams, err := query.Values(params)
	if err != nil {
		return ""
	}
	return getBaseUrl() + endpoint + "?" + urlParams.Encode()
}

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
		data, status := apiRequest(getBaseUrl()+amadeusEndpointOauth2, "POST", params, headers)

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
