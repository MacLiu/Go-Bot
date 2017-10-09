package main

import (
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

//Authentication Urls
const MSA_AAD_AUTH_URL = "https://login.microsoftonline.com/botframework.com/oauth2/v2.0/token";
const AUTH_SCOPE = "https://api.botframework.com/.default";

// Microsoft Bot Information
const CLIENT_ID = "ca67396c-d766-4cce-a1bc-11d069c738a1";
const SECRET_KEY = "F4k37ui6Bfre5uB2DenLKPs";

type Authentication struct {
	TokenType string `json:"token_type"`;
	Expires int `json:"expires_in"`;
	ExtExpires int `json:"ext_expires_in"`
	AccessToken string `json:"access_token"`;
}

func getAccessToken() string {
	// Create body for request
	body := url.Values{}
	body.Add("grant_type", "client_credentials");
	body.Add("client_id", CLIENT_ID);
	body.Add("client_secret", SECRET_KEY);
	body.Add("scope", AUTH_SCOPE);

	// Parse JSON into auth
	raw := sendPostRequest(MSA_AAD_AUTH_URL, body);
	auth := Authentication{};
	json.Unmarshal(raw, &auth);

	return auth.AccessToken;
}

func sendPostRequest(url string, body url.Values) []byte {
	httpClient := http.Client{};
	request, _ := http.NewRequest("POST", url, strings.NewReader(body.Encode()));
	response, _ := httpClient.Do(request)
	responseData, err := ioutil.ReadAll(response.Body)

	// Check for error
	if err != nil {
		fmt.Print(err.Error());
		os.Exit(1);
	}
	return responseData;
}
