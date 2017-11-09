// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package qiwi

// Opt returns func for changing client
// used for New
type Opt func(c *Client)

// Wallet set wallet to client
// example:
// qw := qiwi.New(token, qiwi.Wallet("70000000000"))
func Wallet(wallet string) func(c *Client) {
	// TODO: add validation
	return func(c *Client) {
		c.wallet = wallet
	}
}

// Debug enable debug
func Debug(c *Client) {
	c.debug = true
}
