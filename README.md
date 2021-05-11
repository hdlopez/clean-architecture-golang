![Go](https://github.com/hdlopez/clean-architecture-golang/workflows/Go/badge.svg)

# Clean Architecture in Go

This is my humble version of how to structure a GO application. 

Link to [my article](https://dev.to/hdlopez/como-estructurar-tu-aplicacion-en-go-4hk2) (currently only in spanish) about "How to structure a Go Application".

In order to demonstrate my idea with a simple example application, I created this API which handles messages. 

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
* Use `Context` type to pass parameters across different layers
* Add an example of database handler module

