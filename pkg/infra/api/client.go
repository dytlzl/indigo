package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dytlzl/indigo/pkg/config"
)

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=mock_$GOPACKAGE

type Client interface {
	GenerateAccessToken() (*string, error)
	Get(ctx context.Context, endpoint string) ([]byte, error)
	Post(ctx context.Context, endpoint string, body io.Reader) ([]byte, error)
	Put(ctx context.Context, endpoint string, body io.Reader) ([]byte, error)
	Delete(ctx context.Context, endpoint string) ([]byte, error)
}

// Client -
type client struct {
	conf       config.Config
	HTTPClient *http.Client
}

// AuthStruct -
type AuthStruct struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Code         string `json:"code"`
	GrantType    string `json:"grantType"`
}

// AuthResponse -
type AuthResponse struct {
	AccessToken string `json:"accessToken"`
}

// NewClient -
func NewClient(conf config.Config) Client {
	c := client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		conf:       conf,
	}
	return &c
}

func (c *client) GenerateAccessToken() (*string, error) {
	log.Println("logging in...")
	time.Sleep(time.Second)
	apiKey, apiSecret := c.conf.GetCredential()

	if apiKey == "" || apiSecret == "" {
		return nil, fmt.Errorf("define api key and api secret")
	}
	auth := AuthStruct{
		ClientID:     apiKey,
		ClientSecret: apiSecret,
		GrantType:    "client_credentials",
		Code:         "",
	}
	rb, err := json.Marshal(auth)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.customer.jp/oauth/v1/accesstokens", strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doSignInRequest(req)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", err, string(rb))
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar.AccessToken, nil
}

func (c *client) doSignInRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s, requestHeader: %v", res.StatusCode, body, req.Header)
	}

	return body, err
}

func (c *client) Get(ctx context.Context, endpoint string) ([]byte, error) {
	return c.doRequestWithRetry(ctx, "GET", endpoint, nil)
}

func (c *client) Post(ctx context.Context, endpoint string, body io.Reader) ([]byte, error) {
	return c.doRequestWithRetry(ctx, "POST", endpoint, body)
}

func (c *client) Put(ctx context.Context, endpoint string, body io.Reader) ([]byte, error) {
	return c.doRequestWithRetry(ctx, "PUT", endpoint, body)
}

func (c *client) Delete(ctx context.Context, endpoint string) ([]byte, error) {
	return c.doRequestWithRetry(ctx, "DELETE", endpoint, nil)
}

func (c *client) doRequestWithRetry(ctx context.Context, method, endpoint string, reqBody io.Reader) ([]byte, error) {
	if c.conf.Token() == "" {
		token, err := c.GenerateAccessToken()
		if err != nil {
			return nil, err
		}
		err = c.conf.SetToken(*token)
		if err != nil {
			log.Println("failed to persist token")
		}
	}
	res, resBody, err := c.doRequest(ctx, method, endpoint, reqBody)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusUnauthorized {
		log.Println("token expired.")
		token, err := c.GenerateAccessToken()
		if err != nil {
			return nil, err
		}
		err = c.conf.SetToken(*token)
		if err != nil {
			log.Println("failed to persist token")
		}
		res, resBody, err = c.doRequest(ctx, method, endpoint, reqBody)
		if err != nil {
			return nil, err
		}
	}

	if !(res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated) {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, resBody)
	}
	return resBody, err
}

func (c *client) doRequest(ctx context.Context, method, endpoint string, reqBody io.Reader) (*http.Response, []byte, error) {
	time.Sleep(2 * time.Second)
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s%s", "https://api.customer.jp/webarenaIndigo/v1", endpoint), reqBody)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.conf.Token()))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return res, nil, err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res, nil, err
	}

	return res, resBody, err
}
