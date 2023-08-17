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