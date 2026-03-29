# Calculator

A command-line calculator built in Go that evaluates basic arithmetic expressions on integer operands.

## Usage

```bash
go build -o calculator .
./calculator --lhs 10 --rhs 5 --operator "+"
# 10 + 5 = 15
```

### Flags

| Flag | Alias | Description |
|------|-------|-------------|
| `--lhs` | `-l` | Left-hand side operand (>= 0) |
| `--rhs` | `-r` | Right-hand side operand (>= 0) |
| `--operator` | `-o` | Arithmetic operator (`+`, `-`, `*`, `/`) |

## Tests

```bash
go test ./...
```

## Linter

Install `golanci-lint`

```bash
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
```

Install `lefthook`

```bash
brew install lefthook
```

Run `lefthook`

```bash
lefthook run pre-commit --all-files
```