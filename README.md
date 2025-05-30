# Net-Tools

A network utility toolkit built with Go and Cobra CLI.

## Features

- **Port Scanner**: Scan TCP ports on target hosts
- Interactive CLI interface with tool selection menu

## Installation

```bash
go build -o net-tools.exe .
```

## Usage

Run the application:

```bash
./net-tools.exe
```

The application will show you a menu to select tools:

```
=== Net-Tools ===
Select a tool:
1. Port Scanner
0. Exit

Enter your choice:
```

### Port Scanner

Select option 1 to use the port scanner. You will be prompted to enter:
- Target host (e.g., 192.168.1.1)
- Start port (e.g., 1)
- End port (e.g., 1024)

The scanner will then check all ports in the specified range and report which ones are open.

## Project Structure

```
.
├── cmd/
│   └── root.go          # Cobra CLI root command
├── internal/
│   └── scanner/
│       └── scanner.go   # Port scanner implementation
├── main.go              # Application entry point
├── go.mod               # Go module file
└── README.md            # This file
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework for Go 