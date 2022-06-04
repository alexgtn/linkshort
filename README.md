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