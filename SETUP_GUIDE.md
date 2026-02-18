# 🚀 Покрокова інструкція налаштування CI/CD

## Крок 1: Створення AWS EC2 Instance

### 1.1 Зайди в AWS Console
- Перейди на https://aws.amazon.com/console/
- Увійди в свій акаунт (або створи новий)

### 1.2 Запусти EC2 Instance
1. В пошуку знайди **EC2** → натисни **Launch Instance**
2. Налаштуй instance:

**Name**: `shop-api-server` (або будь-яка назва)

**Application and OS Images (AMI)**:
- Вибери **Amazon Linux 2023** (безкоштовно)

**Instance type**:
- Вибери **t2.micro** (Free tier eligible)

**Key pair (login)**:
- Натисни **Create new key pair**
- Name: `shop-api-key`
- Key pair type: **RSA**
- Private key file format: **.pem**
- Натисни **Create key pair** → файл `.pem` завантажиться

**Network settings**:
- Натисни **Edit**
- Створи Security Group з такими правилами:

| Type | Protocol | Port Range | Source |
|------|----------|------------|--------|
| SSH | TCP | 22 | 0.0.0.0/0 |
| Custom TCP | TCP | 8080 | 0.0.0.0/0 |
| PostgreSQL | TCP | 5432 | 0.0.0.0/0 |

**Configure storage**:
- 8 GB (за замовчуванням, можна збільшити)

3. Натисни **Launch instance**
4. Дочекайся статусу **Running**
5. **Скопіюй Public IPv4 address** (наприклад `54.123.45.67`)

---

## Крок 2: Налаштування GitHub Secrets

### 2.1 Підготуй SSH ключ
1. Знайди завантажений файл `shop-api-key.pem`
2. Відкрий його в текстовому редакторі
3. Скопіюй **весь вміст** (включно з `-----BEGIN RSA PRIVATE KEY-----` і `-----END RSA PRIVATE KEY-----`)

### 2.2 Додай Secrets в GitHub
1. Перейди в свій репозиторій: https://github.com/SashaChun/go_MarketAPI
2. Натисни **Settings** (вгорі справа)
3. В лівому меню: **Secrets and variables** → **Actions**
4. Натисни **New repository secret**

Створи **3 секрети**:

#### Secret 1: EC2_SSH_KEY
- Name: `EC2_SSH_KEY`
- Secret: **вставити весь вміст `.pem` файлу**
- Натисни **Add secret**

#### Secret 2: EC2_HOST
- Name: `EC2_HOST`
- Secret: **IP адреса EC2** (наприклад `54.123.45.67`)
- Натисни **Add secret**

#### Secret 3: EC2_USER
- Name: `EC2_USER`
- Secret: `ec2-user`
- Натисни **Add secret**

---

## Крок 3: Підключення до EC2 та налаштування

### 3.1 Підключись через SSH

**На Mac/Linux**:
```bash
chmod 400 ~/Downloads/shop-api-key.pem
ssh -i ~/Downloads/shop-api-key.pem ec2-user@54.123.45.67
```

**На Windows** (PowerShell):
```powershell
ssh -i C:\Users\YourName\Downloads\shop-api-key.pem ec2-user@54.123.45.67
```

Якщо запитає "Are you sure you want to continue connecting?" → напиши `yes`

### 3.2 Запусти setup скрипт

Після підключення до EC2 виконай:

```bash
# Завантаж setup скрипт
curl -o setup.sh https://raw.githubusercontent.com/SashaChun/go_MarketAPI/main/scripts/setup-ec2.sh

# Зроби його виконуваним
chmod +x setup.sh

# Запусти
./setup.sh
```

Скрипт автоматично встановить:
- Docker
- Docker Compose
- Git
- Склонує твій репозиторій
- Створить `.env` файл

### 3.3 Налаштуй .env файл

```bash
cd ~/go_MarketAPI
nano .env
```

Відредагуй значення (особливо `JWT_SECRET` і `PASSWORD`):

```env
PORT=5432
HOST=postgres
DATABASE_URL=shopdb
DB_USER=shop
PASSWORD=your-strong-password-here
JWT_SECRET=your-super-secret-jwt-key-min-32-chars
```

Збережи: `Ctrl+O` → `Enter` → `Ctrl+X`

### 3.4 Перший запуск

```bash
docker-compose -f docker-compose.prod.yml up -d --build
```

Перевір статус:
```bash
docker-compose -f docker-compose.prod.yml ps
docker-compose -f docker-compose.prod.yml logs -f
```

Якщо все ОК, натисни `Ctrl+C` щоб вийти з логів.

---

## Крок 4: Тестування API

Перевір, що API працює:

```bash
# Health check
curl http://54.123.45.67:8080/health

# Має повернути:
# {"database":"connected","status":"healthy"}
```

Тестуй з локальної машини:

```bash
# Реєстрація
curl -X POST http://54.123.45.67:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","email":"test@example.com","password":"123456"}'

# Створення продукту
curl -X POST http://54.123.45.67:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop","description":"Gaming laptop","price":1500.99,"stock":10}'

# Список продуктів
curl http://54.123.45.67:8080/products
```

---

## Крок 5: Автоматичний деплой

Тепер все готово! Коли ти робиш:

```bash
git add .
git commit -m "Update API"
git push origin main
```

GitHub Actions **автоматично**:
1. ✅ Запустить тести
2. ✅ Підключиться до EC2 через SSH
3. ✅ Зробить `git pull`
4. ✅ Перезапустить Docker контейнери
5. ✅ Виконає health check

Перевірити статус деплою можна тут:
https://github.com/SashaChun/go_MarketAPI/actions

---

## Корисні команди на EC2

### Перегляд логів
```bash
cd ~/go_MarketAPI
docker-compose -f docker-compose.prod.yml logs -f app
docker-compose -f docker-compose.prod.yml logs -f postgres
```

### Перезапуск
```bash
docker-compose -f docker-compose.prod.yml restart app
```

### Зупинка всього
```bash
docker-compose -f docker-compose.prod.yml down
```

### Повне очищення (видалить БД!)
```bash
docker-compose -f docker-compose.prod.yml down -v
docker system prune -af
```

### Backup БД
```bash
docker exec shop-postgres pg_dump -U shop shopdb > backup_$(date +%Y%m%d).sql
```

### Вихід з EC2
```bash
exit
```

---

## Troubleshooting

### Помилка "Connection refused"
Перевір Security Group в AWS Console → EC2 → Security Groups

### Помилка "Permission denied (publickey)"
```bash
chmod 400 shop-api-key.pem
```

### Docker не запускається
```bash
sudo systemctl start docker
sudo systemctl enable docker
```

### Порт 8080 зайнятий
```bash
sudo netstat -tulpn | grep :8080
# Або змінити PORT в .env
```

---

## Що далі?

- [ ] Налаштувати HTTPS (Let's Encrypt + Nginx)
- [ ] Додати моніторинг (Prometheus + Grafana)
- [ ] Налаштувати автоматичні backup'и БД
- [ ] Додати rate limiting
- [ ] Налаштувати CDN для статики

Готово! Твій API тепер на production з автоматичним деплоєм! 🚀
