package mailinator

import (
	"fmt"
	"net/http"
)

// GetDomainOptions .
type GetDomainOptions struct {
	DomainId string `json:"domain_id"`
}

// CreateDomainOptions .
type CreateDomainOptions struct {
	Name string `json:"name"`
}

// DeleteDomainOptions .
type DeleteDomainOptions struct {
	DomainId string `json:"domain_id"`
}

// DomainsList .
type DomainsList struct {
	Domains []Domain `json:"domains"`
}

// Domain .
type Domain struct {
	Id          string `json:"_id"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Name        string `json:"name"`
	Rules       []Rule `json:"rules"`
}

// Rule .
type Rule struct {
	Id          string       `json:"_id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Enabled     bool         `json:"enabled"`
	Match       MatchType    `json:"match_type"`
	Priority    int          `json:"priority"`
	Conditions  []Condition  `json:"conditions"`
	Actions     []ActionRule `json:"actions"`
}

// MatchType .
type MatchType string

const (
	ANY          MatchType = "ANY"
	ALL                    = "ALL"
	ALWAYS_MATCH           = "ALWAYS_MATCH"
)

// Condition .
type Condition struct {
	Operation     OperationType `json:"operation"`
	ConditionData ConditionData `json:"condition_data"`
}

// OperationType .
type OperationType string

const (
	EQUALS OperationType = "EQUALS"
	PREFIX               = "PREFIX"
)

// ConditionData .
type ConditionData struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

// ActionRule .
type ActionRule struct {
	Action     ActionType `json:"action"`
	ActionData ActionData `json:"action_data"`
}

// ActionType .
type ActionType string

const (
	WEBHOOK ActionType = "WEBHOOK"
	DROP               = "DROP"
)

// ActionData .
type ActionData struct {
	Url string `json:"url"`
}

// Fetches a list of all your domains.
func (c *Client) GetDomains() (*DomainsList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	res := DomainsList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Fetches a specific domain
func (c *Client) GetDomain(options *GetDomainOptions) (*Domain, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s", c.baseURL, options.DomainId), nil)
	if err != nil {
		return nil, err
	}

	res := Domain{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// This endpoint creates a private domain attached to your account. Note, the domain must be unique to the system and you must have not reached your maximum number of Private Domains .
func (c *Client) CreateDomain(options *CreateDomainOptions) (*ResponseStatus, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/domains/%s", c.baseURL, options.Name), nil)
	if err != nil {
		return nil, err
	}

	res := ResponseStatus{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// This endpoint deletes a Private Domain .
func (c *Client) DeleteDomain(options *DeleteDomainOptions) (*ResponseStatus, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/domains/%s", c.baseURL, options.DomainId), nil)
	if err != nil {
		return nil, err
	}

	res := ResponseStatus{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
