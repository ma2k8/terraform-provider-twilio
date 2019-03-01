package twilioClient

import (
	"golang.org/x/net/context"
	"net/url"
	"strings"
)

/**
 * Doc: https://jp.twilio.com/docs/iam/api/subaccounts
 */
const accountApiPath string = "Accounts"

type Account struct {
	Sid             string            `json:"sid"`
	FriendlyName    string            `json:"friendly_name"`
	Type            AccountType       `json:"type"`
	AuthToken       string            `json:"auth_token"`
	OwnerAccountSid string            `json:"owner_account_sid"`
	DateCreated     string            `json:"date_created"`
	DateUpdated     string            `json:"date_updated"`
	Status          AccountStatus     `json:"status"`
	SubresourceURIs map[string]string `json:"subresource_uris"`
	URI             string            `json:"uri"`
}

type AccountPage struct {
	Page
	Accounts []Account `json:"accounts"`
}

type AccountStatus string

const Active = AccountStatus("active")
const Suspended = AccountStatus("suspended")
const Closed = AccountStatus("closed")

func (s AccountStatus) ToString() string {
	switch s {
	case Active:
		return "active"
	case Suspended:
		return "suspended"
	case Closed:
		return "closed"
	default:
		return strings.Title(string(s))
	}
}

type AccountType string

const Trial = AccountType("Trial")
const Full = AccountType("Full")

func (s AccountType) ToString() string {
	switch s {
	case Trial:
		return "Trial"
	case Full:
		return "Full"
	default:
		return strings.Title(string(s))
	}
}

type AccountService struct {
	client *TwilioClient
}

func (service *AccountService) Create(ctx context.Context, data url.Values) (*Account, error) {
	var targetUrl = service.client.BaseUrl() + "/" + accountApiPath + ".json"
	account := new(Account)
	err := service.client.post(targetUrl, data, account)
	return account, err
}

func (service *AccountService) Update(ctx context.Context, accountSid string, data url.Values) (*Account, error) {
	var targetUrl = service.client.BaseUrl() + "/" + accountApiPath + "/" + accountSid + ".json"

	account := new(Account)
	err := service.client.post(targetUrl, data, account)
	return account, err
}

func (service *AccountService) Delete(ctx context.Context, accountSid string) (*Account, error) {
	var targetUrl = service.client.BaseUrl() + "/" + accountApiPath + "/" + accountSid + ".json"

	data := url.Values{}
	data.Set("Status", Closed.ToString())
	account := new(Account)
	err := service.client.post(targetUrl, data, account)
	return account, err
}

func (service *AccountService) Get(ctx context.Context, accountSid string) (*Account, error) {
	var targetUrl = service.client.BaseUrl() + "/" + accountApiPath + "/" + accountSid + ".json"

	account := new(Account)
	err := service.client.get(targetUrl, url.Values{}, account)

	return account, err
}

func (service *AccountService) GetPage(ctx context.Context, data url.Values) (*AccountPage, error) {
	iter := service.GetPageIterator(data)
	return iter.Next(ctx)
}

type AccountPageIterator struct {
	p *PageIterator
}

func (service *AccountService) GetPageIterator(data url.Values) *AccountPageIterator {
	iter := NewPageIterator(service.client, data, service.client.BaseUrl()+"/"+accountApiPath+".json")
	return &AccountPageIterator{
		p: iter,
	}
}

func (c *AccountPageIterator) Next(ctx context.Context) (*AccountPage, error) {
	cp := new(AccountPage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
