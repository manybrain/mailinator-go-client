package mailinator

import (
	"fmt"
	"net/http"
)

// TeamStats .
type TeamStats struct {
	Stats []Stat `json:"stats"`
}

// Stat .
type Stat struct {
	Date      string    `json:"date"`
	Retrieved Retrieved `json:"retrieved"`
	Sent      Sent      `json:"sent"`
}

// Retrieved .
type Retrieved struct {
	WebPublic  int `json:"web_public"`
	ApiError   int `json:"api_error"`
	WebPrivate int `json:"web_private"`
	ApiEmail   int `json:"api_email"`
}

// Sent .
type Sent struct {
	SMS   int `json:"sms"`
	Email int `json:"email"`
}

// TeamInfo .
type TeamInfo struct {
	PrivateDomains []PrivateDomain `json:"private_domains"`
	SMSNumbers     []SMSNumber     `json:"sms_number"`
	Members        []Member        `json:"members"`
	PlanData       PlanData        `json:"plan_data"`
	Id             string          `json:"_id"`
	Plan           string          `json:"plan"`
	TeamName       string          `json:"team_name"`
	Token          string          `json:"token"`
	Status         string          `json:"status"`
}

// TeamInfoData .
type TeamInfoData struct {
	ServerTime string   `json:"server_time"`
	Domains    []string `json:"private_domains"`
}

// PrivateDomain .
type PrivateDomain struct {
	PD      string `json:"pd"`
	Enabled bool   `json:"enabled"`
}

// SMSNumber .
type SMSNumber struct {
	Number  string `json:"number"`
	Country string `json:"country"`
	Status  string `json:"status"`
}

// Member .
type Member struct {
	Role  string `json:"role"`
	Id    string `json:"_id"`
	Email string `json:"email"`
}

// PlanData .
type PlanData struct {
	StorageMb              int `json:"storage_mb"`
	NumberOfPrivateDomains int `json:"num_private_domains"`
	EmailReadsPerDay       int `json:"email_reads_per_day"`
	TeamAccounts           int `json:"team_accounts"`
}

// Retrieves stats of team
func (c *Client) GetTeamStats() (*TeamStats, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/team/stats", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	res := TeamStats{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Retrieves team info
func (c *Client) GetTeam() (*TeamInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/team/", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	res := TeamInfo{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Retrieves team info
func (c *Client) GetTeamInfo() (*TeamInfoData, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/teaminfo", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	res := TeamInfoData{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
