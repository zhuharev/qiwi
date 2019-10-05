// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package qiwi

import "context"

// Balance for payment-history endpoints
type Balance struct {
	client *Client
}

// NewBalance returns new Balance obj
func NewBalance(c *Client) *Balance {
	return &Balance{client: c}
}

// Current call api and get current user balance
func (b *Balance) Current(ctx context.Context) (hr BalanceResponse, err error) {
	err = b.client.makeRequest(ctx, EndpointBalance, &hr)
	if err != nil {
		return
	}

	return
}

// BalanceResponse response of balance request
type BalanceResponse struct {
	Accounts []Account
}
