package twilioClient

import (
	"golang.org/x/net/context"
	"net/url"
	"strings"
)

/**
 * Doc: https://www.twilio.com/docs/chat/rest/credentials
 */
const pushCredentialApiPath string = "Credentials"

type PushCredential struct {
	Sid          string             `json:"sid"`
	AccountSid   string             `json:"account_sid"`
	DateCreated  string             `json:"date_created"`
	DateUpdated  string             `json:"date_updated"`
	FriendlyName string             `json:"friendly_name"`
	Type         PushCredentialType `json:"type"`
	Sandbox      bool               `json:"sandbox"`
	Url          string             `json:"url"`
}

type PushCredentialType string

const Apn = PushCredentialType("apn")
const Fcm = PushCredentialType("fcm")
const Gcm = PushCredentialType("gcm")

func (s PushCredentialType) ToString() string {
	switch s {
	case Apn:
		return "apn"
	case Fcm:
		return "fcm"
	case Gcm:
		return "gcm"
	default:
		return strings.Title(string(s))
	}
}

type PushCredentialPage struct {
	Page
	PushCredentials []*PushCredential `json:"credentials"`
}

type PushCredentialService struct {
	client *TwilioClient
}

func (service *PushCredentialService) Create(ctx context.Context, data url.Values) (*PushCredential, error) {
	var targetUrl = service.client.ChatBaseUrl() + "/" + pushCredentialApiPath
	app := new(PushCredential)
	err := service.client.post(targetUrl, data, app)
	return app, err
}

func (service *PushCredentialService) Update(ctx context.Context, appSid string, data url.Values) (*PushCredential, error) {
	var targetUrl = service.client.ChatBaseUrl() + "/" + pushCredentialApiPath + "/" + appSid

	app := new(PushCredential)
	err := service.client.post(targetUrl, data, app)
	return app, err
}

func (service *PushCredentialService) Delete(ctx context.Context, appSid string) error {
	var targetUrl = service.client.ChatBaseUrl() + "/" + pushCredentialApiPath + "/" + appSid

	return service.client.delete(targetUrl)
}

func (service *PushCredentialService) Get(ctx context.Context, appSid string) (*PushCredential, error) {
	var targetUrl = service.client.ChatBaseUrl() + "/" + pushCredentialApiPath + "/" + appSid

	app := new(PushCredential)
	err := service.client.get(targetUrl, url.Values{}, app)

	return app, err
}

func (service *PushCredentialService) GetPage(ctx context.Context, data url.Values) (*PushCredentialPage, error) {
	iter := service.GetPageIterator(data)
	return iter.Next(ctx)
}

type PushCredentialPageIterator struct {
	p *PageIterator
}

func (service *PushCredentialService) GetPageIterator(data url.Values) *PushCredentialPageIterator {
	iter := NewPageIterator(service.client, data, service.client.ChatBaseUrl()+"/"+pushCredentialApiPath)
	return &PushCredentialPageIterator{
		p: iter,
	}
}

func (c *PushCredentialPageIterator) Next(ctx context.Context) (*PushCredentialPage, error) {
	cp := new(PushCredentialPage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
