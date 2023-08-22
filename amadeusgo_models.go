package amadeusgo

type Oauth2ResponseModel struct {
	Type            string `json:"type"`
	Username        string `json:"username"`
	ApplicationName string `json:"application_name"`
	ClientId        string `json:"client_id"`
	TokenType       string `json:"token_type"`
	AccessToken     string `json:"access_token"`
	ExpiresIn       int    `json:"expires_in"`
	State           string `json:"state"`
	Scope           string `json:"scope"`
}

type HotelSearchRequestModel struct {
	Keyword     string `url:"keyword"`
	SubType     string `url:"subType"`
	CountryCode string `url:"countryCode,omitempty"`
	Lang        string `url:"lang,omitempty"`
	Max         string `url:"max,omitempty"`
}
type HotelSearchResultsModel struct {
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
	}
}

type CitySearchRequestModel struct {
	Keyword     string   `url:"keyword"`
	CountryCode string   `url:"countryCode,omitempty"`
	Max         int      `url:"max,omitempty"`
	Include     []string `url:"include,omitempty"`
}

type CitySearchResultsModel struct {
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
}
