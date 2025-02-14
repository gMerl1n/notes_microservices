## Описание

Приложение для хранения и создания заметок по категориям. Приложение состоит из трех микросервисов.

1) Api Gateway (Go, http) - является точкой входа в приложение, в этом сервисе маршрутизируются
запросы, проверяется авторизация и аутентификация пользователей (парсит JWT). Через Api Gateway
происходит обращение к другим сервисам.

| Технология       | Версия  | 
|------------------|---------|
| Go               | 1.23.4  | 
| gorilla/mux      | 1.8.1   | 
| sirupsen/logrus  | 1.9.3   |

2) Auth service (Go, REST API, JWT, Postgres, Redis) - REST API сервис для регистрации и авторизации пользователей. Авторизация
реализована через JWT, токены хранятся в Redis, данные о пользователях в PostgreSQL.

| Технология       | Версия  | 
|------------------|---------|
| Go               | 1.23.4  | 
| gorilla/mux      | 1.8.1   | 
| jackc/pgx/v5     | 5.7.2   |
| go-redis/v9      | 9.7.0   |
| sirupsen/logrus  | 1.9.3   |

3) Note service (Python, FastApi) - REST API сервис для хранения и создания заметок и категорий. 
Для использования требуется авторизация в сервисе Auth service.

| Технология | Версия  | 
|------------|---------|
| Python     | 3.12    | 
| FastApi    | 1.8.1   | 
| Pydantic   | 2.7.1   |
| SQLAlchemy | 2.0.30  |
| Alembic    | 0.115.7 |
| Punq       | 0.7.0   |

Для всех приложений созданы Dockerfile.

#### Схема микросервисов

![](https://github.com/iriskin77/notes_microservices/blob/master/images/microservices.png)


#### Общая схема БД

![](https://github.com/iriskin77/notes_microservices/blob/master/images/db_schema.png)


## Как запустить 

Все три сервиса упакованы в Docker, поэтому:

+ Запуск api-gateway: 
  + Перейти в папку 01_api_gateway,
  + Использовать команду docker-compose build
  + Использовать команду docker-compose up

+ Запуск auth-service: 
  + Перейти в папку 02_auth_service,
  + Использовать команду docker-compose build
  + Использовать команду docker-compose up
  + Перейти в созданный контейнер (docker exec -it <container_id> bash)
  + Создать миграции: migrate -path ./app/migrations -database 'postgres://pguser:pgpassword@localhost:5432/postgres?sslmode=disable' up


+ Запуск api-gateway: 
  + Перейти в папку 03_notes_service,
  + Использовать команду docker-compose build
  + Использовать команду docker-compose up