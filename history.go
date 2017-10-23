package qiwi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// History for payment-history endpoints
type History struct {
	client *Client
}

// NewHistory returns new Histpry obj
func NewHistory(c *Client) *History {
	return &History{client: c}
}

// Payments call api and get payments history
func (h *History) Payments(rows uint) (hr *PaymentsResponse, err error) {
	body, err := h.client.makeRequest(EndpointPaymentsHistory, url.Values{"rows": {fmt.Sprint(rows)}})
	if err != nil {
		return
	}
	defer body.Close()

	hr = new(PaymentsResponse)

	dec := json.NewDecoder(body)
	err = dec.Decode(hr)
	return
}

// PaymentsResponse api response format
type PaymentsResponse struct {
	Data []struct {
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
		CurrencyRate           int           `json:"currencyRate"`
		Extras                 []interface{} `json:"extras"`
		ChequeReady            bool          `json:"chequeReady"`
		BankDocumentAvailable  bool          `json:"bankDocumentAvailable"`
		BankDocumentReady      bool          `json:"bankDocumentReady"`
		RepeatPaymentEnabled   bool          `json:"repeatPaymentEnabled"`
		FavoritePaymentEnabled bool          `json:"favoritePaymentEnabled"`
		RegularPaymentEnabled  bool          `json:"regularPaymentEnabled"`
	} `json:"data"`
	NextTxnID   interface{} `json:"nextTxnId"`
	NextTxnDate interface{} `json:"nextTxnDate"`
}
