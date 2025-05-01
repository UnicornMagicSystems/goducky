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

```bash
go get github.com/marcboeker/go-duckdb
```

### Building and Running the Application(If you want the GUI app version of this rename main.go to main.goZZZ and rename main.goGUI to main.go)

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

1. Creates a new DuckDB database file named `mydatabase.duckdb`
2. Creates a `users` table with columns for ID, name, email, and creation timestamp
3. Inserts sample user data
4. Queries and displays the inserted data

## Example Output

When you run the application, you should see output similar to:

```
Connected to DuckDB database
Created 'users' table
Inserted sample data

Users in the database:
----------------------
ID: 1, Name: John Doe, Email: john@example.com
ID: 2, Name: Jane Smith, Email: jane@example.com
ID: 3, Name: Bob Johnson, Email: bob@example.com

DuckDB database saved to: mydatabase.duckdb
```

## Notes

- The application will remove any existing database file with the same name before creating a new one.
- The database file is stored in the same directory as the executable.
- This example uses the [go-duckdb](https://github.com/marcboeker/go-duckdb) driver.

## License

[BSD 3-Clause License](LICENSE)