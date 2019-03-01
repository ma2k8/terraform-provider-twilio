package twilioClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type TwilioClient struct {
	accountSid, authToken, baseUrl, chatBaseUrl, notifyBaseUrl string

	HttpClient   *http.Client
	RetryTimeout time.Duration

	Accounts *AccountService
}

func NewTwilioClient(accountSid, authToken string) *TwilioClient {
	baseUrl := os.Getenv("TWILIO_BASE_URL")
	if baseUrl == "" {
		baseUrl = "https://api.twilio.com/2010-04-01"
	}

	chatBaseUrl := os.Getenv("TWILIO_CHAT_BASE_URL")
	if chatBaseUrl == "" {
		chatBaseUrl = "https://chat.twilio.com/v2/"
	}

	notifyBaseUrl := os.Getenv("TWILIO_NOTIFY_BASE_URL")
	if notifyBaseUrl == "" {
		notifyBaseUrl = "https://notify.twilio.com/v1/"
	}

	timeoutSec, _ := strconv.Atoi(os.Getenv("TWILIO_API_WAIT"))
	if timeoutSec == 0 {
		timeoutSec = 60
	}

	c := &TwilioClient{
		accountSid:    accountSid,
		authToken:     authToken,
		baseUrl:       baseUrl,
		chatBaseUrl:   chatBaseUrl,
		notifyBaseUrl: notifyBaseUrl,
		HttpClient:    http.DefaultClient,
		RetryTimeout:  time.Duration(timeoutSec) * time.Second,
	}

	c.Accounts = &AccountService{client: c}

	return c
}

func (client *TwilioClient) get(targetUrl string, queryParams url.Values, v interface{}) error {
	if queryParams == nil {
		queryParams = url.Values{}
	}

	req, err := http.NewRequest("GET", targetUrl, nil)
	req.URL.RawQuery = queryParams.Encode()

	if err != nil {
		return err
	}

	req.SetBasicAuth(client.AccountSid(), client.AuthToken())
	httpClient := &http.Client{}

	res, err := httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		if res.StatusCode == 500 {
			return Error{"Server Error"}
		} else {
			twilioError := new(TwilioError)
			_ = json.Unmarshal(body, twilioError)
			return twilioError
		}
	}

	return json.Unmarshal(body, v)
}

func (client *TwilioClient) post(targetUrl string, data url.Values, v interface{}) error {
	req, err := http.NewRequest("POST", targetUrl, strings.NewReader(data.Encode()))

	if err != nil {
		return err
	}

	req.SetBasicAuth(client.AccountSid(), client.AuthToken())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpClient := &http.Client{}

	res, err := httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		if res.StatusCode == 500 {
			return Error{"Server Error"}
		} else {
			twilioError := new(TwilioError)
			json.Unmarshal(body, twilioError)
			return twilioError
		}
	}

	return json.Unmarshal(body, v)
}

func (client *TwilioClient) delete(targetUrl string) error {
	req, err := http.NewRequest("DELETE", targetUrl, nil)

	if err != nil {
		return err
	}

	req.SetBasicAuth(client.AccountSid(), client.AuthToken())
	httpClient := &http.Client{}

	res, err := httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 204 {
		return fmt.Errorf("Non-204 returned from server for DELETE: %d", res.StatusCode)
	}

	return nil
}

func (client *TwilioClient) AccountSid() string {
	return client.accountSid
}

func (client *TwilioClient) AuthToken() string {
	return client.authToken
}

func (client *TwilioClient) BaseUrl() string {
	return client.baseUrl
}

func (client *TwilioClient) ChatBaseUrl() string {
	return client.chatBaseUrl
}

func (client *TwilioClient) NotifyBaseUrl() string {
	return client.notifyBaseUrl
}
