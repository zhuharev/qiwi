// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package qiwi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/fatih/color"
)

const (
	// BaseURL represent base url
	BaseURL = "https://edge.qiwi.com"
	// OpenBaseURL open base url
	OpenBaseURL = "https://qiwi.com"
	// VersionAPI current qiwi.com api version
	VersionAPI = "v1"

	// CurrencyRUB rub id
	CurrencyRUB = "643"
)

const (
	// EndpointProfile endpoint
	EndpointProfile = "person-profile/v1/profile/current"
	// EndpointIdent identification endpoint
	EndpointIdent = "identification/v1/persons/%s/identification" // %s - wallet
	// EndpointPaymentsHistory get history
	EndpointPaymentsHistory = "payment-history/v1/persons/%s/payments" // %s - wallet
	// EndpointStat get stat of payments
	EndpointStat = "payment-history/v1/persons/%s/payments/total" // %s - wallet
	// EndpointTxnInfo get info anout single txn
	EndpointTxnInfo = "payment-history/v1/transactions/%s" // %s - txn_id
	// EndpointBalance get wallet balance
	EndpointBalance = "funding-sources/v1/accounts/current"
	// EndpointCardsDetect detect code of PS
	EndpointCardsDetect = "card/detect.action"
	// EndpointCardsPayment send money from wallet
	EndpointCardsPayment = "sinap/api/v2/terms/%d/payments"
	// EndpointComission a provider comission
	EndpointComission = "sinap/providers/%d/form"
	// EndpointSpecialComission comission for specific amount
	EndpointSpecialComission = "sinap/providers/%d/onlineCommission"
)

var (
	// OpenMethods not require Bearer token
	OpenMethods = map[string]bool{
		EndpointCardsDetect: true,
	}
)

// New returns client
func New(token string, opts ...Opt) *Client {
	c := &Client{
		token:       token,
		baseURL:     BaseURL,
		openBaseURL: OpenBaseURL,
		httpClient:  http.DefaultClient,
		debug:       false,
	}

	c.Payments = NewPayments(c)
	c.Profile = NewProfile(c)
	c.Balance = NewBalance(c)
	c.Cards = NewCards(c)

	for _, fn := range opts {
		fn(c)
	}
	return c
}

// Client main struct
type Client struct {
	baseURL     string
	openBaseURL string
	token       string
	wallet      string

	httpClient *http.Client

	Payments *Payments
	Profile  *Profile
	Balance  *Balance
	Cards    *Cards

	debug bool
}

func (c *Client) makeRequest(endpoint string, res interface{}, params ...url.Values) (err error) {
	var param url.Values
	if len(params) > 0 {
		param = params[0]
	}
	return c.req("GET", endpoint, res, param)
}

func (c *Client) makePostRequest(endpoint string, res interface{}, params ...interface{}) (err error) {
	return c.req("POST", endpoint, res, params...)
}

func (c *Client) req(method, endpoint string, res interface{}, params ...interface{}) (err error) {

	var (
		needWalletID = []string{
			EndpointPaymentsHistory,
			EndpointStat,
		}
		isOpenMethod = OpenMethods[endpoint]
		baseURL      = c.baseURL
	)

	if isOpenMethod {
		baseURL = c.openBaseURL
	}

	for _, withNeedWalletID := range needWalletID {
		if endpoint == withNeedWalletID {
			endpoint = fmt.Sprintf(endpoint, c.wallet)
		}
	}

	uri := fmt.Sprintf("%s/%s", baseURL, endpoint)
	var body io.Reader

	if len(params) > 0 && params[0] != nil {
		if method == "GET" {
			param := params[0].(url.Values)
			if len(param) > 0 {
				query := param.Encode()
				uri = fmt.Sprintf("%s?%s", uri, query)
			}
		} else {
			switch v := params[0].(type) {
			case url.Values:
				if c.debug {
					color.Cyan("body: %v", v.Encode())
				}
				body = strings.NewReader(v.Encode())
			case PaymentRequest,
				SpecialComissionRequest:
				var bts []byte
				bts, err = json.Marshal(v)
				if err != nil {
					return
				}
				if c.debug {
					color.Cyan("body: %s", bts)
				}
				body = bytes.NewReader(bts)
			}

		}
	}

	if c.debug {
		color.Green("Request %s", uri)
	}

	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return
	}
	req.Header.Set("Accept", "application/json")
	if !isOpenMethod {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}
	if method == "POST" {
		if len(params) > 0 {
			switch params[0].(type) {
			case url.Values:
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			case PaymentRequest, SpecialComissionRequest:
				req.Header.Set("Content-Type", "application/json")
			}
		}
	}
	if c.debug {
		color.Cyan("token %s", c.token)
		color.Cyan("%v", req.Header)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if err, has := codeToError[resp.StatusCode]; has {
			return err
		}
		return fmt.Errorf(http.StatusText(resp.StatusCode))
	}

	return c.decodeResponse(resp.Body, res)
}

// SetWallet set wallet for client
func (c *Client) SetWallet(wallet string) {
	c.wallet = wallet
}
