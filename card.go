// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package qiwi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// CardsDetectResponse api response
type CardsDetectResponse struct {
	Code struct {
		Value json.Number `json:"value"`
		Name  string      `json:"_name"`
	} `json:"code"`
	Message json.Number `json:"message"`
}

// Detect detect card PS
func (c *Payments) DetectProviderIDByCardNumber(ctx context.Context, cardNumber string) (id int, err error) {
	var r CardsDetectResponse
	err = c.client.makePostRequest(ctx, EndpointCardsDetect, &r, url.Values{"cardNumber": {cardNumber}})
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
	PaymentMethod PaymentMethod `json:"paymentMethod"`
	Fields        struct {
		Account    string `json:"account"`               // Номер банковской карты
		RemName    string `json:"rem_name,omitempty"`    //Имя отправителя.
		RemNameF   string `json:"rem_name_f,omitempty"`  // Фамилия отправителя.
		RecAddress string `json:"rec_address,omitempty"` //Адрес отправителя
		RecCity    string `json:"rec_city,omitempty"`    //Город отправителя.
		RecCountry string `json:"rec_country,omitempty"` //Страна отправителя.
		RegName    string `json:"reg_name,omitempty"`    //Имя получателя.
		RegNameF   string `json:"reg_name_f,omitempty"`  //Фамилия получателя.
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
func (c *Payments) Payment(ctx context.Context, psID int, amount float64, account string, comments ...string) (res PaymentResponse, err error) {
	req := PaymentRequest{
		ID: strconv.Itoa(int(time.Now().Unix()) * 1000),
	}
	// constants
	req.PaymentMethod.Type = "Account"
	req.PaymentMethod.AccountID = CurrencyRUB

	req.Sum.Amount = amount
	req.Sum.Currency = CurrencyRUB
	req.Fields.Account = account

	if len(comments) > 0 {
		req.Comment = comments[0]
	}

	endpoint := fmt.Sprintf(EndpointPayment, psID)

	err = c.client.makePostRequest(ctx, endpoint, &res, req)
	if err != nil {
		return
	}

	return
}
