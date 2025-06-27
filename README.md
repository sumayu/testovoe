Запуск  ||
        \/
cd src/cmd/server 
 go run main.go
 http запросы 
 curl -X POST http://localhost:8080/task/create создать 
curl http://localhost:8080/task/info инфо
curl -X DELETE http://localhost:8080/task/delete удалить
