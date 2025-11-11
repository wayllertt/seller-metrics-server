# seller-metrics-server
подсчет метрики продавцов на основе заказов

ручки:
POST /orders/demo
GET /sellers/metrics
GET /metrics/seller?seller_id=1

запуск тестов:
go test ./...

запуск:
go run ./cmd/api
