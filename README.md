# polygot

This API should accept a number and API response should be 'true' if the given number is a Two-sided Prime (https://prime-numbers.info/article/two-sided-primes), otherwise 'false'.


## Installation guide
```sh
$ git clone https://github.com/sanketdhoble/polygot-go.git
$ cd polygot-go
```

## Build
The following command will retrieve the dependency package

```
go get
```

## Run
The following command will start API server at port 8080

```
go run prime-checker.go

Url to hit:  

`http://localhost:8080/isTwoSidedPrime/{number}`

