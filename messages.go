package mailinator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchInboxOptions .
type FetchInboxOptions struct {
	Domain        string `json:"domain"`
	Inbox         string `json:"inbox"`
	Skip          int    `json:"skip"`
	Limit         int    `json:"limit"`
	Sort          Sort   `json:"sort"`
	DecodeSubject bool   `json:"decode_subject"`
}

// Sort .
type Sort string

const (
	ascending  Sort = "ascending"
	descending      = "descending"
)

// Inbox .
type Inbox struct {
	Domain   string    `json:"domain"`
	To       string    `json:"to"`
	Messages []Message `json:"msgs"`
}

// Message .
type Message struct {
	Subject    string  `json:"subject"`
	From       string  `json:"from"`
	To         string  `json:"to"`
	Id         string  `json:"id"`
	Time       float64 `json:"time"`
	SecondsAgo float64 `json:"seconds_ago"`
	Domain     string  `json:"domain"`

	IsFirstExchange bool              `json:"is_first_exchange"`
	Fromfull        string            `json:"fromfull"`
	Headers         map[string]string `json:"headers"`
	Parts           []Part            `json:"parts"`
	Origfrom        string            `json:"origfrom"`
	Mrid            string            `json:"mrid"`
	Size            int               `json:"size"`
	Stream          string            `json:"stream"`
	MsgType         string            `json:"msg_type"`
	Source          string            `json:"source"`
	Text            string            `json:"text"`
}

