<p align="center">
  <img src="go-linkshort-logo-scaled.png" />
  <br />
  Kudos to https://gopherize.me/ for the cute logo
</p>

## LinkShort

Video walkthrough: https://www.youtube.com/watch?v=XQpFS3cSMRE

Based on supernova https://github.com/alexgtn/supernova 

Implemented with hexagonal architecture. More: https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/

```
go build -o linkshort

Usage:
  linkshort [command]

Available Commands:
  completion         Generate the autocompletion script for the specified shell
  execute-migration  Runs migration against a local db
  generate-migration Generate migration from schema
  http               gRPC HTTP gateway
  main               gRPC server
```

Run service
```
go run main.go main
go run main.go http
go run main.go execute-migration
```

Run tests with coverage

```
go test ./... -cover -test.short
```

Run fuzz tests

```
go test github.com/alexgtn/go-linkshort/usecase -fuzz=FuzzService_Create
```

### System-level tests

Make sure server is running

```
go run main.go main
go run main.go http
go run main.go execute-migration
```

Run burst tests. Launches 100x requests (create, redirect, create & redirect) concurrently.
```
go test -v -run Test_Main
```
