mysqlService:
	docker run --name mysql-commercial -p 3306:3306 -e MYSQL_ROOT_PASSWORD=verysecretpass -e MYSQL_DATABASE=order -d mysql

orderService:
	DATA_SOURCE_URL=root:verysecretpass@tcp\(localhost:3306\)/order APPLICATION_PORT=3000 PAYMENT_SERVICE_URL=localhost:3001 ENV=development go run cmd/main.go

.PHONY: mysqlService orderService