// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package qiwi

// Balance for payment-history endpoints
type Balance struct {
	client *Client
}

// NewBalance returns new Balance obj
func NewBalance(c *Client) *Balance {
	return &Balance{client: c}
}

// Current call api and get current user balance
func (b *Balance) Current() (hr BalanceResponse, err error) {
	err = b.client.makeRequest(EndpointBalance, &hr)
	if err != nil {
		return
	}

	return
}

// BalanceResponse response of balance request
type BalanceResponse struct {
	Accounts []struct {
		Alias   string `json:"alias"`
		FsAlias string `json:"fsAlias"`
		Title   string `json:"title"`
		Type    struct {
			ID    string `json:"id"`
			Title string `json:"title"`
		} `json:"type"`
		HasBalance bool `json:"hasBalance"`
		Balance    struct {
			Amount   float64 `json:"amount"`
			Currency int     `json:"currency"`
		} `json:"balance"`
		Currency int `json:"currency"`
	} `json:"accounts"`
}
