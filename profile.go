// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package qiwi

import "time"

// Profile for profile endpoints
type Profile struct {
	client *Client
}

// NewProfile returns new Profile obj
func NewProfile(c *Client) *Profile {
	return &Profile{client: c}
}

// Current call api and get current user profile
func (h *Profile) Current() (pr ProfileResponse, err error) {
	err = h.client.makeRequest(EndpointProfile, &pr)
	if err != nil {
		return
	}

	return
}

// ProfileResponse reponse
type ProfileResponse struct {
	AuthInfo struct {
		BoundEmail    string    `json:"boundEmail"`
		IP            string    `json:"ip"`
		LastLoginDate time.Time `json:"lastLoginDate"`
		MobilePinInfo struct {
			LastMobilePinChange time.Time `json:"lastMobilePinChange"`
			MobilePinUsed       bool      `json:"mobilePinUsed"`
			NextMobilePinChange time.Time `json:"nextMobilePinChange"`
		} `json:"mobilePinInfo"`
		PassInfo struct {
			LastPassChange time.Time `json:"lastPassChange"`
			NextPassChange time.Time `json:"nextPassChange"`
			PasswordUsed   bool      `json:"passwordUsed"`
		} `json:"passInfo"`
		PersonID int64 `json:"personId"`
		PinInfo  struct {
			PinUsed bool `json:"pinUsed"`
		} `json:"pinInfo"`
		RegistrationDate time.Time `json:"registrationDate"`
	} `json:"authInfo"`
	ContractInfo struct {
		Blocked            bool          `json:"blocked"`
		ContractID         int64         `json:"contractId"`
		CreationDate       time.Time     `json:"creationDate"`
		Features           []interface{} `json:"features"`
		IdentificationInfo []struct {
			BankAlias           string `json:"bankAlias"`
			IdentificationLevel string `json:"identificationLevel"`
		} `json:"identificationInfo"`
	} `json:"contractInfo"`
	UserInfo struct {
		DefaultPayCurrency int         `json:"defaultPayCurrency"`
		DefaultPaySource   int         `json:"defaultPaySource"`
		Email              interface{} `json:"email"`
		FirstTxnID         int64       `json:"firstTxnId"`
		Language           string      `json:"language"`
		Operator           string      `json:"operator"`
		PhoneHash          string      `json:"phoneHash"`
		PromoEnabled       interface{} `json:"promoEnabled"`
	} `json:"userInfo"`
}
