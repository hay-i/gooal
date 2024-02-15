hmr:
	templ generate --watch --proxy="http://localhost:1323" --cmd="go run ."

start:
	go run .

up:
	docker compose up -d

down:
	docker compose down

db:
	docker exec -it chronologger-mongo-1 mongosh --username root --password example
