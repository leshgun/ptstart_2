## CI/CD
Сборка HTTP-сервера (Rest API) на Python 3.12 в Docker-контейнере

## Запуск

Перед запуском необходимо выставить *ip*-адрес и *port* в файле `Dockerfile`, на которых будет работать сервис Rest API.

Для запуска сервиса:
```bash
docker build -t cicd . && docker run -dit cicd
```

Для запуска в интерактивном режиме (с журналом запросов)
```bash
docker build -t cicd . && docker run -it cicd
```