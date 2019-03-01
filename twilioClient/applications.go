package twilioClient

import (
	"golang.org/x/net/context"
	"net/url"
)

/**
 * Doc: https://www.twilio.com/docs/usage/api/applications
 */
const appApiPath string = "Applications"

type Application struct {
	AccountSid            string `json:"account_sid"`
	APIVersion            string `json:"api_version"`
	DateCreated           string `json:"date_created"`
	DateUpdated           string `json:"date_updated"`
	FriendlyName          string `json:"friendly_name"`
	MessageStatusCallback string `json:"message_status_callback"`
	Sid                   string `json:"sid"`
	SMSFallbackMethod     string `json:"sms_fallback_method"`
	SMSFallbackURL        string `json:"sms_fallback_url"`
	SMSURL                string `json:"sms_url"`
	StatusCallback        string `json:"status_callback"`
	StatusCallbackMethod  string `json:"status_callback_method"`
	URI                   string `json:"uri"`
	VoiceCallerIDLookup   bool   `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod   string `json:"voice_fallback_method"`
	VoiceFallbackURL      string `json:"voice_fallback_url"`
	VoiceMethod           string `json:"voice_method"`
	VoiceURL              string `json:"voice_url"`
}

type ApplicationPage struct {
	Page
	Applications []*Application `json:"applications"`
}

type ApplicationService struct {
	client *TwilioClient
}

func (service *ApplicationService) Create(ctx context.Context, data url.Values) (*Application, error) {
	var targetUrl = service.client.BaseUrl() + "/" + appApiPath + ".json"
	app := new(Application)
	err := service.client.post(targetUrl, data, app)
	return app, err
}

func (service *ApplicationService) Update(ctx context.Context, appSid string, data url.Values) (*Application, error) {
	var targetUrl = service.client.BaseUrl() + "/" + appApiPath + "/" + appSid + ".json"

	app := new(Application)
	err := service.client.post(targetUrl, data, app)
	return app, err
}

func (service *ApplicationService) Delete(ctx context.Context, appSid string) error {
	var targetUrl = service.client.BaseUrl() + "/" + appApiPath + "/" + appSid + ".json"

	return service.client.delete(targetUrl)
}

func (service *ApplicationService) Get(ctx context.Context, appSid string) (*Application, error) {
	var targetUrl = service.client.BaseUrl() + "/" + appApiPath + "/" + appSid + ".json"

	app := new(Application)
	err := service.client.get(targetUrl, url.Values{}, app)

	return app, err
}

func (service *ApplicationService) GetPage(ctx context.Context, data url.Values) (*ApplicationPage, error) {
	iter := service.GetPageIterator(data)
	return iter.Next(ctx)
}

type ApplicationPageIterator struct {
	p *PageIterator
}

func (service *ApplicationService) GetPageIterator(data url.Values) *ApplicationPageIterator {
	iter := NewPageIterator(service.client, data, service.client.BaseUrl()+"/"+appApiPath+".json")
	return &ApplicationPageIterator{
		p: iter,
	}
}

func (c *ApplicationPageIterator) Next(ctx context.Context) (*ApplicationPage, error) {
	cp := new(ApplicationPage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
