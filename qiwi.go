// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package qiwi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	// BaseURL represent base url
	BaseURL = "https://edge.qiwi.com"
	// OpenBaseURL open base url
	OpenBaseURL = "https://qiwi.com"
	// VersionAPI current qiwi.com api version
	VersionAPI = "v1"

	// QiwiProviderID used for payment to qiwi wallet
	QiwiProviderID = 99

	// CurrencyRUB rub id
	CurrencyRUB = "643"
	// CurrencyUSD usd currency id
	CurrencyUSD = "840"
	// CurrencyEUR euro currency id
	CurrencyEUR = "978"
	// CurrencyKZT tenge currency id
	CurrencyKZT = "398"
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
	// EndpointPayment send money from wallet
	EndpointPayment = "sinap/api/v2/terms/%d/payments"
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

	debug bool
}

func (c *Client) makeRequest(ctx context.Context, endpoint string, res interface{}, params ...url.Values) (err error) {
	var param url.Values
	if len(params) > 0 {
		param = params[0]
	}
	return c.req(ctx, "GET", endpoint, res, param)
}

func (c *Client) makePostRequest(ctx context.Context, endpoint string, res interface{}, params ...interface{}) (err error) {
	return c.req(ctx, "POST", endpoint, res, params...)
}

func (c *Client) req(ctx context.Context, method, endpoint string, res interface{}, params ...interface{}) (err error) {

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
					log.Printf("body: %v", v.Encode())
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
					log.Printf("body: %s", bts)
				}
				body = bytes.NewReader(bts)
			}

		}
	}

	if c.debug {
		log.Printf("Request %s", uri)
	}

	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return
	}
	req = req.WithContext(ctx)
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
		log.Printf("token %s", c.token)
		log.Printf("%v", req.Header)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if c.debug {
		log.Printf("response body: %s", respBytes)
	}

	if resp.StatusCode != 200 {
		var errResponse ErrorResponse
		err = json.Unmarshal(respBytes, &errResponse)
		if err != nil && c.debug {
			log.Printf("err unmarshal qiwi error response err=%s", err)
		}
		if errResponse.Code != "" {
			intCode, _ := strconv.Atoi(strings.TrimPrefix(errResponse.Code, "QWPRC-"))
			if err, found := codeToError[intCode]; found {
				return err
			}
			return fmt.Errorf("%s", errResponse.Message)
		}
		if err, has := codeToError[resp.StatusCode]; has {
			return err
		}
		return fmt.Errorf(http.StatusText(resp.StatusCode))
	}

	return c.decodeResponse(ioutil.NopCloser(bytes.NewReader(respBytes)), res)
}

type ErrorResponse struct {
	Code    string
	Message string
}

// SetWallet set wallet for client
func (c *Client) SetWallet(wallet string) {
	c.wallet = wallet
}
