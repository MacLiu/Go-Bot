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
const OPEN_ID_DOC_URL = "https://login.botframework.com/v1/.well-known/openidconfiguration";

// Microsoft Bot Information
const CLIENT_ID = "ca67396c-d766-4cce-a1bc-11d069c738a1";
const SECRET_KEY = "F4k37ui6Bfre5uB2DenLKPs";

/**
 * Authentication struct for the JSON object returned from the auth url.
 */
type Authentication struct {
	TokenType string `json:"token_type"`;
	Expires int `json:"expires_in"`;
	ExtExpires int `json:"ext_expires_in"`
	AccessToken string `json:"access_token"`;
}

/**
 * OpenID Document for the JSON returned from the open doc id url.
 */
type OpenIDDocument struct {
	JWKSUri string `json:"jwks_uri"`;
}

/**
 *  Authentication that contains the access token that is used as the Authorization header of request sent to the
 *  Connector from this Bot.
 */
func getAuthentication() Authentication {
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

	return auth;
}

/**
 * Sends a POST request to the url with the body. Returns the data.
 */
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

/**
 * Jwks uri specifies the location of the document that contains the Bot Connector service's valid signing keys.
 */
func getJWKSUri() string {
	httpClient := http.Client{};
	request, _ := http.NewRequest("GET", OPEN_ID_DOC_URL, nil);
	response, _ := httpClient.Do(request);
	responseData, err := ioutil.ReadAll(response.Body)

	// Check for error
	if err != nil {
		fmt.Print(err.Error());
		os.Exit(1);
	}

	// Parse JSON for jwks_id
	openDocId := OpenIDDocument{};
	json.Unmarshal(responseData, &openDocId);

	return openDocId.JWKSUri;
}
