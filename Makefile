createdb:
	docker run -p 3307:3306 --name db -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=bank -e MYSQL_USER=banker -e MYSQL_PASSWORD=banker123 -d mysql:latest

linkadminer:
	docker run --name adminer_db --link db:db -p 8086:8080 -d adminer 

mysql:
	docker exec -it db mysql -u root -proot

migrateup:
	migrate -path db/migration -database "mysql://root:root@tcp(localhost:3307)/bank" -verbose up


migratedown:
	migrate -path db/migration -database "mysql://root:root@tcp(localhost:3307)/bank" -verbose down

migratedown1:
	migrate -path db/migration -database "mysql://root:root@tcp(localhost:3307)/bank" -verbose down 1

migrateup1:
	migrate -path db/migration -database "mysql://root:root@tcp(localhost:3307)/bank" -verbose up 1

remove:
	docker stop db
	docker stop adminer_db
	docker rm db
	docker rm adminer_db
server:
	go run main.go
.PHONY:
	createdb
	linkadminer
	migrateup
	migratedown
	migratedown1
	migrateup1
	server
	remove