package googleapis

type GoogleApiClient struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
	GrantType    string
}

var apiClient *GoogleApiClient

func Init(clientId string, clientSecret string, redirectUrl string, grantType string) {
	apiClient = &GoogleApiClient{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		RedirectUri:  redirectUrl,
		GrantType:    grantType,
	}

}

func Get() *GoogleApiClient {
	return apiClient
}
