package twilioClient

import (
	"golang.org/x/net/context"
	"net/url"
)

/**
 * Doc: https://www.twilio.com/docs/notify/api/service-resource
 */
const notifyServiceApiPath string = "Services"

type NotifyService struct {
	Sid                                     string            `json:"sid"`
	AccountSid                              string            `json:"account_sid"`
	DateCreated                             string            `json:"date_created"`
	DateUpdated                             string            `json:"date_updated"`
	FriendlyName                            string            `json:"friendly_name"`
	ApnCredentialSid                        string            `json:"apn_credential_sid"`
	GcmCredentialSid                        string            `json:"gcm_credential_sid"`
	FcmCredentialSid                        string            `json:"fcm_credential_sid"`
	MessagingServiceSid                     string            `json:"messaging_service_sid"`
	FacebookMessengerPageId                 string            `json:"facebook_messenger_page_id"`
	AlexaSkillId                            string            `json:"alexa_skill_id"`
	DefaultApnNotificationProtocolVersion   string            `json:"default_apn_notification_protocol_version"`
	DefaultGcmNotificationProtocolVersion   string            `json:"default_gcm_notification_protocol_version"`
	DefaultFcmNotificationProtocolVersion   string            `json:"default_fcm_notification_protocol_version"`
	DefaultAlexaNotificationProtocolVersion string            `json:"default_alexa_notification_protocol_version"`
	LogEnabled                              bool              `json:"log_enabled"`
	Type                                    string            `json:"type"`
	Url                                     string            `json:"url"`
	Links                                   map[string]string `json:"links"`
}

type NotifyServiceType string

type NotifyServicePage struct {
	Page
	NotifyServices []*NotifyService `json:"services"`
}

type NotifyServiceService struct {
	client *TwilioClient
}

func (service *NotifyServiceService) Create(ctx context.Context, data url.Values) (*NotifyService, error) {
	var targetUrl = service.client.NotifyBaseUrl() + "/" + notifyServiceApiPath
	app := new(NotifyService)
	err := service.client.post(targetUrl, data, app)
	return app, err
}

func (service *NotifyServiceService) Update(ctx context.Context, appSid string, data url.Values) (*NotifyService, error) {
	var targetUrl = service.client.NotifyBaseUrl() + "/" + notifyServiceApiPath + "/" + appSid

	app := new(NotifyService)
	err := service.client.post(targetUrl, data, app)
	return app, err
}

func (service *NotifyServiceService) Delete(ctx context.Context, appSid string) error {
	var targetUrl = service.client.NotifyBaseUrl() + "/" + notifyServiceApiPath + "/" + appSid

	return service.client.delete(targetUrl)
}

func (service *NotifyServiceService) Get(ctx context.Context, appSid string) (*NotifyService, error) {
	var targetUrl = service.client.NotifyBaseUrl() + "/" + notifyServiceApiPath + "/" + appSid

	app := new(NotifyService)
	err := service.client.get(targetUrl, url.Values{}, app)

	return app, err
}

func (service *NotifyServiceService) GetPage(ctx context.Context, data url.Values) (*NotifyServicePage, error) {
	iter := service.GetPageIterator(data)
	return iter.Next(ctx)
}

type NotifyServicePageIterator struct {
	p *PageIterator
}

func (service *NotifyServiceService) GetPageIterator(data url.Values) *NotifyServicePageIterator {
	iter := NewPageIterator(service.client, data, service.client.NotifyBaseUrl()+"/"+notifyServiceApiPath)
	return &NotifyServicePageIterator{
		p: iter,
	}
}

func (c *NotifyServicePageIterator) Next(ctx context.Context) (*NotifyServicePage, error) {
	cp := new(NotifyServicePage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
