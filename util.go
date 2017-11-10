package qiwi

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func (q *Client) decodeResponse(body io.ReadCloser, res interface{}) (err error) {
	if q.debug {
		bts, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bts, res)
		if err != nil {
			return err
		}
		log.Printf("response: %s", bts)
		return nil
	}
	dec := json.NewDecoder(body)
	err = dec.Decode(res)
	return
}
