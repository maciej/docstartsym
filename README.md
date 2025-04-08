# docstartsym

A Go analyzer that checks if documentation comments start with the name of the symbol they document.

> **Note**: This is an experimental project and is not recommended for production use at this time. It was created while exploring Cursor AI's capabilities.

## Installation

```bash
go install github.com/maciej/docstartsym/cmd/docstartsym@latest
```

## Usage

```bash
docstartsym [flags] [packages]
```

The analyzer will check that documentation comments start with the name of the symbol they document. For example:

```go
// Incorrect - documentation doesn't start with the function name
func GetUser() {}

// GetUser retrieves a user from the database - correct
func GetUser() {}
```

## License

MIT 