## Qiwi [![GoDoc](https://godoc.org/github.com/zhuharev/qiwi?status.svg)](http://godoc.org/github.com/zhuharev/qiwi)

An golang qiwi.com api client

You need to get access token [here](https://qiwi.com/api)

## Install

`go get -u github.com/zhuharev/qiwi`

## Usage

```go
token = "YOUR_TOKEN"
qw := qiwi.New(token)

// get last 10 payments
resp, err := qw.History.Payments(10)
if err != nil {
  // handle error
}
log.Printf("%v", resp)
```
