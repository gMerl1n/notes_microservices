## Описание

Приложение для хранения и создания заметок по категориям. Приложение состоит из
трех микросервисов.

1) Api Gateway (Go) - является точкой входа в приложение, валидируется
запросы, проверяет авторизацию и аутентификацию пользователей (парсит JWT). Через Api Gateway
происходит обращение к другим сервисам.

2) Auth service (Go) - сервис для регстрации и авторизации пользователей. Авторизация
реализована через JWT, токены хранятся в Redis. 

3) Note service (Python, FastApi) - сервис для хранения и создания заметок/категорий. 
Для использования требуется регистрация/авторизация в сервисе Auth service.

Для всех приложений созданые Dockerfile.


#### Схема микросервисов

![](https://github.com/iriskin77/notes_microservices/blob/master/images/microservices.png)


#### Общая схема БД

![](https://github.com/iriskin77/notes_microservices/blob/master/images/db_schema.png)


### Api Gateway (Python, FastApi)

#### Схема swagger из ApiGateway:

![](https://github.com/iriskin77/notes_microservices/blob/master/images/endpoints.png)

### Auth service (Go, REST Api, Postgres)

#### Описание

Сервис для регистрации и авторизации пользователей. Авторизация реализована через JWT. После авторизации пользователю выдаются 
access-token и refresh-token. Сессии с токенами хранятся в Redis.

Взаимодействие с сервисом происходит по HTTP через Api Gateway.

#### Схема БД

![](https://github.com/iriskin77/notes_microservices/blob/master/images/users_db.png)

#### HTTP Methods

Набор эндпоинтов см. выше (Api Gateway)

### Note Service (Python, FastApi, Postgres)

#### Описание

Сервис для создания и хранения заметок по категориям. Взаимодействие происходит по gRPC
через Api Gateway.

#### Схема БД

![](https://github.com/iriskin77/notes_microservices/blob/master/images/notes_db.png)


