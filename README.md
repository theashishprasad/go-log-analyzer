# Go Log Analyzer

A simple Go CLI application that analyzes application log files and generates statistics for log levels dynamically.

This project was built as part of my Go learning journey focused on backend, cloud, platform engineering, and Go development.

## Project Goal

The goal of this project is to:

* Learn Go project structure and package organization
* Practice file handling and line-by-line processing
* Work with structs, maps, functions, and error handling
* Analyze log files and generate useful statistics
* Build a foundation for more advanced Go CLI tools
* Understand how to design scalable and maintainable Go code

## Features

### Current Features

* Read log files from disk
* Detect log levels dynamically
* Count occurrences of each log level using maps
* Support any log level (INFO, WARN, ERROR, DEBUG, TRACE, etc.)
* Handle file access errors gracefully
* Modular project structure using Go packages
* Safe handling of empty lines in log files

### Planned Features

* Accept log file path as a command-line argument
* Display total log entries processed
* Generate log level percentages
* Track most frequent error messages
* Export results as JSON
* Add unit tests
* Sort output by log level or frequency

## Project Structure

```text
go-log-analyzer/
├── analyzer/
│   └── analyzer.go
├── log/
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
INFO App started

DEBUG Request received
TRACE Request payload
WARN Cache miss


ERROR Database timeout
INFO User logged in
DEBUG Response generated
```

## Sample Output

```text
Log Analysis Result

INFO   : 2
DEBUG  : 2
TRACE  : 1
WARN   : 1
ERROR  : 1
```

> Note: Output order may vary because Go maps do not guarantee iteration order.

## Technologies Used

* Go
* bufio.Scanner
* File I/O
* Maps
* Structs
* Functions
* Error Handling
* Packages

## Learning Outcomes

Through this project I practiced:

* Reading files in Go
* Processing text line-by-line
* Working with maps for dynamic data aggregation
* Creating reusable packages
* Returning multiple values from functions
* Handling errors using idiomatic Go patterns
* Building command-line applications
* Refactoring code from a fixed design to a scalable design

## Key Concepts Learned

### Version 1

* Structs
* Functions
* Packages
* File handling
* Error handling

### Version 2

* Maps
* Dynamic aggregation
* map[string]int
* make()
* Iterating over maps
* Scalable data structures

## Future Enhancements

* Command-line argument support
* Total log line count
* Log level percentages
* Top error frequency analysis
* JSON export
* Concurrent log processing
* Unit tests
* Docker support
* Kubernetes deployment examples