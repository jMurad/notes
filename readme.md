# Notes
Сервис заметок, с возможностью регистрации и авторизации пользователя, добавления заметки и вывод списка заметок для авторизованного пользователя. 

## Содержание
- [API](#api)
- [Технологии](#технологии)
- [Сборка и запуск](#сборка-и-запуск)
- [Тестирование](#тестирование)

## API

#### /users
* `POST` : Create a new user

#### /sessions
* `POST` : Create a new session

#### /private/notes
* `GET` : Get a notes
* `POST` : Create a new note


## Технологии
- [Golang](https://go.dev/)
- [Gorilla Sessions](https://github.com/gorilla/sessions/)
- [gorilla/handlers](https://github.com/gorilla/handlers/)
- [gorilla/mux](https://github.com/gorilla/mux/)
- [Logrus](https://github.com/sirupsen/logrus/)
- [BurntSushi/toml](https://github.com/BurntSushi/toml/)
- [pq](https://github.com/lib/pq/)
- [uuid](https://github.com/google/uuid/)

## Сборка и запуск
Настройки проекта хранятся в файле конфигурации: `configs/notes.toml`

команды записаны в make файле
```sh
#Сборка проекта
make build

#запуск postgres в контейнере Docker
make run-pg

#запуск проекта
./notes
```


```sh
#Production сборка 
make prod
```

## Тестирование
Шаблоны запросов cURL и Postman коллекция хранятся в директории `tools/`

```sh
#Запустить контейнер postgres
make run-pg

#создать тестовую БД
make init-test

#запустить тестовые миграции
make migrate-test

#прогнать тесты
make test
```
