package twilioClient

import (
	"golang.org/x/net/context"
	"net/url"
)

/**
 * Doc: https://www.twilio.com/docs/api/rest/keys
 */
const keyApiPath string = "Keys"

type Key struct {
	DateCreated  string `json:"date_created"`
	DateUpdated  string `json:"date_updated"`
	Sid          string `json:"sid"`
	FriendlyName string `json:"friendly_name"`
	Secret       string `json:"secret"`
}

type KeyPage struct {
	Page
	Keys []*Key `json:"keys"`
}

type KeyService struct {
	client *TwilioClient
}

func (service *KeyService) Create(ctx context.Context, data url.Values) (*Key, error) {
	var targetUrl = service.client.BaseUrl() + "/" + keyApiPath + ".json"
	key := new(Key)
	err := service.client.post(targetUrl, data, key)
	return key, err
}

func (service *KeyService) Update(ctx context.Context, keySid string, data url.Values) (*Key, error) {
	var targetUrl = service.client.BaseUrl() + "/" + keyApiPath + "/" + keySid + ".json"

	key := new(Key)
	err := service.client.post(targetUrl, data, key)
	return key, err
}

func (service *KeyService) Delete(ctx context.Context, keySid string) error {
	var targetUrl = service.client.BaseUrl() + "/" + keyApiPath + "/" + keySid + ".json"

	return service.client.delete(targetUrl)
}

func (service *KeyService) Get(ctx context.Context, keySid string) (*Key, error) {
	var targetUrl = service.client.BaseUrl() + "/" + keyApiPath + "/" + keySid + ".json"

	key := new(Key)
	err := service.client.get(targetUrl, url.Values{}, key)

	return key, err
}

func (service *KeyService) GetPage(ctx context.Context, data url.Values) (*KeyPage, error) {
	iter := service.GetPageIterator(data)
	return iter.Next(ctx)
}

type KeyPageIterator struct {
	p *PageIterator
}

func (service *KeyService) GetPageIterator(data url.Values) *KeyPageIterator {
	iter := NewPageIterator(service.client, data, service.client.BaseUrl()+"/"+keyApiPath+".json")
	return &KeyPageIterator{
		p: iter,
	}
}

func (c *KeyPageIterator) Next(ctx context.Context) (*KeyPage, error) {
	cp := new(KeyPage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