// Part .
type Part struct {
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

// FetchMessageOptions .
type FetchMessageOptions struct {
	Domain    string `json:"domain"`
	Inbox     string `json:"inbox"`
	MessageId string `json:"message_id"`
}

// FetchSMSMessageOptions .
type FetchSMSMessageOptions struct {
	Domain        string `json:"domain"`
	TeamSMSNumber string `json:"YOUR_TEAM_SMS_NUMBER"`
}

// SMSMessage .
type SMSMessage struct {
	Domain   string    `json:"domain"`
	To       string    `json:"to"`
	Messages []Message `json:"msgs"`
}

// FetchAttachmentsOptions .
type FetchAttachmentsOptions struct {
	Domain    string `json:"domain"`
	Inbox     string `json:"inbox"`
	MessageId string `json:"message_id"`
}

// Attachments .
type Attachments struct {
	Attachments []Attachment `json:"attachments"`
}

// Attachment .
type Attachment struct {
	Filename                string `json:"filename"`
	ContentDisposition      string `json:"content-disposition"`
	ContentTransferEncoding string `json:"content-transfer-encoding"`
	ContentType             string `json:"content-type"`
	AttachmentId            int    `json:"attachment-id"`
}

// FetchAttachmentOptions .
type FetchAttachmentOptions struct {
	Domain       string `json:"domain"`
	Inbox        string `json:"inbox"`
	MessageId    string `json:"message_id"`
	AttachmentId int    `json:"attachment_id"`
}

// FetchAttachmentResponse .
type FetchAttachmentResponse struct {
	Bytes       []byte `json:"bytes"`
	ContentType string `json:"content-type"`
	FileName    string `json:"filename"`
}

// FetchMessageLinksOptions .
type FetchMessageLinksOptions struct {
	Domain    string `json:"domain"`
	Inbox     string `json:"inbox"`
	MessageId string `json:"message_id"`
}

// MessageLinks .
type MessageLinks struct {
	Links []string `json:"links"`
}

// DeleteAllDomainMessagesOptions .
type DeleteAllDomainMessagesOptions struct {
	Domain string `json:"domain"`
}

// DeletedMessages .
type DeletedMessages struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

// DeleteAllInboxMessagesOptions .
type DeleteAllInboxMessagesOptions struct {
	Domain string `json:"domain"`
	Inbox  string `json:"inbox"`
}

// DeleteMessageOptions .
type DeleteMessageOptions struct {
	Domain    string `json:"domain"`
	Inbox     string `json:"inbox"`
	MessageId string `json:"message_id"`
}

// InjectMessageOptions .
type InjectMessageOptions struct {
	Domain  string        `json:"domain"`
	Inbox   string        `json:"inbox"`
	Message MessageToPost `json:"message_to_post"`
}

// MessageToPost .
type MessageToPost struct {
	Subject string `json:"subject"`
	From    string `json:"from"`
	Text    string `json:"text"`
}

// InjectedMessage .
type InjectedMessage struct {
	Status     string `json:"status"`
	Id         string `json:"id"`
	RulesFired []Rule `json:"rules_fired"`
}

// Retrieves a list of messages summaries. You can retreive a list by inbox, inboxes, or entire domain.
func (c *Client) FetchInbox(options *FetchInboxOptions) (*Inbox, error) {
	skip := 0
	limit := 50
	sort := Sort("ascending")
	decodeSubject := false

	if options != nil {
		if options.Skip != 0 {
			skip = options.Skip
		}

		if options.Limit != 0 {
			limit = options.Limit
		}

		if options.Sort != Sort("ascending") {
			sort = options.Sort
		}

		if options.DecodeSubject != false {
			decodeSubject = options.DecodeSubject
		}
	}

	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/inboxes/%s?skip=%d&limit=%d&sort=%v&decode_subject=%t", c.baseURL, options.Domain, options.Inbox, skip, limit, sort, decodeSubject), &buf)
	if err != nil {
		return nil, err
	}

	res := Inbox{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Retrieves a specific message by id.
func (c *Client) FetchMessage(options *FetchMessageOptions) (*Message, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/inboxes/%s/messages/%s", c.baseURL, options.Domain, options.Inbox, options.MessageId), &buf)
	if err != nil {
		return nil, err
	}

	res := Message{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Retrieves a specific SMS message by sms number.
func (c *Client) FetchSMSMessage(options *FetchSMSMessageOptions) (*SMSMessage, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/inboxes/%s", c.baseURL, options.Domain, options.TeamSMSNumber), &buf)
	if err != nil {
		return nil, err
	}

	res := SMSMessage{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Retrieves a list of attachments for a message. Note attachments are expected to be in Email format.
func (c *Client) FetchAtachments(options *FetchAttachmentsOptions) (*Attachments, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/inboxes/%s/messages/%s/attachments", c.baseURL, options.Domain, options.Inbox, options.MessageId), &buf)
	if err != nil {
		return nil, err
	}

	res := Attachments{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Retrieves a specific attachment.
func (c *Client) FetchAttachment(options *FetchAttachmentOptions) (*FetchAttachmentResponse, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/inboxes/%s/messages/%s/attachments/%d", c.baseURL, options.Domain, options.Inbox, options.MessageId, options.AttachmentId), &buf)
	if err != nil {
		return nil, err
	}

	res := FetchAttachmentResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Retrieves all links found within a given email
func (c *Client) FetchMessageLinks(options *FetchMessageLinksOptions) (*MessageLinks, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/inboxes/%s/messages/%s/links", c.baseURL, options.Domain, options.Inbox, options.MessageId), &buf)
	if err != nil {
		return nil, err
	}

	res := MessageLinks{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Deletes ALL messages from a Private Domain. Caution: This action is irreversible.
func (c *Client) DeleteAllDomainMessages(options *DeleteAllDomainMessagesOptions) (*DeletedMessages, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/domains/%s/inboxes", c.baseURL, options.Domain), &buf)
	if err != nil {
		return nil, err
	}

	res := DeletedMessages{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Deletes ALL messages from a specific private inbox.
func (c *Client) DeleteAllInboxMessages(options *DeleteAllInboxMessagesOptions) (*DeletedMessages, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/domains/%s/inboxes/%s", c.baseURL, options.Domain, options.Inbox), &buf)
	if err != nil {
		return nil, err
	}

	res := DeletedMessages{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Deletes a specific messages
func (c *Client) DeleteMessage(options *DeleteMessageOptions) (*DeletedMessages, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/domains/%s/inboxes/%s/messages/%s", c.baseURL, options.Domain, options.Inbox, options.MessageId), &buf)
	if err != nil {
		return nil, err
	}

	res := DeletedMessages{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Deliver a JSON message into your private domain.
func (c *Client) InjectMessage(options *InjectMessageOptions) (*InjectedMessage, error) {
	jsonReq, _ := json.Marshal(options.Message)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/domains/%s/inboxes/%s/messages", c.baseURL, options.Domain, options.Inbox), bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}

	res := InjectedMessage{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
