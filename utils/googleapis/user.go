package googleapis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	mgmtError "management-backend/utils/error"
)

type UserInfo struct {
	FamilyName    string `json:"family_name"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	Email         string `json:"email"`
	GivenName     string `json:"given_name"`
	Id            string `json:"id"`
	VerifiedEmail bool   `json:"verified_email"`
}

type UserInfoError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ErrorResponse struct {
	UserInfoError UserInfoError `json:"error"`
}

func (g *GoogleApiClient) FetchSelfInfo(accessToken string) (UserInfo, error) {
	var user UserInfo

	req, _ := http.NewRequest("GET", "https://www.googleapis.com/userinfo/v2/me", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := http.DefaultClient

	resp, err := client.Do(req)

	if err != nil {
		return user, err
	}

	jsonByte, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return user, err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		var errorResponse ErrorResponse
		err = json.Unmarshal(jsonByte, &errorResponse)

		if err != nil {
			return user, err
		}
		return user, mgmtError.Error{
			Code:        errorResponse.UserInfoError.Code,
			ErrorType:   errorResponse.UserInfoError.Status,
			Description: errorResponse.UserInfoError.Message,
		}
	}

	err = json.Unmarshal(jsonByte, &user)

	return user, err
}
