# bavovna-lint

opinionated golang linters

## Linters

- `pkg/analyzers/appendr` - `append` statements usage linter
- `pkg/analyzers/elser` - `else` statements usage linter
- `pkg/analyzers/readall` - `ioutil.ReadAll` usage linter

# ioutil.ReadAll linter only

```
go get -u github.com/gopkgz/bavovna-lint/cmd/bavovna-lint
go vet -vettool ~/bin/bavovna-lint
```

# All linters

```
go get -u github.com/gopkgz/bavovna-lint/cmd/bavovna-lint-all
go vet -vettool ~/bin/bavovna-lint-all ./...
```

# Skip files/directories

Set `IGNORE` environment variable to comma-separated glob-like file patterns.

Examples:
- `IGNORE="**/*_test.go"` ignores all the test files
- `IGNORE="**/cmd/**"` ignores everything under cmd directory

```
IGNORE="**/*_test.go" go vet -vettool ~/bin/bavovna-lint-all ./...
```

NOTE: uses [github.com/gobwas/glob](https://github.com/gobwas/glob) with `'/'` separator.

## Building and running from source

```
make build
go vet -vettool ./build/bin/bavovna-lint-all ./...
```

## Development

```
make ci      # vet + fmt + lint + self-lint + test
make cover   # coverage HTML at build/coverage.html
make clean   # remove build/
```
