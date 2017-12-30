// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package qiwi

import (
	"fmt"
	"net/url"
	"time"
)

// Payments for payment-history endpoints
type Payments struct {
	client *Client
}

// NewPayments returns new Payments obj
func NewPayments(c *Client) *Payments {
	return &Payments{client: c}
}

// History call api and get payments history
func (p *Payments) History(rows uint, params ...url.Values) (pr PaymentsResponse, err error) {
	param := url.Values{}

	{
		if len(params) > 0 {
			param = params[0]
		}
		param["rows"] = []string{fmt.Sprint(rows)}
	}

	err = p.client.makeRequest(EndpointPaymentsHistory, &pr, param)
	if err != nil {
		return
	}

	return
}

// PaymentsResponse api response format
type PaymentsResponse struct {
	Data        []Txn       `json:"data"`
	NextTxnID   interface{} `json:"nextTxnId"`
	NextTxnDate interface{} `json:"nextTxnDate"`
}

// Txn represent qiwi transaction
type Txn struct {
	TxnID      int64     `json:"txnId"`
	PersonID   int64     `json:"personId"`
	Date       time.Time `json:"date"`
	ErrorCode  int       `json:"errorCode"`
	Error      string    `json:"error"`
	Status     string    `json:"status"`
	Type       string    `json:"type"`
	StatusText string    `json:"statusText"`
	TrmTxnID   string    `json:"trmTxnId"`
	Account    string    `json:"account"`
	Sum        struct {
		Amount   float64 `json:"amount"`
		Currency int     `json:"currency"`
	} `json:"sum"`
	Commission struct {
		Amount   float64 `json:"amount"`
		Currency int     `json:"currency"`
	} `json:"commission"`
	Total struct {
		Amount   float64 `json:"amount"`
		Currency int     `json:"currency"`
	} `json:"total"`
	Provider struct {
		ID          int           `json:"id"`
		ShortName   string        `json:"shortName"`
		LongName    string        `json:"longName"`
		LogoURL     string        `json:"logoUrl"`
		Description string        `json:"description"`
		Keys        string        `json:"keys"`
		SiteURL     string        `json:"siteUrl"`
		Extras      []interface{} `json:"extras"`
	} `json:"provider"`
	Source                 interface{}   `json:"source"`
	Comment                string        `json:"comment"`
	CurrencyRate           float64       `json:"currencyRate"`
	Extras                 []interface{} `json:"extras"`
	ChequeReady            bool          `json:"chequeReady"`
	BankDocumentAvailable  bool          `json:"bankDocumentAvailable"`
	BankDocumentReady      bool          `json:"bankDocumentReady"`
	RepeatPaymentEnabled   bool          `json:"repeatPaymentEnabled"`
	FavoritePaymentEnabled bool          `json:"favoritePaymentEnabled"`
	RegularPaymentEnabled  bool          `json:"regularPaymentEnabled"`
}

// StatResponse response of stat endpoind
type StatResponse struct {
	IncomingTotal []struct {
		Amount   float64 `json:"amount"`
		Currency int     `json:"currency"`
	} `json:"incomingTotal"`
	OutgoingTotal []struct {
		Amount   float64 `json:"amount"`
		Currency int     `json:"currency"`
	} `json:"outgoingTotal"`
}

// Stat get sum of incoming and outgoing payments
func (p *Payments) Stat(startDate, endDate time.Time, params ...url.Values) (res StatResponse, err error) {
	param := url.Values{}

	{
		if len(params) > 0 {
			param = params[0]
		}
		param["startDate"] = []string{startDate.Format(time.RFC3339)}
		param["endDate"] = []string{endDate.Format(time.RFC3339)}
	}

	err = p.client.makeRequest(EndpointStat, &res, param)
	if err != nil {
		return
	}
	return
}

// ComissionResponse  json unserialazer
type ComissionResponse struct {
	Content struct {
		Terms struct {
			Commission struct {
				Ranges []struct {
					Bound float64 `json:"bound"`
					Fixed float64 `json:"fixed"`
					Rate  float64 `json:"rate"`
					Min   float64 `json:"min"`
					Max   float64 `json:"max"`
				} `json:"ranges"`
			} `json:"commission"`
		} `json:"terms"`
	} `json:"content"`
}

// Comission get provider comission
func (p *Payments) Comission(providerID int) (res ComissionResponse, err error) {
	var (
		endpoint = fmt.Sprintf(EndpointComission, providerID)
	)
	err = p.client.makeRequest(endpoint, &res)
	if err != nil {
		return
	}
	return
}

// SpecialComissionRequest json serializer
type SpecialComissionRequest struct {
	Account       string `json:"account"`
	PaymentMethod struct {
		Type      string `json:"type"`
		AccountID string `json:"accountId"`
	} `json:"paymentMethod"`
	PurchaseTotals struct {
		Total struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		} `json:"total"`
	} `json:"purchaseTotals"`
}

// SpecialComissionResponse json unserialazer
type SpecialComissionResponse struct {
	ProviderID  int `json:"providerId"`
	WithdrawSum struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"withdrawSum"`
	EnrollmentSum struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"enrollmentSum"`
	QwCommission struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"qwCommission"`
	FundingSourceCommission struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"fundingSourceCommission"`
	WithdrawToEnrollmentRate float64 `json:"withdrawToEnrollmentRate"`
}

// SpecialComission get provider comission
func (p *Payments) SpecialComission(providerID int, to string, amount float64) (res SpecialComissionResponse, err error) {
	req := SpecialComissionRequest{Account: to}
	req.PaymentMethod.Type = "Account"
	req.PaymentMethod.AccountID = CurrencyRUB
	req.PurchaseTotals.Total.Amount = amount
	req.PurchaseTotals.Total.Currency = CurrencyRUB

	var (
		endpoint = fmt.Sprintf(EndpointSpecialComission, providerID)
	)
	err = p.client.makePostRequest(endpoint, &res, req)
	return
}
