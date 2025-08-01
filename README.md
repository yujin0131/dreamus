# Dreamus CI/CD Build Project  by Yujin Lee

## 빌드 방법 및 애플리케이션 실행 방법

1. create a `.env` file in the project root directory

```env
DB_USER=dreamus
DB_PASSWORD=Wlsehf0014@
DB_HOST=mysql
DB_PORT=3306
DB_NAME=dreamus_db

MYSQL_ROOT_PASSWORD=WlsehfRoot123!
MYSQL_DATABASE=dreamus_db
MYSQL_USER=dreamus
MYSQL_PASSWORD=Wlsehf0014@
```

2. Install `Docker` and `Docker Compose`, and start Docker daemon.

3. Run `docker-compose up --build`