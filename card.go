// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package qiwi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// Cards for payment-history endpoints
type Cards struct {
	client *Client
}

// NewCards returns new Cards obj
func NewCards(c *Client) *Cards {
	return &Cards{client: c}
}

// CardsDetectResponse api response
type CardsDetectResponse struct {
	Code struct {
		Value json.Number `json:"value"`
		Name  string      `json:"_name"`
	} `json:"code"`
	Message json.Number `json:"message"`
}

// Detect detect card PS
func (c *Cards) Detect(cardNumber string) (id int, err error) {
	var r CardsDetectResponse
	err = c.client.makePostRequest(EndpointCardsDetect, &r, url.Values{"cardNumber": {cardNumber}})
	if err != nil {
		return
	}

	if r.Code.Value.String() != "0" {
		return 0, fmt.Errorf("%s", r.Message.String())
	}

	idInt64, err := r.Message.Int64()
	if err != nil {
		return
	}
	return int(idInt64), nil
}

// PaymentRequest request of payment
type PaymentRequest struct {
	ID  string `json:"id"`
	Sum struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"sum"`
	PaymentMethod struct {
		Type      string `json:"type"`
		AccountID string `json:"accountId"`
	} `json:"paymentMethod"`
	Fields struct {
		Account string `json:"account"`
	} `json:"fields"`
	// Qiwi to qiwi related field
	Comment string `json:"comment,omitempty"`
}

// PaymentResponse foemat of payment response
type PaymentResponse struct {
	ID     string `json:"id"`
	Terms  string `json:"terms"`
	Fields struct {
		Account string `json:"account"`
	} `json:"fields"`
	Sum struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"sum"`
	Source      string `json:"source"`
	Transaction struct {
		ID    string `json:"id"`
		State struct {
			Code string `json:"code"`
		} `json:"state"`
	} `json:"transaction"`
	// Qiwi to qiwi related field
	Comment string `json:"comment,omitempty"`
}

// Payment make mayment
func (c *Cards) Payment(psID int, amount float64, cardNumber string, comments ...string) (res PaymentResponse, err error) {
	req := PaymentRequest{
		ID: strconv.Itoa(int(time.Now().Unix()) * 1000),
	}
	// constants
	req.PaymentMethod.Type = "Account"
	req.PaymentMethod.AccountID = CurrencyRUB

	req.Sum.Amount = amount
	req.Sum.Currency = CurrencyRUB
	req.Fields.Account = cardNumber

	if len(comments) > 0 {
		req.Comment = comments[0]
	}

	endpoint := fmt.Sprintf(EndpointCardsPayment, psID)

	err = c.client.makePostRequest(endpoint, &res, req)
	if err != nil {
		return
	}

	return
}
