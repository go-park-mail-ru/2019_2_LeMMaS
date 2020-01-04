# LeMMaS Backend [![Build Status](https://travis-ci.org/go-park-mail-ru/2019_2_LeMMaS.svg?branch=master)](https://travis-ci.org/go-park-mail-ru/2019_2_LeMMaS)

Игра про голодные шарики, вдохновленная [Agar.io](http://agar.io/).

[![image](https://user-images.githubusercontent.com/6276455/69713801-3a1aa980-1116-11ea-82db-902277aefbe3.png)](https://lemmas.ru/)

### Архитектура

Проект следует принципам [чистой архитектуры](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).  Используются PostgreSQL, Redis, Sentry, Prometheus, Grafana. 
Микросервисы API, Game, User и Auth общаются с помощью gRPC.

<img width="450" alt="architecture" src="https://user-images.githubusercontent.com/6276455/71639740-4520a680-2c8d-11ea-9b34-3c6910806d6d.png">

### Тесты

Все основные функции приложения протестированы (покрытие 40%).

### Разработка

Запуск: `./docker/bin/start` Документация по API [тут](https://go-park-mail-ru.github.io/2019_2_LeMMaS).

## Авторы

- [Сударев Максим](https://github.com/smi97)
- [Можевикина Леонарда](https://github.com/ledka17)
- [Кобзев Антон](https://github.com/kzon)

## Менторы

-   [Елизавета Щербакова](https://github.com/Liza-Shch)
-   [Дмитрий Палий](https://github.com/stanf0rd)

### [Frontend репозиторий](https://github.com/frontend-park-mail-ru/2019_2_LeMMaS)
