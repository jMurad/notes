# Notes
Сервис заметок, с возможностью регистрации и авторизации пользователя, добавления заметки и вывод списка заметок для авторизованного пользователя. 

## Содержание
- [API](#api)
- [Зависимости](#зависимости)
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


## Зависимости
- [Gorilla Sessions](https://github.com/gorilla/sessions/)
- [gorilla/handlers](https://github.com/gorilla/handlers/)
- [gorilla/mux](https://github.com/gorilla/mux/)
- [Logrus](https://github.com/sirupsen/logrus/)
- [BurntSushi/toml](https://github.com/BurntSushi/toml/)
- [pq](https://github.com/lib/pq/)
- [uuid](https://github.com/google/uuid/)

## Сборка и запуск
Настройки проекта хранятся в файле конфигурации: `configs/notes.toml`

Скрипты для создания рабочей и тестовой БД лежат в директории: `tools/`

#### Описание команд в make файле
    build - сборка проекта
    img-build - сборка Docker образа проекта 
    img-del - удаление Docker образа проекта 
    start - запуск postgres и проекта через docker-compose файл
    stop - остановка и удаление запущенных контейнеров
    init - скрипт для создания рабочей БД
    init-test - скрипт для создания тестовой БД
    run-pg - запуск Docker контейнера postgres
    del-pg - остановка и удаление контейнера postgres
    migrate - миграции
    migrate-test - тестовые миграции
    prod - сборка проекта, запуск docker-compose файла, создание БД и запуск миграций
    test - автоматическое выполнение тестовых файлов

### Сборка проекта
```sh
make build
make run-pg
./notes
```

### Production сборка 
```sh
make prod
```

## Тестирование

```sh
make run-pg
make init-test
make migrate-test
make test
```
Шаблоны запросов cURL и Postman коллекция хранятся в директории `test/`