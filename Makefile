.PHONY: all clean

all: onion.go 
	@go generate
	@go build
	sudo ./onion serve

cleandb: Onion.db
	rm -f Onion.db
clean:
	@go clean
	rm -f onion_resource.go
	rm -f web_service.go
	rm -f main.go
	tree
	cat onion.go

migrate:
	./onion migratedb

run:
	./onion serve

test:
	@go test
