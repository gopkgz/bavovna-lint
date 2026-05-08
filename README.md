# bavovna-lint

opinionated golang linters

## Linters

- `pkg/analyzers/appendr` - `append` statements usage linter
- `pkg/analyzers/elser` - `else` statements usage linter
- `pkg/analyzers/readall` - `ioutil.ReadAll` usage linter

## Usage

bavovna-lint ships as a [golangci-lint module plugin](https://golangci-lint.run/plugins/module-plugins/). Add it to your `.custom-gcl.yml`:

```yaml
version: v2.12.2
name: custom-gcl
destination: .
plugins:
  - module: github.com/gopkgz/bavovna-lint
    version: vX.Y.Z
```

Build a custom golangci-lint binary:

```
golangci-lint custom
```

Enable `bavovna` in `.golangci.yml`:

```yaml
linters:
  enable:
    - bavovna
  settings:
    custom:
      bavovna:
        type: module
        description: bavovna-lint analyzers (appendr, elser, readall).
        original-url: github.com/gopkgz/bavovna-lint
```

Run:

```
./custom-gcl run ./...
```

Skip files via the standard golangci-lint `issues.exclude-rules` / `run.exclude-dirs` mechanisms in `.golangci.yml`.

## Building and running from source

```
make lint    # builds custom-gcl and runs it on this repo
```

## Development

```
make ci      # vet + fmt + lint + test
make cover   # coverage HTML at build/coverage.html
make clean   # remove build/
```
