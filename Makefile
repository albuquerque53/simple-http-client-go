up:
	docker-compose up -d
down:
	docker-compose down
app:
	docker exec -it go_app bash
run:
	go run app/main.go
