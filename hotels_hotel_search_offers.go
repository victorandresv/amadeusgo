package amadeusgo

import (
	"encoding/json"
	"errors"
)

const (
	endpointHotelSearchOffers string = "/v3/shopping/hotel-offers"

	HotelsHotelSearchOffersPaymentPolicyNone      = "NONE"
	HotelsHotelSearchOffersPaymentPolicyGuarantee = "GUARANTEE"
	HotelsHotelSearchOffersPaymentPolicyDeposit   = "DEPOSIT"

	HotelsHotelSearchOffersBoardTypeRoomOnly     = "ROOM_ONLY"
	HotelsHotelSearchOffersBoardTypeBreakfast    = "BREAKFAST"
	HotelsHotelSearchOffersBoardTypeHalfBoard    = "HALF_BOARD"
	HotelsHotelSearchOffersBoardTypeFullBoard    = "FULL_BOARD"
	HotelsHotelSearchOffersBoardTypeAllInclusive = "ALL_INCLUSIVE"
)

type HotelsHotelSearchOffersRequest struct {
	HotelIds           string `url:"hotelIds"`
	Adults             int32  `url:"adults,omitempty"`
	CheckInDate        string `url:"checkInDate,omitempty"`
	CheckOutDate       string `url:"checkOutDate,omitempty"`
	CountryOfResidence string `url:"countryOfResidence,omitempty"`
	RoomQuantity       int32  `url:"roomQuantity,omitempty"`
	PriceRange         string `url:"priceRange,omitempty"`
	Currency           string `url:"currency,omitempty"`
	PaymentPolicy      string `url:"paymentPolicy,omitempty"`
	BoardType          string `url:"boardType,omitempty"`
	IncludeClosed      bool   `url:"includeClosed,omitempty"`
	BestRateOnly       bool   `url:"bestRateOnly,omitempty"`
	Lang               string `url:"lang,omitempty"`
}

type HotelsHotelSearchOffersResults struct {
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

// HotelsHotelSearchOffers Search hotel availability and prices
func HotelsHotelSearchOffers(params HotelsHotelSearchOffersRequest) (resultsModel *HotelsHotelSearchOffersResults, resultsString string, err error) {
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
