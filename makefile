dbName := election

#app
tidy:
	go mod tidy
run:
	go run main.go
test:
	go test ./... -v -cover

initdb:
	docker run --name postgres -d -e POSTGRES_USER=brian -e POSTGRES_PASSWORD=brian -p 5432:5432 postgres
startdb:
	docker start postgres
stopdb:
	docker stop postgres
deldb:
	docker stop postgres && docker rm postgres
createdb:
	docker exec -it postgres createdb ${dbName} --username=brian --owner=brian;
dropdb:
	docker exec -it backend dropdb ${dbName}--username=brian
migrateUp: 
	migrate -path ./db/migrations -database "postgresql://brian:brian@127.0.0.1:5432/election?sslmode=disable" -verbose up
migrateDown:
	 migrate -path ./db/migrations -database "postgresql://brian:brian@127.0.0.1:5432/election?sslmode=disable" -verbose down
