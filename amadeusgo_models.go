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

type HotelSearchOffersRequestModel struct {
	HotelIds           []string `url:"hotelIds"`
	Adults             int32    `url:"adults,omitempty"`
	CheckInDate        string   `url:"checkInDate,omitempty"`
	CheckOutDate       string   `url:"checkOutDate,omitempty"`
	CountryOfResidence string   `url:"countryOfResidence,omitempty"`
	RoomQuantity       int32    `url:"roomQuantity,omitempty"`
	PriceRange         string   `url:"priceRange,omitempty"`
	Currency           string   `url:"currency,omitempty"`
	PaymentPolicy      string   `url:"paymentPolicy,omitempty"`
	BoardType          string   `url:"boardType,omitempty"`
	IncludeClosed      bool     `url:"includeClosed,omitempty"`
	BestRateOnly       bool     `url:"bestRateOnly,omitempty"`
	Lang               string   `url:"lang,omitempty"`
}

type HotelSearchOffersResultsModel struct {
	Data []struct {
		Type  string `json:"type"`
		Hotel struct {
			Type      string  `json:"type"`
			HotelId   string  `json:"hotelId"`
			ChainCode string  `json:"chainCode"`
			DupeId    string  `json:"dupeId"`
			Name      string  `json:"name"`
			CityCode  string  `json:"cityCode"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"hotel"`
		Available bool `json:"available"`
		Offers    []struct {
			Id                  string `json:"id"`
			CheckInDate         string `json:"checkInDate"`
			CheckOutDate        string `json:"checkOutDate"`
			RateCode            string `json:"rateCode"`
			RateFamilyEstimated struct {
				Code string `json:"code"`
				Type string `json:"type"`
			} `json:"rateFamilyEstimated"`
			Room struct {
				Type          string `json:"type"`
				TypeEstimated struct {
					Category string `json:"category"`
					Beds     int    `json:"beds"`
					BedType  string `json:"bedType"`
				} `json:"typeEstimated"`
				Description struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"description"`
			} `json:"room"`
			Guests struct {
				Adults int `json:"adults"`
			} `json:"guests"`
			Price struct {
				Currency   string `json:"currency"`
				Base       string `json:"base"`
				Total      string `json:"total"`
				Variations struct {
					Average struct {
						Base string `json:"base"`
					} `json:"average"`
					Changes []struct {
						StartDate string `json:"startDate"`
						EndDate   string `json:"endDate"`
						Total     string `json:"total"`
					} `json:"changes"`
				} `json:"variations"`
			} `json:"price"`
			Policies struct {
				PaymentType  string `json:"paymentType"`
				Cancellation struct {
					Description struct {
						Text string `json:"text"`
					} `json:"description"`
					Type string `json:"type"`
				} `json:"cancellation"`
			} `json:"policies"`
			Self string `json:"self"`
		} `json:"offers"`
		Self string `json:"self"`
	} `json:"data"`
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
	Included struct {
		Airports map[string]struct {
			Name     string `json:"name"`
			IataCode string `json:"iataCode"`
			SubType  string `json:"subType"`
		} `json:"airports"`
	} `json:"included"`
}
