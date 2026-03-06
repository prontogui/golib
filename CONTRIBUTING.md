# Contributing to ProntoGUI™ Go Library

Thank you for your interest in contributing to the ProntoGUI Go Library! This is the core Go module for [ProntoGUI](https://github.com/prontogui), providing primitives, fields, and communication infrastructure.

## Getting Started

### Prerequisites

- Go 1.23 or later
- A working Go development environment
- Protocol Buffer compiler (`protoc`) and Go gRPC plugins (see [README](README.md) for setup)

### Building and Testing

Clone the repository and verify everything works:

```bash
git clone https://github.com/prontogui/golib.git
cd golib
go build ./...
go test ./...
```

To regenerate protobuf code:

```bash
make
```

## How to Contribute

### Reporting Issues

- Use the GitHub issue tracker to report bugs or suggest improvements.
- Include the Go version, OS, and steps to reproduce the issue.

### Submitting Changes

1. Fork the repository.
2. Create a feature branch from `main`:
   ```bash
   git checkout -b my-feature
   ```
3. Make your changes (see guidelines below).
4. Ensure the project builds and tests pass:
   ```bash
   go build ./...
   go vet ./...
   go test ./...
   ```
5. Commit your changes with a clear, descriptive message.
6. Push to your fork and open a pull request against `main`.

### Adding a New Primitive

Please email us directly at support@prontogui.com and describe the primitive you would like to propose. The code representing the primitives is generated from a specification that is controlled by ProntoGUI, LLC. New primitives are added based on demand and our roadmap for the product.

## Code Guidelines

- Follow standard Go conventions and formatting (`gofmt`).
- Write unit tests for new functionality.
- Use clear variable and function names.
- Include the BSD 3-Clause license header in all new Go source files.

## License

By contributing, you agree that your contributions will be licensed under the [BSD 3-Clause License](LICENSE) that covers this project.

## Questions?

If you have questions about contributing, feel free to open an issue for discussion.
