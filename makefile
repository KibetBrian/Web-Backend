database := electiondb
dbName := election

initdb:
	docker run --name $(database) -d -e POSTGRES_USER=brian -e POSTGRES_PASSWORD=brian -p 5432:5432 postgres
startdb:
	docker start $(database)
stopdb:
	docker stop $(database)
deldb:
	docker stop $(database) && docker rm $(database)
createdb:
	docker exec -it ${database} createdb $(dbName) --username=brian --owner=brian;
dropdb:
	docker exec -it backend dropdb $(dbName)--username=brian
