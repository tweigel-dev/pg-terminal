test-database-start:
	docker rm -f test-database
	docker run --rm -p 5432:5432 -e POSTGRES_PASSWORD=changeme -e POSTGRES_USER=postgres -e POSTGRES_DB=test -d  --name test-database  postgres