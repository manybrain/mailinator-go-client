package mailinator

import (
	"fmt"
	"net/http"
)

// InstantTOTP2FACodeOptions .
type InstantTOTP2FACodeOptions struct {
	TotpSecretKey string `json:"totpSecretKey"`
}

// InstantTOTP2FACode .
type InstantTOTP2FACode struct {
	TimeStep         int      `json:"time_step"`
	FutureCodes      []string `json:"futurecodes"`
	NextResetSeconds int      `json:"next_reset_secs"`
	Passcode         string   `json:"passcode"`
}

// Authenticators .
type Authenticators struct {
	Passcodes []Authenticator `json:"passcodes"`
}

// Authenticator .
type Authenticator struct {
	Id               string   `json:"id"`
	TimeStep         int      `json:"time_step"`
	FutureCodes      []string `json:"futurecodes"`
	NextResetSeconds int      `json:"next_reset_secs"`
	Passcode         string   `json:"passcode"`
}

// GetAuthenticatorsByIdOptions .
type GetAuthenticatorsByIdOptions struct {
	Id string `json:"id"`
}

// Instant TOTP 2FA code.
func (c *Client) InstantTOTP2FACode(options *InstantTOTP2FACodeOptions) (*InstantTOTP2FACode, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/totp/%s", c.baseURL, options.TotpSecretKey), nil)
	if err != nil {
		return nil, err
	}

	res := InstantTOTP2FACode{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Fetches Authenticators
func (c *Client) GetAuthenticators() (*Authenticators, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/authenticators", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	res := Authenticators{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Fetch the TOTP 2FA code from one of your saved Keys
func (c *Client) GetAuthenticatorsById(options *GetAuthenticatorsByIdOptions) (*Authenticator, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/authenticators/%s", c.baseURL, options.Id), nil)
	if err != nil {
		return nil, err
	}

	res := Authenticator{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Fetches Authenticator
func (c *Client) GetAuthenticator() (*Authenticators, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/authenticator", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	res := Authenticators{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Fetches Authenticator By Id
func (c *Client) GetAuthenticatorById(options *GetAuthenticatorsByIdOptions) (*Authenticator, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/authenticator/%s", c.baseURL, options.Id), nil)
	if err != nil {
		return nil, err
	}

	res := Authenticator{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
