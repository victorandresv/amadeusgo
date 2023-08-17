package amadeusgo

var clientId string
var clientSecret string
var oauth2Data Oauth2ResponseModel

func SetCredentials(ClientId string, ClientSecret string) {
	clientId = ClientId
	clientSecret = ClientSecret
}
