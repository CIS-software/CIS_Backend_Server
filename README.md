# Приложение для секции карате (Backend)
## Назначение проекта
Это некая социальная сеть для участников секции, в ней они могут смотреть новости, связанные с секцией, и расписание тренировок.

## Установка проекта
1) Скачать проект и докер
2) Собрать все сервисы, описанные в docker-compose. Команда в терминале `docker compose build`
3) Запустить все сервисы, описанные в docker-compose. Команда в терминале `docker compose up -d`
4) Выполнить миграцию таблицы в БД. Путь к файлу с таблицей: `migrations/migratinos-up.sql`
5) Создать bucket для Minio по адресу http://127.0.0.1:9001 с именем указаным в `config/config.go`

**P.S** Все данные(логины, пароли, адреса, имена и др.) есть в [config/config.go](https://github.com/CIS-software/CIS_Backend_Server/blob/main/config/config.go)

## API
Документация: https://cis-software.stoplight.io/docs/cis-backend-server/49e6161557049-users

Postman: https://www.postman.com/technical-cosmologist-73620193/workspace/cis-software-cis-backend-server
