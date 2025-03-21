package mailinator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// PrivateWebhookOptions .
type PrivateWebhookOptions struct {
	WebhookToken string  `json:"wh-token"`
	Webhook      Webhook `json:"webhook"`
}

// PrivateInboxWebhookOptions .
type PrivateInboxWebhookOptions struct {
	WebhookToken string  `json:"wh-token"`
	Webhook      Webhook `json:"webhook"`
	Inbox        string  `json:"inbox"`
}

// PrivateCustomServiceWebhookOptions .
type PrivateCustomServiceWebhookOptions struct {
	WebhookToken  string  `json:"wh-token"`
	Webhook       Webhook `json:"webhook"`
	CustomService string  `json:"customService"`
}

// PrivateCustomServiceInboxWebhookOptions .
type PrivateCustomServiceInboxWebhookOptions struct {
	WebhookToken  string  `json:"wh-token"`
	Webhook       Webhook `json:"webhook"`
	Inbox         string  `json:"inbox"`
	CustomService string  `json:"customService"`
}

// Webhook .
type Webhook struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
	To      string `json:"to"`
}

// ResponseStatusWithId .
type ResponseStatusWithId struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

// This command will Webhook messages into your Private Domain .
// The incoming Webhook will arrive in the inbox designated by the "to" field in the incoming request payload .
// Webhooks into your Private System do NOT use your regular API Token .
// This is because a typical use case is to enter the Webhook URL into 3rd-party systems(i.e.Twilio, Zapier, IFTTT, etc) and you should never give out your API Token .
// Check your Team Settings where you can create "Webhook Tokens" designed for this purpose .
func (c *Client) PrivateWebhook(options *PrivateWebhookOptions) (*ResponseStatusWithId, error) {
	jsonReq, err := json.Marshal(options.Webhook)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/domains/private/webhook?whtoken=%s", c.baseURL, options.WebhookToken), bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}

	res := ResponseStatusWithId{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// This command will deliver the message to the :inbox inbox .
// Incoming Webhooks are delivered to Mailinator inboxes and from that point onward are not notably different than other messages in the system (i.e. emails) .
// As normal, Mailinator will list all messages in the Inbox page and via the Inbox API calls .
// If the incoming JSON payload does not contain a "from" or "subject", then dummy values will be inserted in these fields .
// You may retrieve such messages via the Web Interface, the API, or the Rule System .
func (c *Client) PrivateInboxWebhook(options *PrivateInboxWebhookOptions) (*ResponseStatusWithId, error) {
	jsonReq, err := json.Marshal(options.Webhook)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/domains/private/webhook/%s?whtoken=%s", c.baseURL, options.Inbox, options.WebhookToken), bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}

	res := ResponseStatusWithId{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// If you have a Twilio account which receives incoming SMS messages. You may direct those messages through this facility to inject those messages into the Mailinator system .
// Mailinator intends to apply specific mappings for certain services that commonly publish webhooks .
// If you test incoming Messages to SMS numbers via Twilio, you may use this endpoint to correctly map "to", "from", and "subject" of those messages to the Mailinator system.By default, the destination inbox is the Twilio phone number .
func (c *Client) PrivateCustomServiceWebhook(options *PrivateCustomServiceWebhookOptions) error {
	jsonReq, err := json.Marshal(options.Webhook)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/domains/private/%s?whtoken=%s", c.baseURL, options.CustomService, options.WebhookToken), bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}

	res := new(string)

	if err := c.sendRequestWithOptions(req, res, true); err != nil {
		return err
	}

	return nil
}

// The SMS message will arrive in the Private Mailinator inbox corresponding to the Twilio Phone Number. (only the digits, if a plus sign precedes the number it will be removed)
// If you wish the message to arrive in a different inbox, you may append the destination inbox to the URL .
func (c *Client) PrivateCustomServiceInboxWebhook(options *PrivateCustomServiceInboxWebhookOptions) error {
	jsonReq, err := json.Marshal(options.Webhook)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/domains/private/%s/%s?whtoken=%s", c.baseURL, options.CustomService, options.Inbox, options.WebhookToken), bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}

	res := new(string)

	if err := c.sendRequestWithOptions(req, res, true); err != nil {
		return err
	}

	return nil
}
