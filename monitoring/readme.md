## Monitoring
Мониторинг доступности сайта "https://ptsecurity.com" при помощи Prometheus и Blackbox-exporter.  
Примеры конфигураций можно найти в файлах:
- `blackbox.yml` - пример стандартных GET- и ICMP-запросов
- `prometheus.yml` - сбор метрик из Blackbox-exporter, описывающие адрес "https://ptsecurity.com"

## Запуск
Данный пример сбора метрик из Blackbox-exportet можно запустить в докер контейнере через команду `docker-compose`:
```bash
docker-compose up --build
```