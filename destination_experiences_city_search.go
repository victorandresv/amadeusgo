package amadeusgo

import "encoding/json"

const (
	endpointCitySearch                            string = "/v1/reference-data/locations/cities"
	DestinationExperiencesCitySearchParamAirports string = "AIRPORTS"
)

type DestinationExperiencesCitySearchRequest struct {
	Keyword     string   `url:"keyword"`
	CountryCode string   `url:"countryCode,omitempty"`
	Max         int      `url:"max,omitempty"`
	Include     []string `url:"include,omitempty"`
}

type DestinationExperiencesCitySearchResults struct {
	Meta struct {
		Count int `json:"count"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"meta"`
	Data []struct {
		Type    string `json:"type"`
		SubType string `json:"subType"`
		Name    string `json:"name"`
		Address struct {
			CountryCode string `json:"countryCode"`
			StateCode   string `json:"stateCode"`
		} `json:"address"`
		GeoCode struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"geoCode"`
	} `json:"data"`
	Included struct {
		Airports map[string]struct {
			Name     string `json:"name"`
			IataCode string `json:"iataCode"`
			SubType  string `json:"subType"`
		} `json:"airports"`
	} `json:"included"`
}

// DestinationExperiencesCitySearch Search city autocomplete list
func DestinationExperiencesCitySearch(params DestinationExperiencesCitySearchRequest) (*DestinationExperiencesCitySearchResults, string, error) {
	var resultsModel *DestinationExperiencesCitySearchResults
	url := apiAmadeusPrepareCall(params, endpointCitySearch)
	result, status := apiAmadeusCall(url, nil)
	if status == 200 {
		err := json.Unmarshal(result, &resultsModel)
		if err != nil {
			return nil, "", err
		}
	}
	return resultsModel, string(result), nil
}
