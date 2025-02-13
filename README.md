# Тестовое задание в компанию Tages

## ТЗ
>Необходимо написать сервис на Golang работающий по gRPC.
>Сервис должен:
>1) Принимать бинарные файлы (изображения) от клиента и сохранять их на жесткий диск.
>2) Иметь возможность просмотра списка всех загруженных файлов в формате:
> Имя файла | Дата создания | Дата обновления
> 3) Отдавать файлы клиенту.
> 4) Ограничивать количество одновременных подключений с клиента:
> - на загрузку/скачивание файлов - 10 конкурентных запросов;
> - на просмотр списка файлов - 100 конкурентных запросов.

## Описание проекта
- gRPC сервис работает на порту 50051 внутри докер сети и пробрасывается на порт 9051 внутри docker-compose.yml
- переменные окружения задаются в ```.env```
- конфиги в ```storage/config/```
- точка входа находится в ```storage/cmd/main.go```
- обработка запросов разделена не сколько слоев:
  - **controller** - слой контроллера ```storage/internal/controller/grpc/```
  - **service** - слой бизнес логики ```storage/internal/service/```
  - **repository** - слой репозитория для работы с бд ```storage/internal/repository/```
- структуры в ```storage/internal/entity/```
- соединение с postgres в ```storage/internal/db/postgres/```
- миграции в ```storage/tools/migrations/```
- **.proto** и сгенерированные файлы находятся в ```storage/pkg/proto/storage/```



## Запуск проекта
- Если установлен Make
  
  ```
  make run-docker
  ```
- Если нет Make
  
  ```
  docker compose --env-file .env -f docker-compose.yml up
  ```

## Телеграм для обратной связи
[https://t.me/daniilnaumovets](https://t.me/daniilnaumovets)




