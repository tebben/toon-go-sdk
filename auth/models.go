package auth

// AuthenticationEvent send over ToonAuthenticator channel
type AuthenticationEvent int

// Different authentication events send over the ToonAuthenticator channel
const (
	TokenReceived AuthenticationEvent = iota
	TokenRefreshing
	TokenError
)

// OAuthToken response returned from Toon
type OAuthToken struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             string `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
}

// FormValue contains a key and value which will be
// used for posting form data
type FormValue struct {
	key   string
	value string
}
