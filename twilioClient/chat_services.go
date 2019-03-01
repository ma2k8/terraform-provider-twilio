package twilioClient

import (
	"golang.org/x/net/context"
	"net/url"
	"strings"
)

/**
 * Doc: https://www.twilio.com/docs/chat/rest/services
 */
const chatServiceApiPath string = "Services"

type ChatService struct {
	Sid                          string             `json:"sid"`
	AccountSid                   string             `json:"account_sid"`
	ConsumptionReportInterval    string             `json:"consumption_report_interval"`
	DateCreated                  string             `json:"date_created"`
	DateUpdated                  string             `json:"date_updated"`
	DefaultChannelCreatorRoleSid string             `json:"default_channel_creator_role_sid"`
	DefaultServiceRoleSid        string             `json:"default_service_role_sid"`
	FriendlyName                 string             `json:"friendly_name"`
	Limits                       ChatServiceLimit   `json:"limits"`
	Links                        map[string]string  `json:"links"`
	PostWebhookUrl               string             `json:"post_webhook_url"`
	PreWebhookUrl                string             `json:"pre_webhook_url"`
	PostWebhookRetryCount        int                `json:"post_webhook_retry_count"`
	PreWebhookRetryCount         int                `json:"pre_webhook_retry_count"`
	ReachabilityEnabled          bool               `json:"reachability_enabled"`
	ReadStatusEnabled            bool               `json:"read_status_enabled"`
	TypingIndicatorTimeout       int                `json:"typing_indicator_timeout"`
	Url                          string             `json:"url"`
	WebhookFilters               []ChatWebhookEvent `json:"webhook_filters"`
	WebhookMethod                string             `json:"webhook_method"`
	Media                        ChatServiceMedia   `json:"media"`
	//Notifications
}

type ChatServiceLimit struct {
	ChannelMembers string `json:"channel_members"`
	UserChannels   string `json:"user_channels"`
}

type ChatServiceMedia struct {
	SizeLimitMb          string `json:"size_limit_mb"`
	CompatibilityMessage string `json:"compatibility_message"`
}

type ChatWebhookEvent string

const OnMessageSend = ChatWebhookEvent("onMessageSend")
const OnMessageRemove = ChatWebhookEvent("onMessageRemove")
const OnMessageUpdate = ChatWebhookEvent("onMessageUpdate")
const OnMediaMessageSend = ChatWebhookEvent("onMediaMessageSend")
const OnChannelAdd = ChatWebhookEvent("onChannelAdd")
const OnChannelUpdate = ChatWebhookEvent("onChannelUpdate")
const OnChannelDestroy = ChatWebhookEvent("onChannelDestroy")
const OnMemberAdd = ChatWebhookEvent("onMemberAdd")
const OnMemberRemove = ChatWebhookEvent("onMemberRemove")
const OnUserAdded = ChatWebhookEvent("onUserAdded")
const OnUserUpdate = ChatWebhookEvent("onUserUpdate")

func (s ChatWebhookEvent) ToString() string {
	switch s {
	case OnMessageSend:
		return "onMessageSend"
	case OnMessageRemove:
		return "onMessageRemove"
	case OnMessageUpdate:
		return "onMessageUpdate"
	case OnMediaMessageSend:
		return "onMediaMessageSend"
	case OnChannelAdd:
		return "onChannelAdd"
	case OnChannelUpdate:
		return "onChannelUpdate"
	case OnChannelDestroy:
		return "onChannelDestroy"
	case OnMemberAdd:
		return "onMemberAdd"
	case OnMemberRemove:
		return "onMemberRemove"
	case OnUserAdded:
		return "onUserAdded"
	case OnUserUpdate:
		return "onUserUpdate"
	default:
		return strings.Title(string(s))
	}
}

type ChatServicePage struct {
	Page
	ChatServices []*ChatService `json:"services"`
}

type ChatServiceService struct {
	client *TwilioClient
}

func (service *ChatServiceService) Create(ctx context.Context, data url.Values) (*ChatService, error) {
	var targetUrl = service.client.ChatBaseUrl() + "/" + chatServiceApiPath
	app := new(ChatService)
	err := service.client.post(targetUrl, data, app)
	return app, err
}

func (service *ChatServiceService) Update(ctx context.Context, appSid string, data url.Values) (*ChatService, error) {
	var targetUrl = service.client.ChatBaseUrl() + "/" + chatServiceApiPath + "/" + appSid

	app := new(ChatService)
	err := service.client.post(targetUrl, data, app)
	return app, err
}

func (service *ChatServiceService) Delete(ctx context.Context, appSid string) error {
	var targetUrl = service.client.ChatBaseUrl() + "/" + chatServiceApiPath + "/" + appSid

	return service.client.delete(targetUrl)
}

func (service *ChatServiceService) Get(ctx context.Context, appSid string) (*ChatService, error) {
	var targetUrl = service.client.ChatBaseUrl() + "/" + chatServiceApiPath + "/" + appSid

	app := new(ChatService)
	err := service.client.get(targetUrl, url.Values{}, app)

	return app, err
}

func (service *ChatServiceService) GetPage(ctx context.Context, data url.Values) (*ChatServicePage, error) {
	iter := service.GetPageIterator(data)
	return iter.Next(ctx)
}

type ChatServicePageIterator struct {
	p *PageIterator
}

func (service *ChatServiceService) GetPageIterator(data url.Values) *ChatServicePageIterator {
	iter := NewPageIterator(service.client, data, service.client.ChatBaseUrl()+"/"+chatServiceApiPath)
	return &ChatServicePageIterator{
		p: iter,
	}
}

func (c *ChatServicePageIterator) Next(ctx context.Context) (*ChatServicePage, error) {
	cp := new(ChatServicePage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
