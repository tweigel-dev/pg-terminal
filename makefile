run:
	go run main.go


test-database-start:
	docker rm -f tesla-test-database
	docker run --rm -p 5432:5432 -e POSTGRES_PASSWORD=changeme -e POSTGRES_USER=postgres -e POSTGRES_DB=tesla-test -d  --name tesla-test-database  postgres