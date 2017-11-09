## Qiwi [![GoDoc](https://godoc.org/github.com/zhuharev/qiwi?status.svg)](http://godoc.org/github.com/zhuharev/qiwi)

An golang qiwi.com api client

Qiwi.com - CIS and Russian Payment system and this package relevant for Russian-speaking users. All the following documentation will be in Russian.

Чтобы получить токен доступа, перейдите по [ссылке](https://qiwi.com/api)

## Установка

`go get -u github.com/zhuharev/qiwi`

## Использование

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
