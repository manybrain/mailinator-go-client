// +build integration

package mailinator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ENV_API_TOKEN                  = "MAILINATOR_TEST_API_TOKEN"
	ENV_DOMAIN_PRIVATE             = "MAILINATOR_TEST_DOMAIN_PRIVATE"
	ENV_INBOX                      = "MAILINATOR_TEST_INBOX"
	ENV_PHONE_NUMBER               = "MAILINATOR_TEST_PHONE_NUMBER"
	ENV_MESSAGE_ID_WITH_ATTACHMENT = "MAILINATOR_TEST_MESSAGE_WITH_ATTACHMENT_ID"
	ENV_ATTACHMENT_ID              = 0
	ENV_DELETE_DOMAIN              = "MAILINATOR_TEST_DELETE_DOMAIN"
	ENV_INBOX_ALL                  = "*"
)

// will be set by TestGetDomains
var domain Domain

// will be set by TestGetRules
var rule Rule

// will be set by TestFetchInbox
var message Message

//Domains tests.
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

	res, err := c.GetDomain(&GetDomainOptions{domain.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.NotNil(t, res.Id, "expecting non-nil domain id result")
}


//Stats tests.
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


//Rules tests.
func TestCreateRule(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)
	rule := RuleToCreate{
		Name:        "RuleName",
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

	res, err := c.CreateRule(&CreateRuleOptions{domain.Id, rule})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.NotNil(t, res.Id, "expecting non-nil rule id result")
}

func TestGetAllRules(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.GetAllRules(&GetAllRulesOptions{domain.Id})
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

	res, err := c.EnableRule(&EnableRuleOptions{domain.Id, rule.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}

func TestDisableRule(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.DisableRule(&DisableRuleOptions{domain.Id, rule.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}


func TestGetRule(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.GetRule(&GetRuleOptions{domain.Id, rule.Id})
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

func TestFetchMessage(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchMessage(&FetchMessageOptions{domain.Name, ENV_INBOX_ALL, message.Id})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestFetchSMSMessage(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchSMSMessage(&FetchSMSMessageOptions{domain.Name, ENV_PHONE_NUMBER})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestAttachments(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchAtachments(&FetchAttachmentsOptions{domain.Name, ENV_INBOX, ENV_MESSAGE_ID_WITH_ATTACHMENT})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestFetchAttachment(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchAttachment(&FetchAttachmentOptions{domain.Name, ENV_INBOX, ENV_MESSAGE_ID_WITH_ATTACHMENT, ENV_ATTACHMENT_ID})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	assert.NotNil(t, res.Bytes, "expecting non-nil bytes result")
	assert.NotNil(t, res.ContentType, "expecting non-nil content type result")
	assert.NotNil(t, res.FileName, "expecting non-nil file name result")
}

func TestFetchMessageLinks(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	res, err := c.FetchMessageLinks(&FetchMessageLinksOptions{domain.Name, ENV_INBOX, ENV_MESSAGE_ID_WITH_ATTACHMENT})
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

func TestInjectMessage(t *testing.T) {
	c := NewMailinatorClient(ENV_API_TOKEN)

	message := MessageToPost{
		Subject: "Testing message",
		From:    "test_email@test.com",
		Text:    "Hello World!",
	}
	res, err := c.InjectMessage(&InjectMessageOptions{domain.Name, ENV_INBOX_ALL, message})
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "ok", res.Status, "expecting correct status")
	}
}