## Qiwi [![GoDoc](https://godoc.org/github.com/zhuharev/qiwi?status.svg)](http://godoc.org/github.com/zhuharev/qiwi)

An golang qiwi.com api client

Qiwi.com - CIS and Russian Payment system and this package relevant for Russian-speaking users. All the following documentation will be in Russian.

Чтобы получить токен доступа, перейдите по [ссылке](https://qiwi.com/api)

## История изменений

Доступна в файле [CHANGELOG.md](CHANGELOG.md).

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

## Пример платежа

Для того, чтобы перевести деньги по номеру карты, мы должны определить `ID` провайдера, через которого будет проведён платёж.

```go
	// токен
	token := "QIWI_TOKEN"
	// номер карты, на которую мы хотим сделать перевод
	cartNumber := "4377723744084975"
	// сумма перевода
	amount := 10.0

	qw := New(token)

	ctx := context.Background()

	// Определяем провайдера
	providerID, err := qw.Payments.DetectProviderIDByCardNumber(ctx, cartNumber)
	checkErr(err)

	// Проверяем комиссию
	resp, err := qw.Payments.SpecialComission(ctx, providerID, cartNumber, amount)
	checkErr(err)
	log.Printf("Комиссия: %.2f", resp.QwCommission.Amount)
	log.Printf("Сумма платежа с комиссией: %.2f", resp.WithdrawSum.Amount)

	// Делаем платёж
	paymentResponse, err := qw.Payment(ctx, providerID, cartNumber, amount)
	checkErr(err)
	log.Printf("ID транзакции: %s", paymentResponse.Transaction.ID)
```