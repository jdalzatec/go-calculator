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

* [Install `golanci-lint`](https://golangci-lint.run/docs/welcome/install/local/).
* [Install `lefthook`](https://lefthook.dev/install/)
* Run `lefthook`
    ```bash
    lefthook run pre-commit --all-files
    ```