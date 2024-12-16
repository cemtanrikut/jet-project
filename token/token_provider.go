package token

// TokenProvider defines the interface for generating access tokens.
type TokenProvider interface {
	GetAccessToken() (string, error)
}
