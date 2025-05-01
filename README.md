# Go DuckDB Example

A simple Go application demonstrating how to use DuckDB with Go's standard database/sql package.

## Prerequisites

Before you begin, you'll need to install Go on your system.

### Installing Go

1. Download Go from the [official website](https://golang.org/dl/).
2. Follow the installation instructions for your operating system:
   - **Windows**: Run the downloaded MSI installer and follow the prompts.
   - **macOS**: Open the downloaded package file and follow the prompts, or use Homebrew: `brew install go`.
   - **Linux**: Extract the archive to `/usr/local` with `tar -C /usr/local -xzf go[version].linux-amd64.tar.gz` and add `/usr/local/go/bin` to your PATH.

Verify your installation by running:

```bash
go version
```

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/UnicornMagicSystems/goducky.git
cd goducky
```

### Install Dependencies

The application uses the `go-duckdb` driver to connect to DuckDB. Install it with:
(If you want the GUI app version of this rename main.go to main.goZZZ and rename main.goGUI to main.go)
```bash
go get github.com/marcboeker/go-duckdb
go mod tidy
```

### Building and Running the Application

#### Build and Run Locally

To build and run the application on your current platform:

```bash
go build -o goducky
```

Then run the executable:

```bash
./goducky
```

#### Cross-Compilation

You can compile the application for different operating systems:

**For Windows (from any platform):**

```bash
GOOS=windows GOARCH=amd64 go build -o goducky.exe
```

**For macOS (from any platform):**

```bash
GOOS=darwin GOARCH=amd64 go build -o goducky-macos
```

**For Linux (from any platform):**

```bash
GOOS=linux GOARCH=amd64 go build -o goducky-linux
```

## Application Overview

This application demonstrates basic DuckDB operations using Go:

1. The first time you run the app it will creates a new DuckDB database file named `people.duckdb` if it doesn't already exist.
2. Creates a `people` table with columns for ID, first_name, last_name, and age.
3. Users can create/update/delete the records that are saved to a people.duckdb file.
4. If you close the app and reopen it will read the saved people.duckdb file so you can continue working on your people list.

```

## License

[BSD 3-Clause License](LICENSE)
