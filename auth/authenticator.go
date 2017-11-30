package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Authentication server settings
var (
	AuthHost = "https://api.toon.eu"
	AuthURL  = AuthHost + "/authorize/legacy"
	TokenURL = AuthHost + "/token"
)

// ToonAuthenticator description
type ToonAuthenticator struct {
	clientID         string
	clientSecret     string
	tenantID         string
	redirectURI      string
	callbackHost     string
	callbackEndpoint string
	TokenError       string
	callbackPort     int
	IsAuthenticating bool
	Token            OAuthToken
	Events           chan AuthenticationEvent
}

// NewToonAuthenticator create a new Toon Authenticator
func NewToonAuthenticator(clientID, clientSecret, tenantID, redirectURI, callbackHost, callbackEndpoint string, callbackPort int) *ToonAuthenticator {
	ta := &ToonAuthenticator{
		clientID:         clientID,
		clientSecret:     clientSecret,
		tenantID:         tenantID,
		redirectURI:      redirectURI,
		callbackHost:     callbackHost,
		callbackEndpoint: callbackEndpoint,
		callbackPort:     callbackPort,
		Events:           make(chan AuthenticationEvent),
	}

	return ta
}

// StartGetToken description
func (auth *ToonAuthenticator) StartGetToken(username, password string) {
	auth.IsAuthenticating = true
	go auth.startCallbackServer()

	// sleep for a moment to be sure the HTTP server started before authentication
	time.Sleep(300 * time.Millisecond)
	go auth.login(username, password)
}

// StartRefreshToken description
func (auth *ToonAuthenticator) StartRefreshToken() {
	auth.IsAuthenticating = true
	auth.Events <- TokenRefreshing

	go func() {
		resp, err := postFormData(TokenURL,
			FormValue{key: "client_id", value: auth.clientID},
			FormValue{key: "client_secret", value: auth.clientSecret},
			FormValue{key: "grant_type", value: "refresh_token"},
			FormValue{key: "refresh_token", value: auth.Token.RefreshToken},
		)

		auth.parseAndSetToken(err, resp)
	}()
}

// login sends form data to the login page, no need for user interaction this way, callback will be triggered
func (auth *ToonAuthenticator) login(username, password string) {
	resp, err := postFormData(AuthURL,
		FormValue{key: "client_id", value: auth.clientID},
		FormValue{key: "username", value: username},
		FormValue{key: "password", value: password},
		FormValue{key: "redirecturi", value: auth.redirectURI},
		FormValue{key: "tenant_id", value: auth.tenantID},
		FormValue{key: "response_type", value: "code"},
		FormValue{key: "state", value: ""},
		FormValue{key: "scope", value: ""},
	)

	if err != nil || resp.StatusCode != 200 {
		auth.TokenError = fmt.Sprintf("Unable to get OAuth access token, login failed: %v", err)
		auth.Events <- TokenError
	}
}

// request the OAuth service for a new token
func (auth *ToonAuthenticator) getToken(code string) {
	resp, err := postFormData(TokenURL,
		FormValue{key: "client_id", value: auth.clientID},
		FormValue{key: "client_secret", value: auth.clientSecret},
		FormValue{key: "grant_type", value: "authorization_code"},
		FormValue{key: "code", value: code},
	)

	auth.parseAndSetToken(err, resp)
}

func (auth *ToonAuthenticator) parseAndSetToken(err error, resp *http.Response) {
	if err != nil || resp.StatusCode != 200 {
		auth.TokenError = fmt.Sprintf("%v", err)
		auth.Events <- TokenError
	}

	token := OAuthToken{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &token)
	if err != nil {
		auth.TokenError = fmt.Sprintf("Unable to parse received OAuth token: %v", err)
		auth.Events <- TokenError
	}

	auth.IsAuthenticating = false
	auth.Token = token
	auth.Events <- TokenReceived
}

// start a HTTP server with a handler to retrieve the OAuth code after login in
func (auth *ToonAuthenticator) startCallbackServer() {
	host := fmt.Sprintf("%s:%v", auth.callbackHost, auth.callbackPort)
	http.HandleFunc(auth.callbackEndpoint, func(w http.ResponseWriter, r *http.Request) {
		callbackHandler(w, r, auth)
	})

	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// callbackHandler handles the incoming callback after login in, retrieves a code which
// is used to get an access token
func callbackHandler(w http.ResponseWriter, r *http.Request, auth *ToonAuthenticator) {
	code := r.FormValue("code")
	if len(code) == 0 {
		auth.TokenError = "No code found in OAuth2 callback"
		auth.Events <- TokenError
		return
	}

	auth.getToken(code)
}

func postFormData(endpoint string, formvalue ...FormValue) (*http.Response, error) {
	form := url.Values{}
	for _, v := range formvalue {
		form.Add(v.key, v.value)
	}

	hc := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return hc.Do(req)
}
