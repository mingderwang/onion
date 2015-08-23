.PHONY: all clean

all: onion.go 
	@go generate
	@go build

clean:
	@go clean
	rm -f *_resource.go
	rm -f main.go
	rm -f web_service.go

cleandb: Onion.db
	rm -f Onion.db

migrate:
	./onion migratedb

run:
	./onion serve

test:
	@go test
