# Subscription Service

REST-сервис для работы с онлайн-подписками пользователей.

## Функциональность

- CRUD по подпискам
- Подсчет общей суммы подписок с фильтрами
- Swagger-документация
- Логирование
- Docker Compose

## Запуск

```bash
docker-compose up --build
```

## Архитектура

```text
subscription-service/
├── cmd/
│   └── app/                      # точка входа приложения: запуск сервера
│       └── main.go
│
├── api/
│   ├── handler/                  # обработчики (handlers) HTTP-запросов
│   │   ├── health.go             # эндпоинт для проверки "живости" сервера (/healthz)
│   │   └── subscription_handler.go  # CRUD + расчёт стоимости подписок
│   └── router.go                 # регистрация всех маршрутов
│
├── pkg/
│   ├── config/                   # загрузка переменных окружения из .env
│   │   └── loader.go
│   │
│   ├── database/                 # подключение к PostgreSQL и миграции
│   │   └── postgres.go
│   │
│   ├── logger/                   # инициализация логгера (zap)
│   │   └── logger.go
│   │
│   └── models/                   # GORM-модель подписки
│       └── subscription.go
│
├── docs/                         # Swagger-документация
│   └── docs.go                   # автосгенерированный код документации
│
├── .env.example                  # пример конфигурации окружения
├── docker-compose.yml            # Docker Compose: app + PostgreSQL
├── Dockerfile                    # инструкция по сборке Go-приложения
├── go.mod                        # модули Go
├── go.sum                        # контрольные суммы зависимостей
└── README.md                     # документация проекта

```