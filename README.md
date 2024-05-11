## Описание

Приложение для хранения и создания заметок по категориям. Приложение состоит из
трех микросервисов.

1) Api Gateway (Python, FastApi) - является точкой входа в приложение, валидируется
запросы, проверяет авторизацию и аутентификацию пользователей (парсит JWT). Через Api Gateway
происходит обращение к другим сервисам.

2) Auth service (Go) - сервис для регстрации и авторизации пользователей. Авторизация
реализована через JWT, токены хранятся в Redis. 

3) Note service (Python, gRPC) - сервис для хранения и создания заметок/категорий. 
Для использования требуется регистрация/авторизация в сервисе Auth service.

Для всех приложений созданые Dockerfile.

Разумеется, для такого простого приложения не требуется использования микросервисной архитектуры, gRPC, разных языков программирования и прочих усложнений.
Вполне достаточно использовать даже простой Django с его html шаблонами.
ОДНАКО очень-очень хотелось попробовать написать микросервисы (впервые в жизни), да еще и на разных языках. Поэтому и написал.

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

#### HTTP Methods

Набор эндпоинтов см. выше (Api Gateway)

### Note Service (Python, gRPC, Postgres)

#### Описание

Сервис для создания и хранения заметок по категориям. Взаимодействие происходит по gRPC
через Api Gateway.

#### Схема БД

![](https://github.com/iriskin77/notes_microservices/blob/master/images/endpoints.png)

#### gRPC methods:

Реализованные gRPC методы:

Можно посмотреть выше (Api Gateway) или здесь в виде protobuf:

```
service Note {
    rpc CreateNote (CreateNoteRequest) returns (CreateNoteResponse);
    rpc GetNote (GetNoteRequest) returns (GetNoteResponse);
    rpc GetNotes (GetListNotesRequest) returns (GetListNotesResponse);
    rpc CreateCategory (CreateCategoryRequest) returns (CreateCategoryResponse);
    rpc GetNotesByCategory (GetNotesByCategoryRequest) returns (GetNotesByCategoryResponse);
    rpc UpdateNote (UpdateNoteRequest) returns (UpdateNoteResponse);
    rpc DeleteNote (DeleteNoteRequest) returns (DeleteNoteReponse);
}
```


