ser:
	go run main.go

network:
	docker network create 200lab-network

mysql:
	docker run --name 200lab-mysql  --network 200lab-network -e MYSQL_ROOT_PASSWORD=12345678 -e MYSQL_USER="food_delivery" -e MYSQL_PASSWORD=12345678 -e MYSQL_DATABASE="food_delivery" -p 3307:3306 -d mysql:latest

createdb:
	docker exec -it 200lab-mysql createdb --username=root --owner=root 200lab_restaurant