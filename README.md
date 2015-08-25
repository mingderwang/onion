# onion

#Pre-install (on Ubuntu 14.x for example)

You need to pre-install golang, git, and make from respective package.

##install golang:
refer to https://golang.org/dl/

##install git:
```
sudo apt-get update
sudo apt-get install git
```

##install make: 
```
sudo apt-get install build-essential
```

#Build
```
go get ./...
go get github.com/mingderwang/ginger
go generate
go build
```

#Migrate database
```
./onion migratedb
```

#Run service
```
./onion serve
```


#Clean
to clean auto-generated file
```
make clean
```

#Welcome to contribute and share

Licensed by MIT License
Copyright 2015 Ming-der Wang <ming@log4analytics.com>
