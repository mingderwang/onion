.PHONY: all clean

all: onion.go 
	@go build
	./onion serve

cleandb: Onion.db
	rm -f Onion.db
clean:
	@go clean

migrate:
	./onion migratedb

run:
	./onion serve

test:
	@go test
