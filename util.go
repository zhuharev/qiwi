package qiwi

import (
	"encoding/json"
	"io"
)

func (q *Client) decodeResponse(body io.ReadCloser, res interface{}) (err error) {
	dec := json.NewDecoder(body)
	err = dec.Decode(res)
	return
}
