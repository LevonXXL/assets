### Установка goos (для миграции, если требуется)

```bash 
make init 
```

### Мигации хранятся в ./migrations
#### миграция +1
```bash 
PG_DB_DSN='postgresql://user:password@host:port/assets_db' make migration up 
```

#### миграция -1
```bash 
PG_DB_DSN='postgresql://user:password@host:port/assets_db' make migration down 
```

#### Сборка
```bash 
make build
```

### Переменные окружения (enviroment):
```bash 
APP_PORT=10101;PG_DB_DSN=postgresql://asset:asset-password@127.0.0.1:5432/asset_db
SESSION_DURATION_HOURS=24
USE_HTTPS=true //если не true, то будет http
```


### Пути:
/api/auth - авторизация (POST)

/api/upload-asset/{asset_name} - загрузка файла "asset_name" (POST)

/api/asset/{asset_name} - получение данных из файла "asset_name" (GET)

/api/delete-asset/{asset_name} - удаление файла asset_name (DELETE)

/api/assets - список файлов (доступны без авторизации) без данных с данными pagination, доступны query параметры: per_page (100 по умолчанию) и page (1 по умолчанию)

### HTTPS
В директории /certs лежит самоподписаный сертификат. 
Для работы по https нужно установить server.crt в доверенные сертификаты системы

в переменных окружения установить
```bash 
USE_HTTPS=true
```
и добавить в /etc/hosts строку

127.0.0.1  aterentev.lc
