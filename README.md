<p align="center">
  <img src="go-linkshort-logo-scaled.png" />
  <br />
  Kudos to https://gopherize.me/ for the cute logo
</p>

## LinkShort

Run tests with coverage

```
go test ./... -cover
```

Run fuzz tests

```
go test github.com/alexgtn/go-linkshort/usecase -fuzz=FuzzService_Create
```

### System-level tests

Make sure server is running

```
go run main.go main
```

Run burst tests. Launches 100x requests (create, redirect, create & redirect) concurrently.
```
go test -v -run Test_Main
```