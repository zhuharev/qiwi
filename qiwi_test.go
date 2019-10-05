// QIWI_TOKEN=xxx go test -v *.go
package qiwi

import (
	"context"
	"os"
	"testing"
)

// Пример платежа на карту
func TestPaymentToCard(t *testing.T) {
	// токен
	token := os.Getenv("QIWI_TOKEN")
	if os.Getenv("QIWI_TOKEN") == "" {
		t.Errorf("please, set QIWI_TOKEN env variable")
		return
	}
	// номер карты, на которую мы хотим сделать перевод
	cartNumber := "4377723744084975"
	// сумма перевода
	amount := 10.0

	qw := New(token)

	ctx := context.Background()

	// Определяем провайдера
	providerID, err := qw.Payments.DetectProviderIDByCardNumber(ctx, cartNumber)
	checkErr(t, err, "Detect provider")

	// Проверяем комиссию
	resp, err := qw.Payments.SpecialComission(ctx, providerID, cartNumber, amount)
	checkErr(t, err, "Fetch comission")
	t.Logf("Комиссия: %.2f", resp.QwCommission.Amount)
	t.Logf("Сумма платежа с комиссией: %.2f", resp.WithdrawSum.Amount)

	// Делаем платёж
	//paymentResponse, err := qw.Payment(ctx, providerID, cartNumber, amount)
	//checkErr(err)
	//t.Logf("ID транзакции: %s", paymentResponse.Transaction.ID)
}

// Приме перевода на QIWI-кошелек
func TestPaymentToQiwi(t *testing.T) {
	// токен
	token := os.Getenv("QIWI_TOKEN")
	if os.Getenv("QIWI_TOKEN") == "" {
		t.Errorf("please, set QIWI_TOKEN env variable")
		return
	}
	// номер телефона, на qiwi-кошелек которого мы хотим сделать перевод
	account := "+79997651151"
	// сумма перевода
	amount := 10.0

	qw := New(token)
	qw.debug = true

	ctx := context.Background()

	// Проверяем комиссию
	resp, err := qw.Payments.SpecialComission(ctx, QiwiProviderID, account, amount)
	checkErr(t, err, "Fetch comission")
	t.Logf("Комиссия: %v", resp)
	t.Logf("Комиссия: %.2f", resp.QwCommission.Amount)
	t.Logf("Сумма платежа с комиссией: %.2f", resp.WithdrawSum.Amount)

	// Делаем платёж
	// paymentResponse, err := qw.Payments.Payment(ctx, QiwiProviderID, amount, account, "привет")
	// checkErr(t, err, "Make payment")
	// t.Logf("ID транзакции: %s", paymentResponse.Transaction.ID)
}

func TestPaymentHistory(t *testing.T) {
	// токен
	token := os.Getenv("QIWI_TOKEN")
	if os.Getenv("QIWI_TOKEN") == "" {
		t.Errorf("please, set QIWI_TOKEN env variable")
		return
	}

	qw := New(token, Wallet("79002287508"))
	qw.debug = true

	ctx := context.Background()

	pr, err := qw.Payments.History(ctx, 10)
	checkErr(t, err, "Make payment")
	for _, tx := range pr.Data {
		t.Logf("Транзакция: номер: %v статус: %v тип: %s получатель/отправитель: %s сумма: %v комиссия: %v", tx.TxnID, tx.Status, tx.Type, tx.Account, tx.Sum.Amount, tx.Commission.Amount)
	}
}

func checkErr(t *testing.T, err error, format string, args ...interface{}) {
	if err != nil {
		t.Errorf(format+": "+err.Error(), args...)
	}
}
