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

type HotelSearchByGeocodeRequestModel struct {
	Latitude    float64  `url:"latitude"`
	Longitude   float64  `url:"longitude"`
	Radius      int      `url:"radius,omitempty"`
	RadiusUnit  string   `url:"radiusUnit,omitempty"`
	ChainCodes  []string `url:"chainCodes,omitempty"`
	Amenities   []string `url:"amenities,omitempty"`
	Ratings     []string `url:"ratings,omitempty"`
	HotelSource []string `url:"hotelSource,omitempty"`
}

type HotelSearchByGeocodeResultsModel struct {
	Data []struct {
		ChainCode string `json:"chainCode"`
		IataCode  string `json:"iataCode"`
		DupeId    int    `json:"dupeId"`
		Name      string `json:"name"`
		HotelId   string `json:"hotelId"`
		GeoCode   struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"geoCode"`
		Address struct {
			CountryCode string `json:"countryCode"`
		} `json:"address"`
		Distance struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"distance"`
	} `json:"data"`
	Meta struct {
		Count int `json:"count"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"meta"`
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
	} `json:"data"`
	Meta struct {
		Count int `json:"count"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"meta"`
}
