//go:build integration
// +build integration

package mailinator

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	ENV_API_TOKEN                  = "MAILINATOR_TEST_API_TOKEN"
	ENV_DOMAIN_PRIVATE             = "MAILINATOR_TEST_DOMAIN_PRIVATE"
	ENV_INBOX                      = "MAILINATOR_TEST_INBOX"
	ENV_PHONE_NUMBER               = "MAILINATOR_TEST_PHONE_NUMBER"
	ENV_MESSAGE_ID_WITH_ATTACHMENT = "MAILINATOR_TEST_MESSAGE_ID_WITH_ATTACHMENT"
	ENV_ATTACHMENT_ID              = 0
	ENV_DELETE_DOMAIN              = "MAILINATOR_TEST_DELETE_DOMAIN"
	ENV_INBOX_ALL                  = "MAILINATOR_TEST_INBOX_ALL"
	ENV_WEBHOOKTOKEN_PRIVATEDOMAIN = "MAILINATOR_TEST_WEBHOOKTOKEN_PRIVATEDOMAIN"
	ENV_WEBHOOKTOKEN_CUSTOMSERVICE = "MAILINATOR_TEST_WEBHOOKTOKEN_CUSTOMSERVICE"
	ENV_AUTH_SECRET                = "MAILINATOR_TEST_AUTH_SECRET"
	ENV_AUTH_ID                    = "MAILINATOR_TEST_AUTH_ID"
	ENV_WEBHOOK_INBOX              = "MAILINATOR_TEST_WEBHOOK_INBOX"
	ENV_WEBHOOK_CUSTOMSERVICE      = "MAILINATOR_TEST_WEBHOOK_CUSTOMSERVICE"
)

// will be set by TestGetDomains
var domain Domain

// will be set by TestGetRules
var rule Rule

// will be set by TestFetchInbox
var message Message

// GenerateRandomName generates a random name for testing purposes.
func GenerateRandomName() string {
	rand.Seed(time.Now().UnixNano())
	// Define the character set from which to generate the random name.
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	nameLength := 8 // Adjust the desired length of the random name here

	// Build the random name character by character.
	var nameBuilder strings.Builder
	for i := 0; i < nameLength; i++ {
		randomIndex := rand.Intn(len(charSet))
		nameBuilder.WriteByte(charSet[randomIndex])
	}
	return nameBuilder.String()
}

// Authenticators tests.
func TestInstantTOTP2FACode(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.InstantTOTP2FACode(&InstantTOTP2FACodeOptions{ENV_AUTH_SECRET})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	if res != nil {
		assert.NotNil(t, res.FutureCodes, "expecting non-nil FutureCodes result")
	}
}

func TestGetAuthenticators(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.GetAuthenticators()
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	if res != nil {
		assert.NotNil(t, res.Passcodes, "expecting non-nil Passcodes result")
	}
}

func TestGetAuthenticatorsById(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.GetAuthenticatorsById(&GetAuthenticatorsByIdOptions{ENV_AUTH_ID})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	if res != nil {
		assert.NotNil(t, res.FutureCodes, "expecting non-nil FutureCodes result")
	}
}

func TestGetAuthenticator(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.GetAuthenticator()
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	if res != nil {
		assert.NotNil(t, res.Passcodes, "expecting non-nil Passcodes result")
	}
}

func TestGetAuthenticatorById(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.GetAuthenticatorById(&GetAuthenticatorsByIdOptions{ENV_AUTH_ID})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	if res != nil {
		assert.NotNil(t, res.FutureCodes, "expecting non-nil FutureCodes result")
	}
}

// Domains tests.
func TestGetDomains(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.GetDomains()
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	if res != nil {
		assert.NotNil(t, res.Domains, "expecting non-nil domains result")
		if len(res.Domains) > 0 {
			assert.NotNil(t, res.Domains[0].Id, "expecting non-nil domain id result")
			domain = res.Domains[0]
		}
	}
}

func TestGetDomain(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.GetDomain(&GetDomainOptions{domain.Name})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.NotNil(t, res.Id, "expecting non-nil domain id result")
}

