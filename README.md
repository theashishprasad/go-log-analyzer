# Go Log Analyzer

A simple Go CLI application that analyzes application log files and generates statistics for log levels such as INFO, WARN, and ERROR.

This project was built as part of my Go learning journey focused on backend, cloud, and platform engineering concepts.

## Project Goal

The goal of this project is to:

* Learn Go project structure and package organization
* Practice file handling and line-by-line processing
* Work with structs, functions, and error handling
* Analyze log files and generate useful statistics
* Build a foundation for more advanced Go CLI tools

## Features

### Current Features

* Read log files from disk
* Count INFO log entries
* Count WARN log entries
* Count ERROR log entries
* Handle file access errors gracefully
* Modular project structure using Go packages

### Planned Features

* Accept log file path as a command-line argument
* Display total log entries processed
* Generate log level percentages
* Track most frequent error messages
* Export results as JSON
* Add unit tests

## Project Structure

```text
go-log-analyzer/
├── analyzer/
│   └── analyzer.go
├── logs/
│   └── app.log
├── main.go
├── go.mod
└── README.md
```

## How to Run

Clone the repository:

```bash
git clone https://github.com/theashishprasad/go-log-analyzer.git
cd go-log-analyzer
```

Run the application:

```bash
go run main.go
```

## Sample Input

Example log file:

```text
INFO Application started
INFO User logged in
WARN Slow response
ERROR Database timeout
INFO Request completed
ERROR API failed
WARN Cache miss
INFO Health check passed
```

## Sample Output

```text
Log Analysis Result

INFO  : 4
WARN  : 2
ERROR : 2
```

## Technologies Used

* Go
* bufio.Scanner
* File I/O
* Structs
* Functions
* Error Handling
* Packages

## Learning Outcomes

Through this project I practiced:

* Reading files in Go
* Processing text line-by-line
* Working with structs
* Creating reusable packages
* Returning multiple values from functions
* Handling errors using idiomatic Go patterns
* Building command-line applications

## Future Enhancements

* Command-line argument support
* JSON export
* Concurrent log processing
* Unit tests
* Docker support
* Kubernetes deployment examples