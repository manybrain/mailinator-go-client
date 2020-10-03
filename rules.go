package mailinator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateRuleOptions .
type CreateRuleOptions struct {
	DomainId     string       `json:"domain_id"`
	RuleToCreate RuleToCreate `json:"rule_to_create"`
}

// RuleToCreate .
type RuleToCreate struct {
	Description string       `json:"description"`
	Enabled     bool         `json:"enabled"`
	Match       MatchType    `json:"match"`
	Name        string       `json:"name"`
	Priority    int          `json:"priority"`
	Conditions  []Condition  `json:"conditions"`
	Actions     []ActionRule `json:"actions"`
}

// EnableRuleOptions .
type EnableRuleOptions struct {
	DomainId string `json:"domain_id"`
	RuleId   string `json:"rule_id"`
}

// DisableRuleOptions .
type DisableRuleOptions struct {
	DomainId string `json:"domain_id"`
	RuleId   string `json:"rule_id"`
}

// ResponseStatus .
type ResponseStatus struct {
	Status string `json:"status"`
}

// GetAllRulesOptions .
type GetAllRulesOptions struct {
	DomainId string `json:"domain_id"`
}

// GetRuleOptions .
type GetRuleOptions struct {
	DomainId string `json:"domain_id"`
	RuleId   string `json:"rule_id"`
}

// DeleteRuleOptions .
type DeleteRuleOptions struct {
	DomainId string `json:"domain_id"`
	RuleId   string `json:"rule_id"`
}

// Rules .
type Rules struct {
	Rules []Rule `json:"rules"`
}

// Creates a Rule. Note that in the examples, ":domain_id" can be one of your private domains.
func (c *Client) CreateRule(options *CreateRuleOptions) (*Rule, error) {
	jsonReq, _ := json.Marshal(options.RuleToCreate)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/domains/%s/rules", c.baseURL, options.DomainId), bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}

	res := Rule{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Enable an existing Rule
func (c *Client) EnableRule(options *EnableRuleOptions) (*ResponseStatus, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/domains/%s/rules/%s/enable", c.baseURL, options.DomainId, options.RuleId), &buf)
	if err != nil {
		return nil, err
	}

	res := ResponseStatus{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Disable an existing Rule
func (c *Client) DisableRule(options *DisableRuleOptions) (*ResponseStatus, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/domains/%s/rules/%s/disable", c.baseURL, options.DomainId, options.RuleId), &buf)
	if err != nil {
		return nil, err
	}

	res := ResponseStatus{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Fetches a All Rules for a Domain
func (c *Client) GetAllRules(options *GetAllRulesOptions) (*Rules, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/rules", c.baseURL, options.DomainId), nil)
	if err != nil {
		return nil, err
	}

	res := Rules{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Fetches a Rules for a Domain
func (c *Client) GetRule(options *GetRuleOptions) (*Rule, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/rules/%s", c.baseURL, options.DomainId, options.RuleId), nil)
	if err != nil {
		return nil, err
	}

	res := Rule{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Deletes a specific Rule from a Domain
func (c *Client) DeleteRule(options *DeleteRuleOptions) (*ResponseStatus, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/domains/%s/rules/%s", c.baseURL, options.DomainId, options.RuleId), nil)
	if err != nil {
		return nil, err
	}

	res := ResponseStatus{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
