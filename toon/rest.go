package toon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/tebben/toon-go-sdk/auth"
)

var (
	httpClient = &http.Client{Timeout: 10 * time.Second}
)

func test(url string, auth *auth.ToonAuthenticator) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth.Token.AccessToken))
	resp, _ := httpClient.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

func get(url string, auth *auth.ToonAuthenticator, target interface{}, isRetry bool) *ErrorResponse {
	// if currently authenticating wait 10 seconds and check every 100ms
	if auth.IsAuthenticating {
		for i := 0; i < 100; i++ {
			time.Sleep(100 * time.Millisecond)
			if !auth.IsAuthenticating {
				break
			}
		}
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth.Token.AccessToken))
	resp, err := httpClient.Do(req)
	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&target)
		if err == nil {
			return nil
		}

		return &ErrorResponse{Fault: Fault{Faultstring: fmt.Sprintf("%v", err), Detail: FaultDetail{Errorcode: "Unable to parse JSON"}}}
	}

	// Status not ok
	errorResponse := &ErrorResponse{}
	err = json.NewDecoder(resp.Body).Decode(errorResponse)
	if err != nil {
		if !isRetry && resp.StatusCode == http.StatusUnauthorized && errorResponse.Fault.Faultstring == "Access Token expired" {
			auth.StartRefreshToken()
			return get(url, auth, target, true)
		}

		return errorResponse
	}

	return &ErrorResponse{Fault: Fault{Faultstring: fmt.Sprintf("%v", err), Detail: FaultDetail{Errorcode: fmt.Sprintf("Statuscode: %v", resp.StatusCode)}}}
}

func post() {

}

func put() {

}

func delete() {

}

func constructEndpointURI(endpoint string, params map[string]string, agreementID string) string {
	uri := apiEndpoint
	if len(agreementID) == 0 {
		uri = fmt.Sprintf("%s%s", uri, endpoint)
	} else {
		uri = fmt.Sprintf("%s/%s%s", uri, agreementID, endpoint)
	}

	if params != nil && len(params) > 0 {
		for k, v := range params {
			prefix := "?"
			if strings.Contains(uri, "?") {
				prefix = "&"
			}

			uri = fmt.Sprintf("%s%s%s=%s", uri, prefix, k, v)
		}
	}

	return uri
}
