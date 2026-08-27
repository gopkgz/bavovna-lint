# bavovna-lint

opinionated golang linters

## Linters

- `pkg/analyzers/appendr` - `append` statements usage linter
- `pkg/analyzers/elser` - `else` statements usage linter
- `pkg/analyzers/readall` - `ioutil.ReadAll` usage linter

## Usage

bavovna-lint ships as a [golangci-lint module plugin](https://golangci-lint.run/plugins/module-plugins/). Add it to your `.custom-gcl.yml`:

```yaml
version: v2.13.1
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

Enable the linters in `.golangci.yml`. There are two packagings — pick **one**, not both. The bundle and the standalone linters share the same underlying analyzers, so enabling a bundle together with any of its constituent linters makes golangci-lint attribute a finding to a nondeterministic linter name (last writer wins over a map iterated in random order); suppression then becomes unreliable. Enable one mode, not both.

### Granular (recommended)

Each analyzer is its own linter, so `//nolint:appendr` / `//nolint:elser` / `//nolint:readall` are recognized and findings are attributed to that name:

```yaml
linters:
  enable:
    - appendr
    - elser
    - readall
  settings:
    custom:
      appendr:
        type: module
        description: flags append usage.
        original-url: github.com/gopkgz/bavovna-lint
      elser:
        type: module
        description: flags unnecessary else.
        original-url: github.com/gopkgz/bavovna-lint
      readall:
        type: module
        description: flags ioutil.ReadAll usage.
        original-url: github.com/gopkgz/bavovna-lint
```

### Bundle (back-compat)

All three analyzers under one linter name; suppress with `//nolint:bavovna`:

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

### Suppressing a finding

Suppression uses golangci-lint's standard `//nolint` handling — bavovna no longer filters comments internally. That handling is **name-specific** and **single-line**:

- Name the linter: `//nolint:appendr` / `//nolint:elser` / `//nolint:readall` in granular mode, or `//nolint:bavovna` in bundle mode.
- Put the directive on the line of the flagged token. For a multi-line `append(...)`, `else`, or `ioutil.ReadAll(...)`, that is the line carrying the keyword — not a continuation or closing line.
- Migration caveat: a bare `//nolint` still suppresses, but a `//nolint:<other-linter>` that merely shares the line no longer does. The old internal handling matched by position and ignored the name; golangci-lint matches by name.

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

## Releases

Daily `bavovna-lint-<version>-<architecture>` releases contain matching `golangci-lint` and `bavovna-lint` binaries, checksums, and a build manifest for Linux amd64 and arm64.

## Public Go CI images

This repository builds its own public Go 1.27.0 Alpine carriers:

- `ghcr.io/gopkgz/bavovna-lint-ci-go-nocgo` for normal build and lint work
- `ghcr.io/gopkgz/bavovna-lint-ci-go-cgo` for race tests that require cgo

Each publication has an exact Go-version tag and an immutable `sha-<source commit>`
tag. Consumers pin a tag and digest. A rerun preserves an existing source tag
instead of pushing it again, while repairing the Go-version tag from that
preserved digest if an earlier publication stopped partway through. The workflow
then logs out and pulls the source tag by digest before accepting either a new or
existing publication, so a package that is not anonymously readable cannot pass
publication.
