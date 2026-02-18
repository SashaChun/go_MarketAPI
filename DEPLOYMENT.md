# 🚀 Deployment Guide

## AWS EC2 Setup

### 1. Створення EC2 Instance

1. Зайди в AWS Console → EC2
2. Launch Instance:
   - **AMI**: Amazon Linux 2023
   - **Instance Type**: t2.micro (або більше)
   - **Security Group**: відкрий порти 22 (SSH), 8080 (API), 5432 (PostgreSQL)
3. Створи або вибери SSH key pair
4. Launch instance

### 2. Підключення до EC2

```bash
ssh -i your-key.pem ec2-user@your-ec2-ip
```

### 3. Налаштування EC2

Запусти setup скрипт:

```bash
curl -o setup.sh https://raw.githubusercontent.com/YOUR_USERNAME/untitled9/main/scripts/setup-ec2.sh
chmod +x setup.sh
./setup.sh
```

Або вручну:

```bash
# Оновлення системи
sudo yum update -y

# Встановлення Docker
sudo yum install -y docker
sudo systemctl start docker
sudo systemctl enable docker
sudo usermod -aG docker ec2-user

# Встановлення Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Встановлення Git
sudo yum install -y git

# Клонування репозиторію
cd ~
git clone https://github.com/YOUR_USERNAME/untitled9.git
cd untitled9

# Створення .env файлу
cp .env.example .env
nano .env  # Відредагуй значення
```

### 4. GitHub Secrets

Додай в GitHub репозиторій (Settings → Secrets and variables → Actions):

- **EC2_SSH_KEY**: Вміст твого `.pem` файлу
- **EC2_HOST**: IP адреса твого EC2 instance
- **EC2_USER**: `ec2-user` (для Amazon Linux)

### 5. Перший Deploy

```bash
cd ~/untitled9
docker-compose -f docker-compose.prod.yml up -d --build
```

Перевір статус:

```bash
docker-compose -f docker-compose.prod.yml ps
docker-compose -f docker-compose.prod.yml logs -f
```

### 6. Автоматичний Deploy

Після push в `main` branch, GitHub Actions автоматично:
1. Запустить тести
2. Підключиться до EC2
3. Зробить pull останніх змін
4. Перезапустить контейнери
5. Виконає health check

## Корисні команди

### Логи
```bash
docker-compose -f docker-compose.prod.yml logs -f app
docker-compose -f docker-compose.prod.yml logs -f postgres
```

### Перезапуск
```bash
docker-compose -f docker-compose.prod.yml restart app
```

### Зупинка
```bash
docker-compose -f docker-compose.prod.yml down
```

### Повне очищення
```bash
docker-compose -f docker-compose.prod.yml down -v
docker system prune -af
```

### Backup БД
```bash
docker exec shop-postgres pg_dump -U shop shopdb > backup.sql
```

### Restore БД
```bash
cat backup.sql | docker exec -i shop-postgres psql -U shop -d shopdb
```

## API Endpoints

Після деплою API буде доступне на:

```
http://YOUR_EC2_IP:8080
```

### Auth
- `POST /auth/register` - Реєстрація
- `POST /auth/login` - Логін

### Products
- `POST /products` - Створити продукт
- `GET /products` - Список продуктів
- `GET /products/:id` - Отримати продукт
- `PUT /products/:id` - Оновити продукт
- `DELETE /products/:id` - Видалити продукт

## Troubleshooting

### Перевірка портів
```bash
sudo netstat -tulpn | grep :8080
```

### Перевірка Docker
```bash
docker ps
docker logs shop-api
```

### Перевірка БД
```bash
docker exec -it shop-postgres psql -U shop -d shopdb
```

### Firewall (якщо потрібно)
```bash
sudo firewall-cmd --permanent --add-port=8080/tcp
sudo firewall-cmd --reload
```

## Security Checklist

- [ ] Змінити `JWT_SECRET` на випадковий рядок
- [ ] Використовувати сильний пароль для PostgreSQL
- [ ] Налаштувати HTTPS (Let's Encrypt + Nginx)
- [ ] Обмежити доступ до PostgreSQL (тільки з localhost)
- [ ] Регулярно оновлювати Docker images
- [ ] Налаштувати автоматичні backup'и БД
- [ ] Використовувати AWS Secrets Manager для чутливих даних
