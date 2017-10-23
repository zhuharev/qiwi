package qiwi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	// BaseURL represent base url
	BaseURL = "https://edge.qiwi.com"
	// VersionAPI current qiwi.com api version
	VersionAPI = "v1"
)

const (
	// EndpointProfile endpoint
	EndpointProfile = "person-profile/v1/profile/current"
	// EndpointIdent identification endpoint
	EndpointIdent           = "identification/v1/persons/%s/identification"  // %s - wallet
	EndpointPaymentsHistory = "payment-history/v1/persons/%s/payments"       // %s - wallet
	EndpointStat            = "payment-history/v1/persons/%s/payments/total" // %s - wallet
	EndpointTxnInfo         = "payment-history/v1/transactions/%s"           // %s - txn_id
	EndpointBalance         = "funding-sources/v1/accounts/current"
)

// New returns client
func New(token string, opts ...Opt) *Client {
	c := &Client{
		token:      token,
		baseURL:    BaseURL,
		httpClient: http.DefaultClient,
	}

	c.History = NewHistory(c)

	for _, fn := range opts {
		fn(c)
	}
	return c
}

// Client main struct
type Client struct {
	baseURL string
	token   string
	wallet  string

	httpClient *http.Client

	History *History
}

func (c *Client) makeRequest(endpoint string, params ...url.Values) (io.ReadCloser, error) {

	// TODO: refactor this
	if endpoint == EndpointPaymentsHistory {
		endpoint = fmt.Sprintf(endpoint, c.wallet)
	}

	uri := fmt.Sprintf("%s/%s", c.baseURL, endpoint)

	if len(params) > 0 {
		query := params[0].Encode()
		uri = fmt.Sprintf("%s?%s", uri, query)
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := http.DefaultClient.Do(req)
	return resp.Body, err
}