func TestCreateDomain(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	domainName := "testgo.testinator.com"

	res, err := c.CreateDomain(&CreateDomainOptions{domainName})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestDeleteDomain(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	domainName := "testgo.testinator.com"

	res, err := c.DeleteDomain(&DeleteDomainOptions{domainName})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

// Stats tests.
func TestGetTeamStats(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.GetTeamStats()
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.NotNil(t, res.Stats, "expecting non-nil result")
}

func TestGetTeam(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.GetTeam()
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.NotNil(t, res.Id, "expecting non-nil team id result")
}

// Rules tests.
func TestCreateRule(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)
	randomName := GenerateRandomName() // Generate a random name
	rule := RuleToCreate{
		Name:        randomName,
		Priority:    15,
		Description: "Description",
		Conditions: []Condition{
			Condition{
				Operation: OperationType("PREFIX"),
				ConditionData: ConditionData{
					Field: "to",
					Value: "raul",
				},
			},
		},
		Enabled: true,
		Match:   MatchType("ANY"),
		Actions: []ActionRule{
			ActionRule{
				Action: ActionType("WEBHOOK"),
				ActionData: ActionData{
					Url: "https://google.com",
				},
			},
		},
	}

	res, err := c.CreateRule(&CreateRuleOptions{domain.Name, rule})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.NotNil(t, res.Id, "expecting non-nil rule id result")
}

func TestGetAllRules(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.GetAllRules(&GetAllRulesOptions{domain.Name})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	if res != nil {
		assert.NotNil(t, res.Rules, "expecting non-nil rules result")
		if len(res.Rules) > 0 {
			assert.NotNil(t, res.Rules[0].Id, "expecting non-nil rule id result")
			rule = res.Rules[0]
		}
	}
}

func TestEnableRule(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.EnableRule(&EnableRuleOptions{domain.Name, rule.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}

func TestDisableRule(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.DisableRule(&DisableRuleOptions{domain.Name, rule.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}

func TestGetRule(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.GetRule(&GetRuleOptions{domain.Name, rule.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.NotNil(t, res.Id, "expecting non-nil rule id result")
}

func TestDeleteRule(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.DeleteRule(&DeleteRuleOptions{ENV_DELETE_DOMAIN, rule.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}

//Messages tests.

func TestFetchInbox(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchInbox(&FetchInboxOptions{Domain: domain.Name, Inbox: ENV_INBOX_ALL})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	if res != nil {
		assert.NotNil(t, res.Messages, "expecting non-nil messages result")
		if len(res.Messages) > 0 {
			assert.NotNil(t, res.Messages[0].Id, "expecting non-nil message id result")
			message = res.Messages[0]
		}
	}
}

func TestFetchInboxMessage(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchInboxMessage(&FetchInboxMessageOptions{domain.Name, ENV_INBOX_ALL, message.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestFetchMessage(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchMessage(&FetchMessageOptions{domain.Name, message.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestFetchSMSMessage(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchSMSMessage(&FetchSMSMessageOptions{domain.Name, ENV_PHONE_NUMBER})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestInboxMessageAttachments(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchInboxMessageAtachments(&FetchInboxMessageAttachmentsOptions{domain.Name, ENV_INBOX, ENV_MESSAGE_ID_WITH_ATTACHMENT})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestMessageAttachments(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchMessageAtachments(&FetchMessageAttachmentsOptions{domain.Name, ENV_MESSAGE_ID_WITH_ATTACHMENT})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestInboxMessageFetchAttachment(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchInboxMessageAttachment(&FetchInboxMessageAttachmentOptions{domain.Name, ENV_INBOX, ENV_MESSAGE_ID_WITH_ATTACHMENT, ENV_ATTACHMENT_ID})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	assert.NotNil(t, res.Bytes, "expecting non-nil bytes result")
	assert.NotNil(t, res.ContentType, "expecting non-nil content type result")
	assert.NotNil(t, res.FileName, "expecting non-nil file name result")
}

func TestMessageFetchAttachment(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchMessageAttachment(&FetchMessageAttachmentOptions{domain.Name, ENV_MESSAGE_ID_WITH_ATTACHMENT, ENV_ATTACHMENT_ID})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	assert.NotNil(t, res.Bytes, "expecting non-nil bytes result")
	assert.NotNil(t, res.ContentType, "expecting non-nil content type result")
	assert.NotNil(t, res.FileName, "expecting non-nil file name result")
}

func TestFetchMessageLinks(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchMessageLinks(&FetchMessageLinksOptions{domain.Name, ENV_MESSAGE_ID_WITH_ATTACHMENT})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestFetchInboxMessageLinks(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchInboxMessageLinks(&FetchInboxMessageLinksOptions{domain.Name, ENV_INBOX, ENV_MESSAGE_ID_WITH_ATTACHMENT})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestDeleteMessage(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.DeleteMessage(&DeleteMessageOptions{ENV_DELETE_DOMAIN, ENV_INBOX_ALL, message.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}

func TestDeleteAllInboxMessages(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.DeleteAllInboxMessages(&DeleteAllInboxMessagesOptions{ENV_DELETE_DOMAIN, ENV_INBOX})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}

func TestDeleteAllDomainMessages(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.DeleteAllDomainMessages(&DeleteAllDomainMessagesOptions{ENV_DELETE_DOMAIN})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}

func TestPostMessage(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	message := MessageToPost{
		Subject: "Testing message",
		From:    "test_email@test.com",
		Text:    "Hello World!",
	}
	res, err := c.PostMessage(&PostMessageOptions{domain.Name, ENV_INBOX_ALL, message})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}

func TestFetchInboxMessageSmtpLog(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchInboxMessageSmtpLog(&FetchInboxMessageSmtpLogOptions{domain.Name, ENV_INBOX, message.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestFetchMessageSmtpLog(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchMessageSmtpLog(&FetchMessageSmtpLogOptions{domain.Name, message.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestFetchInboxMessageRaw(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchInboxMessageRaw(&FetchInboxMessageRawOptions{domain.Name, ENV_INBOX, message.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestFetchMessageRaw(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchMessageRaw(&FetchMessageRawOptions{domain.Name, message.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

//Webhooks tests.

// Common webhook for testing
var testWebhook = Webhook{
	From:    "sender@example.com",
	Subject: "Test Subject",
	Text:    "Hello, this is a test message.",
	To:      "recipient@example.com",
}

func TestPublicWebhook(t *testing.T) {
	c := NewMailinatorClient("")

	options := &PublicWebhookOptions{
		Webhook: testWebhook,
	}

	res, err := c.PublicWebhook(options)
	if err != nil {
		t.Errorf("PublicWebhook failed: %v", err)
	}

	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}

func TestPublicInboxWebhook(t *testing.T) {
	c := NewMailinatorClient("")

	options := &PublicInboxWebhookOptions{
		Webhook: testWebhook,
		Inbox:   ENV_WEBHOOK_INBOX,
	}

	res, err := c.PublicInboxWebhook(options)
	if err != nil {
		t.Errorf("PublicInboxWebhook failed: %v", err)
	}

	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}

func TestPublicCustomServiceWebhook(t *testing.T) {
	c := NewMailinatorClient("")

	options := &PublicCustomServiceWebhookOptions{
		Webhook:       testWebhook,
		CustomService: ENV_WEBHOOK_CUSTOMSERVICE,
	}

	err := c.PublicCustomServiceWebhook(options)
	if err != nil {
		t.Errorf("PublicCustomServiceWebhook failed: %v", err)
	}
}

func TestPublicCustomServiceInboxWebhook(t *testing.T) {
	c := NewMailinatorClient("")

	options := &PublicCustomServiceInboxWebhookOptions{
		Webhook:       testWebhook,
		CustomService: ENV_WEBHOOK_CUSTOMSERVICE,
		Inbox:         ENV_WEBHOOK_INBOX,
	}

	err := c.PublicCustomServiceInboxWebhook(options)
	if err != nil {
		t.Errorf("PublicCustomServiceInboxWebhook failed: %v", err)
	}
}

func TestPrivateWebhook(t *testing.T) {
	c := NewMailinatorClient("")

	options := &PrivateWebhookOptions{
		WebhookToken: ENV_WEBHOOKTOKEN_PRIVATEDOMAIN,
		Webhook:      testWebhook,
	}

	res, err := c.PrivateWebhook(options)
	if err != nil {
		t.Errorf("PrivateWebhook failed: %v", err)
	}

	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}

func TestPrivateInboxWebhook(t *testing.T) {
	c := NewMailinatorClient("")

	options := &PrivateInboxWebhookOptions{
		WebhookToken: ENV_WEBHOOKTOKEN_PRIVATEDOMAIN,
		Webhook:      testWebhook,
		Inbox:        ENV_WEBHOOK_INBOX,
	}

	res, err := c.PrivateInboxWebhook(options)
	if err != nil {
		t.Errorf("PrivateInboxWebhook failed: %v", err)
	}

	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}

func TestPrivateCustomServiceWebhook(t *testing.T) {
	c := NewMailinatorClient("")

	options := &PrivateCustomServiceWebhookOptions{
		WebhookToken:  ENV_WEBHOOKTOKEN_PRIVATEDOMAIN,
		Webhook:       testWebhook,
		CustomService: ENV_WEBHOOK_CUSTOMSERVICE,
	}

	err := c.PrivateCustomServiceWebhook(options)
	if err != nil {
		t.Errorf("PrivateCustomServiceWebhook failed: %v", err)
	}
}

func TestPrivateCustomServiceInboxWebhook(t *testing.T) {
	c := NewMailinatorClient("")

	options := &PrivateCustomServiceInboxWebhookOptions{
		WebhookToken:  ENV_WEBHOOKTOKEN_PRIVATEDOMAIN,
		Webhook:       testWebhook,
		CustomService: ENV_WEBHOOK_CUSTOMSERVICE,
		Inbox:         ENV_WEBHOOK_INBOX,
	}

	err := c.PrivateCustomServiceInboxWebhook(options)
	if err != nil {
		t.Errorf("PrivateCustomServiceInboxWebhook failed: %v", err)
	}
}
