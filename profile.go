package qiwi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

// History for payment-history endpoints
type Profile struct {
	client *Client
}

// NewProfile returns new Profile obj
func NewProfile(c *Client) *Profile {
	return &Profile{client: c}
}

// Current call api and get current user profile
func (h *Profile) Current() (hr *ProfileResponse, err error) {
	body, err := h.client.makeRequest(EndpointProfile)
	if err != nil {
		return
	}
	defer body.Close()

	bts, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}

	buf := bytes.NewReader(bts)

	log.Printf("[profile resp] %s", bts)

	hr = new(ProfileResponse)

	dec := json.NewDecoder(buf)
	err = dec.Decode(hr)
	return
}

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
