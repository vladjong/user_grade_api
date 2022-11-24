# Use_Grade_API

## Описание

Сервис, который позволяет работать с структурой пользователя

Порт `:3333` - метод get для вывода модели пользователя
Порт `:1111` - метод set для добавления пользователя, пользователь можно добавлять частями

## Стек

- `Go`
- `Docker`
- Фреймворк [Gin](https://github.com/gin-gonic/gin)
- Брокер сообщения [Kafka Sarama](https://github.com/Shopify/sarama)
- Конфигурация приложения [viper](https://github.com/spf13/viper)
- Логер [logrus](https://github.com/sirupsen/logrus)

## Реализованный функционал

- [x] Метод get
- [x] Метод set
- [x] Basic auth (middleware)
- [x] Функционал добавления данных частями в методе set
- [x] Брокер сообщения
- [x] Метод `/backup`

## Запуск

1. Склонировать репозиторий

```
git clone https://github.com/vladjong/user_grade_api
```

2. Включить докер

3. Открыть терминал и набрать:
```
make kafka
```
Запускает докер контейнер kafka

4. Открыть терминал и набрать:
```
make
```
Запускает сервис

5. Проверка на стиль

```
make lint
```

## Тестирование

### Порт `1111`

- `/api` - REST API

#### Post

- `/` Метод добавления пользователя

body:
```
type UserGrade struct {
        UserId        string json:"user_id" validate:"required"
        PostpaidLimit int    json:"postpaid_limit"
        Spp           int    json:"spp"
        ShippingFee   int    json:"shipping_fee"
        ReturnFee     int    json:"return_fee"
    }
```

Curl:
```
curl -X POST -H "Content-Type: application/json" \
-d '{
    "user_id": "2",
    "postpaid_limit": 10,
    "spp": 1,
    "shipping_fee": 12,
    "return_fee": 13
}' \
-u 'admin:qwerty' \
http://localhost:1111/api/
```

### Порт `3333`

- `/api` - REST API

#### GET

- `/:id` Метод для вывода модели пользователя

Curl:
```
curl -X 'GET' \
  'http://localhost:3333/api/2' \
  -H 'accept: application/json'
```

- `/backup` Метод для сохранения данных на локальный компьютер

Curl:
```
curl -X 'GET' \
  'http://localhost:3333/backup' \
  -H 'accept: application/json'
```

### Pkg

#### Async_map

Своя структура для хранения данных. Примитивы синхронизации `mutex`

#### Checker

Пакет, который проверяет структуру пользователя и отдает корректную структуру

#### Kafka `:9092`

Пакет, который позволяет работать с брокер сообщениями
