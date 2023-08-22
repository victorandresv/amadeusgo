### Unofficial Go SDK for Amadeus API Self Service 

For use this library you need create an account on https://developers.amadeus.com,
after that you need create an app for get your client_id and client_secret on https://developers.amadeus.com/my-apps

#### Get library:
````
go get github.com/victorandresv/amadeusgo
````

### How to use
#### Initialize
````
amadeusgo.SetCredentials("client_id", "client_secret")
````

#### Get a City Search
````
result, data, err := amadeusgo.DestinationExperiencesCitySearch(amadeusgo.DestinationExperiencesCitySearchRequest{
    Keyword: "miami",
    Include: amadeusgo.DestinationExperiencesCitySearchParamAirports,
})
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
    fmt.Println(data)
}
````