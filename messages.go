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
	Cursor        string `json:"cursor"`
	Full          bool   `json:"full"`
	Delete        string `json:"delete"`
	Wait          string `json:"wait"`
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
	Cursor   string    `json:"cursor"`
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

	IsFirstExchange bool                   `json:"is_first_exchange"`
	Fromfull        string                 `json:"fromfull"`
	Headers         map[string]interface{} `json:"headers"`
	Parts           []Part                 `json:"parts"`
	Origfrom        string                 `json:"origfrom"`
	Mrid            string                 `json:"mrid"`
	Size            int                    `json:"size"`
	Stream          string                 `json:"stream"`
	MsgType         string                 `json:"msg_type"`
	Source          string                 `json:"source"`
	Text            string                 `json:"text"`
}

// Part .
type Part struct {
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

// FetchInboxMessageOptions .
type FetchInboxMessageOptions struct {
	Domain    string `json:"domain"`
	Inbox     string `json:"inbox"`
	MessageId string `json:"message_id"`
}

// FetchMessageOptions .
type FetchMessageOptions struct {
	Domain    string `json:"domain"`
	MessageId string `json:"message_id"`
	Delete    string `json:"delete"`
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

// FetchInboxMessageAttachmentsOptions .
type FetchInboxMessageAttachmentsOptions struct {
	Domain    string `json:"domain"`
	Inbox     string `json:"inbox"`
	MessageId string `json:"message_id"`
}

// FetchMessageAttachmentsOptions .
type FetchMessageAttachmentsOptions struct {
	Domain    string `json:"domain"`
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

// FetchInboxMessageAttachmentOptions .
type FetchInboxMessageAttachmentOptions struct {
	Domain       string `json:"domain"`
	Inbox        string `json:"inbox"`
	MessageId    string `json:"message_id"`
	AttachmentId int    `json:"attachment_id"`
}

// FetchMessageAttachmentOptions .
type FetchMessageAttachmentOptions struct {
	Domain       string `json:"domain"`
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
	MessageId string `json:"message_id"`
}

// FetchMessageLinksFullOptions .
type FetchMessageLinksFullOptions struct {
	Domain    string `json:"domain"`
	MessageId string `json:"message_id"`
}

// FetchInboxMessageLinksOptions .
type FetchInboxMessageLinksOptions struct {
	Domain    string `json:"domain"`
	Inbox     string `json:"inbox"`
	MessageId string `json:"message_id"`
}

// MessageLinks .
type MessageLinks struct {
	Links []string `json:"links"`
}

// MessageLinksFull .
type MessageLinksFull struct {
	Links []LinkEntity `json:"links"`
}

// LinkEntity .
type LinkEntity struct {
	Link string `json:"link"`
	Text string `json:"text"`
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

// PostMessageOptions .
type PostMessageOptions struct {
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

// PostedMessage .
type PostedMessage struct {
	Status     string `json:"status"`
	Id         string `json:"id"`
	RulesFired []Rule `json:"rules_fired"`
}

// FetchMessageSmtpLogOptions .
type FetchMessageSmtpLogOptions struct {
	Domain    string `json:"domain"`
	MessageId string `json:"message_id"`
}

// FetchInboxMessageSmtpLogOptions .
type FetchInboxMessageSmtpLogOptions struct {
	Domain    string `json:"domain"`
	Inbox     string `json:"inbox"`
	MessageId string `json:"message_id"`
}

// MessageSmtpLogs .
type MessageSmtpLogs struct {
	LogEntries []EmailLogEntry `json:"log"`
}

// EmailLogEntry .
type EmailLogEntry struct {
	Log   string `json:"log"`
	Time  string `json:"time"`
	Event string `json:"event"`
}

// FetchMessageRawOptions .
type FetchMessageRawOptions struct {
	Domain    string `json:"domain"`
	MessageId string `json:"message_id"`
}

// FetchInboxMessageRawOptions .
type FetchInboxMessageRawOptions struct {
	Domain    string `json:"domain"`
	Inbox     string `json:"inbox"`
	MessageId string `json:"message_id"`
}

// MessageRaw .
type MessageRaw struct {
	RawData string `json:"rawData"`
}

// FetchLatestMessagesOptions .
type FetchLatestMessagesOptions struct {
	Domain string `json:"domain"`
}

// FetchLatestInboxMessagesOptions .
type FetchLatestInboxMessagesOptions struct {
	Domain string `json:"domain"`
	Inbox  string `json:"inbox"`
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

	url := fmt.Sprintf("%s/domains/%s/inboxes/%s?skip=%d&limit=%d&sort=%v&decode_subject=%t", c.baseURL, options.Domain, options.Inbox, skip, limit, sort, decodeSubject)

	if options.Cursor != "" {
		url = fmt.Sprintf("%s&cursor=%s", url, options.Cursor)
	}

	if options.Full != false {
		url = fmt.Sprintf("%s&full=%s", url, "true")
	}

	if options.Delete != "" {
		url = fmt.Sprintf("%s&delete=%s", url, options.Delete)
	}

	if options.Wait != "" {
		url = fmt.Sprintf("%s&wait=%s", url, options.Wait)
	}

	req, err := http.NewRequest("GET", url, &buf)
	if err != nil {
		return nil, err
	}

	res := Inbox{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Retrieves a specific message by id for specific inbox.
func (c *Client) FetchInboxMessage(options *FetchInboxMessageOptions) (*Message, error) {
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

// Retrieves a specific message by id.
func (c *Client) FetchMessage(options *FetchMessageOptions) (*Message, error) {
	var buf bytes.Buffer

	url := fmt.Sprintf("%s/domains/%s/messages/%s", c.baseURL, options.Domain, options.MessageId)

	if options.Delete != "" {
		url = fmt.Sprintf("%s?delete=%s", url, options.Delete)
	}

	req, err := http.NewRequest("GET", url, &buf)
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

// Retrieves a list of attachments for a message for specific inbox. Note attachments are expected to be in Email format.
func (c *Client) FetchInboxMessageAtachments(options *FetchInboxMessageAttachmentsOptions) (*Attachments, error) {
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

// Retrieves a list of attachments for a message. Note attachments are expected to be in Email format.
func (c *Client) FetchMessageAtachments(options *FetchMessageAttachmentsOptions) (*Attachments, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/messages/%s/attachments", c.baseURL, options.Domain, options.MessageId), &buf)
	if err != nil {
		return nil, err
	}

	res := Attachments{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Retrieves a specific attachment for specific inbox .
func (c *Client) FetchInboxMessageAttachment(options *FetchInboxMessageAttachmentOptions) (*FetchAttachmentResponse, error) {
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

// Retrieves a specific attachment.
func (c *Client) FetchMessageAttachment(options *FetchMessageAttachmentOptions) (*FetchAttachmentResponse, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/messages/%s/attachments/%d", c.baseURL, options.Domain, options.MessageId, options.AttachmentId), &buf)
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

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/messages/%s/links", c.baseURL, options.Domain, options.MessageId), &buf)
	if err != nil {
		return nil, err
	}

	res := MessageLinks{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Retrieves all links full found within a given email
func (c *Client) FetchMessageLinksFull(options *FetchMessageLinksFullOptions) (*MessageLinksFull, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/messages/%s/linksfull", c.baseURL, options.Domain, options.MessageId), &buf)
	if err != nil {
		return nil, err
	}

	res := MessageLinksFull{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Retrieves all links found within a given email for specific inbox .
func (c *Client) FetchInboxMessageLinks(options *FetchInboxMessageLinksOptions) (*MessageLinks, error) {
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
func (c *Client) PostMessage(options *PostMessageOptions) (*PostedMessage, error) {
	jsonReq, _ := json.Marshal(options.Message)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/domains/%s/inboxes/%s/messages", c.baseURL, options.Domain, options.Inbox), bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}

	res := PostedMessage{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// This endpoint retrieves smtp log from the email .
func (c *Client) FetchMessageSmtpLog(options *FetchMessageSmtpLogOptions) (*MessageSmtpLogs, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/messages/%s/smtplog", c.baseURL, options.Domain, options.MessageId), &buf)
	if err != nil {
		return nil, err
	}

	res := MessageSmtpLogs{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// This endpoint retrieves smtp log from the email for specific inbox .
func (c *Client) FetchInboxMessageSmtpLog(options *FetchInboxMessageSmtpLogOptions) (*MessageSmtpLogs, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/inboxes/%s/messages/%s/smtplog", c.baseURL, options.Domain, options.Inbox, options.MessageId), &buf)
	if err != nil {
		return nil, err
	}

	res := MessageSmtpLogs{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// This endpoint retrieves raw info from the email .
func (c *Client) FetchMessageRaw(options *FetchMessageRawOptions) (*string, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/messages/%s/raw", c.baseURL, options.Domain, options.MessageId), &buf)
	if err != nil {
		return nil, err
	}

	res := new(string)
	if err := c.sendRequestWithOptions(req, res, true); err != nil {
		return nil, err
	}

	return res, nil
}

// This endpoint retrieves raw info from the email for specific inbox .
func (c *Client) FetchInboxMessageRaw(options *FetchInboxMessageRawOptions) (*string, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/inboxes/%s/messages/%s/raw", c.baseURL, options.Domain, options.Inbox, options.MessageId), &buf)
	if err != nil {
		return nil, err
	}

	res := new(string)
	if err := c.sendRequestWithOptions(req, res, true); err != nil {
		return nil, err
	}

	return res, nil
}

// That fetches the latest 5 FULL messages .
func (c *Client) FetchLatestMessages(options *FetchLatestMessagesOptions) (*Inbox, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/messages/*", c.baseURL, options.Domain), &buf)
	if err != nil {
		return nil, err
	}

	res := Inbox{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// That fetches the latest 5 FULL messages for specific inbox .
func (c *Client) FetchLatestInboxMessages(options *FetchLatestInboxMessagesOptions) (*Inbox, error) {
	var buf bytes.Buffer

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domains/%s/inboxes/%s/messages/*", c.baseURL, options.Domain, options.Inbox), &buf)
	if err != nil {
		return nil, err
	}

	res := Inbox{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
