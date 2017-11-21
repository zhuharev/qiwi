// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package qiwi

import "fmt"

var (
	// ErrSyntaxError error with bad request
	ErrSyntaxError = fmt.Errorf("Ошибка синтаксиса запроса (неправильный формат данных)")
	// ErrTokenInvalid error
	ErrTokenInvalid = fmt.Errorf("Неверный токен или истек срок действия токена")
	// ErrNotFound error
	ErrNotFound = fmt.Errorf("Не найден кошелёк, транзакция или отсутствуют платежи с указанными признаками")
	// ErrRateLimitReached error
	ErrRateLimitReached = fmt.Errorf("Слишком много запросов, сервис временно недоступен")
)

var (
	codeToError = map[int]error{
		400: ErrSyntaxError,
		401: ErrTokenInvalid,
		404: ErrNotFound,
		423: ErrRateLimitReached,
	}
)
