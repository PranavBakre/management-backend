package googleapis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	mgmtError "management-backend/utils/error"
)

type TokenResponse struct {
	IdToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expiry       int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

type TokenResponseError struct {
	ErrorType   string `json:"error"`
	Description string `json:"error_description"`
}

func (apiClient *GoogleApiClient) GetToken(accessType string, code string) (TokenResponse, error) {
	var token = TokenResponse{}
	resp, err := http.PostForm("https://oauth2.googleapis.com/token",
		url.Values{
			"code":          {code},
			"client_id":     {apiClient.ClientId},
			"client_secret": {apiClient.ClientSecret},
			"redirect_uri":  {apiClient.RedirectUri},
			"grant_type":    {apiClient.GrantType}})

	if err != nil {
		return token, err
	}

	jsonByte, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		var tokenResponseError TokenResponseError
		err = json.Unmarshal(jsonByte, &tokenResponseError)
		if err != nil {
			return token, err
		}
		return token, mgmtError.Error{
			Code:        resp.StatusCode,
			ErrorType:   tokenResponseError.ErrorType,
			Description: tokenResponseError.Description,
		}
	}

	if err != nil {
		return token, err
	}

	err = json.Unmarshal(jsonByte, &token)

	return token, err
}
