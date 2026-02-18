# Shop API

REST API для інтернет-магазину на Go + Gin + GORM + PostgreSQL з автоматичним деплоєм на AWS EC2.

## Технології

- **Go 1.25**
- **Gin** - HTTP framework
- **GORM** - ORM
- **PostgreSQL 16** - База даних
- **JWT** - Аутентифікація
- **Docker** - Контейнеризація
- **GitHub Actions** - CI/CD
- **AWS EC2** - Хостинг

## Швидкий старт

### Локальна розробка

1. Клонуй репозиторій:
```bash
git clone https://github.com/YOUR_USERNAME/untitled9.git
cd untitled9
```

2. Створи `.env` файл:
```bash
cp .env.example .env
```

3. Запусти через Docker Compose:
```bash
docker-compose up -d
```

4. API доступне на `http://localhost:8080`

### Без Docker

1. Встанови PostgreSQL
2. Створи базу даних `shopdb`
3. Налаштуй `.env`
4. Запусти:
```bash
go mod download
go run cmd/api/main.go
```

## API Endpoints

### Health Check
```
GET /health
```

### Authentication
```
POST /auth/register
POST /auth/login
```

### Products (CRUD)
```
POST   /products
GET    /products
GET    /products/:id
PUT    /products/:id
DELETE /products/:id
```

## Deployment

Детальна інструкція по деплою на AWS EC2: [DEPLOYMENT.md](DEPLOYMENT.md)

### Швидкий деплой

1. Створи EC2 instance
2. Налаштуй GitHub Secrets:
   - `EC2_SSH_KEY`
   - `EC2_HOST`
   - `EC2_USER`
3. Push в `main` branch - автоматичний деплой через GitHub Actions

## Структура проєкту

```
.
├── cmd/
│   └── api/
│       └── main.go              # Entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Конфігурація
│   ├── db/
│   │   └── db.go                # Підключення до БД
│   └── http/
│       ├── controler/
│       │   ├── auth/            # Auth handlers
│       │   ├── products/        # Product handlers
│       │   └── health.go        # Health check
│       └── model/
│           ├── user_dto.go      # User models
│           └── product.go       # Product model
├── scripts/
│   ├── setup-ec2.sh             # EC2 setup script
│   └── deploy.sh                # Manual deploy script
├── .github/
│   └── workflows/
│       └── deploy.yml           # CI/CD pipeline
├── docker-compose.yml           # Dev environment
├── docker-compose.prod.yml      # Production
├── Dockerfile
└── .env.example
```

## Розробка

### Запуск тестів
```bash
go test -v ./...
```

### Логи Docker
```bash
docker-compose logs -f
```

### Міграції
Автоматично через GORM AutoMigrate при старті.

## License

MIT
