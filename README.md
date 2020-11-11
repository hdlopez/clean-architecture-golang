# Clean Archiecture in Go

This is my humble version of how to structure an application in Golang. 

In order to demostrate my idea with an almost real example, I created a API which handles messages. 

## Frameworks

The example uses the following frameworks:

* [Gin-Gonic](https://github.com/gin-gonic/gin) to make the API layer
* [Resty](https://github.com/go-resty/resty) to make API calls

## Run

Clone the repository and execute the following command on the root folder

``` bash
go run main.go
```

## Test

``` bash
go test ./...
```

## Backlog

> This is a work in progress. There's still a lot of work to do. PRs are welcome :)
* Add (a LOT) of unit tests
* Add CI integration
* Use `Context` type to pass parameters across different layers
* Add an example of database handler module

