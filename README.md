# go_webcalc

## Описание

**Сервис go_webcalc разработан на языке программирования Go и предназначен для выполнения арифметических операций через HTTP-запросы. Он обрабатывает строки с математическими выражениями, выполняет вычисления и возвращает результаты. Этот проект может быть использован как самостоятельное API или интегрирован с другими сервисами для удобного выполнения математических операций.

## Функции
Возможности проекта включают в себя обработку арифметических выражений с операциями сложения, вычитания, умножения и деления, JSON API для взаимодействия через POST-запросы, обработку ошибок, включая деление на ноль, и наличие готовых тестов для проверки корректности работы.

### Требования
Для работы необходима установленная версия Go 1.18 и выше, а также наличие Git. Также нужен доступ в интернет для загрузки необходимых зависимостей.
### Установка
Склонируйте репозиторий командой:
```bash

git clone https://github.com/jKakZyo/go_webcalc.git

cd go_webcalc


Установите необходимые зависимости:

go mod tidy


Для запуска в режиме разработки используйте команду:

go run cmd/go_webcalc/main.go

Сервер будет доступен по адресу: http://localhost:8080/api/v1/calculate.

Для запуска в режиме пользователя соберите бинарный файл:

go build -o go_webcalc cmd/go_webcalc/main.go

Запустите файл:

./go_webcalc


API позволяет отправлять POST запросы на адрес /calculate, принимая математическое выражение и возвращая результат.

Пример запроса:

{ "expression": "2 - 5 * (3 + 9)" }


Пример ответа:

{ "result": -58 }


При некорректном выражении сервер вернет ошибку HTTP 400 Internal Server Error.

Для запуска тестов используйте:

go test -v ./...


Пример использования API с помощью curl:

curl -X POST http://localhost:8080/calculate \

-H "Content-Type: application/json" \

-d '{"expression": "(2 + 5) * 10"}'


Ответ:

{ "result": 70 }
