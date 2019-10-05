# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2019-10-05

### Изменено

- Перешли на модули
- Убрали зависимости
- Добавлен контекст всем вызовам
- Все методы `Cards` перенесены в `Payments`
- Удалена структура `Cards`
- Добавлена поддержка распознавания ошибки "Не достаточно средств"
- Добавлены тесты и в них есть пример вывода на карту, вывода на Qiwi-кошелек, запрос истории платежей

## [0.0.5] - 2017-11-21
### Исправлено
- Обработка ошибок при запросах к АПИ

## [0.0.4] - 2017-11-21
### Добавлено
- Обработка ошибок

### Изменено
- Теперь данные логируются только если  дебаг включён

## [0.0.3] - 2017-11-12
### Added
- Добавлен опциональный аргумент при создании платежа, который даёт возможность отправлять комментарий вместе с платежом (только для переводов на qiwi-кошельки).

## [0.0.2] - 2017-11-10
### Changed
- History to Payments
- makeRequest now not return io.ReadCloser and at once unmarshal json

## [0.0.1] - 2017-11-09
### Added
- Changelog
- Русификация
