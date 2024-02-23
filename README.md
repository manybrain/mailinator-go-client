# mailinator-go-client

[Mailianator.com](https://www.mailinator.com/) Go client.

## Install

```
go get -u github.com/manybrain/mailinator-go-client
```

## Usage

To start using the API you need to first create an account at [mailinator.com](https://www.mailinator.com/).

Once you have an account you will need an API Token which you can generate in [mailinator.com/v3/#/#team_settings_pane](https://www.mailinator.com/v3/#/#team_settings_pane).

Then you can configure the library with:

```go
import "github.com/manybrain/mailinator-go-client"

// Replace API_TOKEN with your real key
client := mailinator.NewMailinatorClient("API_TOKEN")
```

## Examples

##### Domains methods:

- Get AllDomains / Domain:

  ```go
    import "github.com/manybrain/mailinator-go-client"

	client := mailinator.NewMailinatorClient("yourApiTokenHere")

    //Get All Domains
	res, err := client.GetDomains()
	
	//Get Domain
	res, err := client.GetDomain(&GetDomainOptions{"yourDomainIdHere"})
    // ...
  ```

##### Rules methods:

- Create / Delete Rule:

  ```go
    import "github.com/manybrain/mailinator-go-client"

	client := mailinator.NewMailinatorClient("yourApiTokenHere")

    //Create Rule
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

	res, err := client.CreateRule(&CreateRuleOptions{"yourDomainIdHere", rule})
            
    //Delete Rule
    res, err := client.DeleteRule(&DeleteRuleOptions{"yourDomainIdHere", "yourRuleIdHere"})
    // ...
  ```

- Enable/Disable Rule:

  ```go
    import "github.com/manybrain/mailinator-go-client"

	client := mailinator.NewMailinatorClient("yourApiTokenHere")

    //Enable Rule
    res, err := client.EnableRule(&EnableRuleOptions{"yourDomainIdHere", "yourRuleIdHere"})
    
    //Disable Rule
    res, err := client.DisableRule(&DisableRuleOptions{"yourDomainIdHere", "yourRuleIdHere"})
  ```

- Get All Rules / Rule:

```go
    import "github.com/manybrain/mailinator-go-client"

	client := mailinator.NewMailinatorClient("yourApiTokenHere")

    //Get All Rules
    res, err := client.GetAllRules(&GetAllRulesOptions{"yourDomainIdHere"})
    
    //Get Rule
    res, err := client.GetRule(&GetRuleOptions{"yourDomainIdHere", "yourRuleIdHere"})
```

##### Messages methods:

- Inject Message:

  ```go
    import "github.com/manybrain/mailinator-go-client"

	client := mailinator.NewMailinatorClient("yourApiTokenHere")

	message := MessageToPost{
			Subject: "Testing message",
			From:    "test_email@test.com",
			Text:    "Hello World!",
		}
		res, err := client.PostMessage(&PostMessageOptions{"yourDomainNameHere", "yourInboxHere", message})
    // ...
  ```

- Fetch Inbox / Message / SMS Messages / Attachments / Attachment:

  ```go
    import "github.com/manybrain/mailinator-go-client"

	client := mailinator.NewMailinatorClient("yourApiTokenHere")

    //Fetch Inbox
    res, err := client.FetchInbox(&FetchInboxOptions{Domain: "yourDomainNameHere", Inbox: "yourInboxHere"})
    
    //Fetch Message
	res, err := client.FetchInboxMessage(&FetchInboxMessageOptions{"yourDomainNameHere", "yourInboxHere", "yourMessageIdHere"})
    
    //Fetch SMS Messages
	res, err := client.FetchSMSMessage(&FetchSMSMessageOptions{"yourDomainNameHere", "yourTeamSMSNumberHere"})
    
    //Fetch Attachments
	res, err := client.FetchInboxMessageAtachments(&FetchInboxMessageAttachmentsOptions{"yourDomainNameHere", "yourInboxHere", "yourMessageIdWithAttachmentHere"})
    
    //Fetch Attachment
	res, err := client.FetchInboxMessageAttachment(&FetchInboxMessageAttachmentOptions{"yourDomainNameHere", "yourInboxHere", "yourMessageIdWithAttachmentHere", "yourAttachmentIdHere"})
            
    //Fetch Message Links
	res, err := client.FetchInboxMessageLinks(&FetchInboxMessageLinksOptions{"yourDomainNameHere", "yourInboxHere", "yourMessageIdHere"})
  ```

- Delete Message / AllInboxMessages / AllDomainMessages

  ```go
    import "github.com/manybrain/mailinator-go-client"

	client := mailinator.NewMailinatorClient("yourApiTokenHere")

    //Delete Message
	res, err := client.DeleteMessage(&DeleteMessageOptions{"yourDomainNameHere", "yourInboxHere", "yourMessageIdHere"})
    
    //Delete All Inbox Messages
	res, err := client.DeleteAllInboxMessages(&DeleteAllInboxMessagesOptions{"yourDomainNameHere", "yourInboxHere"})
    
    //Delete All Domain Messages
	res, err := client.DeleteAllDomainMessages(&DeleteAllDomainMessagesOptions{"yourDomainNameHere"})
  ```
  
##### Stats methods:

- Get Team / Team Stats:

  ```go
    import "github.com/manybrain/mailinator-go-client"

	client := mailinator.NewMailinatorClient("yourApiTokenHere")

    //Get Team
	res, err := client.GetTeamStats()
	
    //Get TeamStats
    res, err := client.GetTeam()
    // ...
  ```

## Testing

Run integration tests with real API Key.

```
go test -v -tags=integration
```

Most of the tests require env variables with valid values. Visit tests source code and review `integration_test.go` file. The more env variables you set, the more tests are run.

* `MAILINATOR_TEST_API_TOKEN` - API tokens for authentication; basic requirement across many tests;see also https://manybrain.github.io/m8rdocs/#api-authentication
* `MAILINATOR_TEST_DOMAIN_PRIVATE` - private domain; visit https://www.mailinator.com/
* `MAILINATOR_TEST_INBOX` - some already existing inbox within the private domain
* `MAILINATOR_TEST_PHONE_NUMBER` - associated phone number within the private domain; see also https://manybrain.github.io/m8rdocs/#fetch-an-sms-messages
* `MAILINATOR_TEST_MESSAGE_WITH_ATTACHMENT_ID` - existing message id within inbox (see above) within private domain (see above); see also https://manybrain.github.io/m8rdocs/#fetch-message
* `MAILINATOR_TEST_ATTACHMENT_ID` - existing message id within inbox (see above) within private domain (see above); see also https://manybrain.github.io/m8rdocs/#fetch-message
* `MAILINATOR_TEST_DELETE_DOMAIN` - don't use it unless you are 100% sure what you are doing
* `MAILINATOR_TEST_WEBHOOKTOKEN_PRIVATEDOMAIN` - private domain for webhook token
* `MAILINATOR_TEST_WEBHOOKTOKEN_CUSTOMSERVICE` - custom service for webhook token
* `MAILINATOR_TEST_AUTH_SECRET` - authenticator secret
* `MAILINATOR_TEST_AUTH_ID` - authenticator id
* `MAILINATOR_TEST_WEBHOOK_INBOX` - inbox for webhook
* `MAILINATOR_TEST_WEBHOOK_CUSTOMSERVICE` - custom service for webhook

