Запуск  ||
        \/
cd src/cmd/server 
 go run main.go
 http запросы 
 curl -X POST http://localhost:8080/task/create создать таску
curl http://localhost:8080/task/info запрос результата работы таски
curl -X DELETE http://localhost:8080/task/delete удалить таску

таска - абстракция какой-то работы. Лицом таски является контекст который живет 5 минут (после завершения контекста выводится статус 200)
