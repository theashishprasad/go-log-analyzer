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
* Generate percentage distribution for log levels
* Track most frequent error messages
* Support any log level (INFO, WARN, ERROR, DEBUG, TRACE, etc.)
* Handle file access errors gracefully
* Modular project structure using Go packages
* Safe handling of empty lines in log files

### Planned Features

* Accept log file path as a command-line argument
* Export results as JSON
* Add unit tests
* Sort output by log level or frequency
* Concurrent log processing

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
ERROR Database timeout
DEBUG Response generated
ERROR API failed
```

## Sample Output

```text
Log Analysis Result

Total Logs : 8

INFO  : 2 (25.00%)
DEBUG : 2 (25.00%)
TRACE : 1 (12.50%)
WARN  : 1 (12.50%)
ERROR : 3 (37.50%)

Top Errors

Database timeout : 2
API failed : 1
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
* String Processing

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
* Performing percentage calculations
* Formatting floating-point output
* Extracting and processing structured log data
* Aggregating repeated error messages

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

### Version 3

* Percentage calculations
* Integer vs floating-point division
* float64 conversions
* Formatted output with fmt.Printf
* Basic statistical reporting
* Division-by-zero handling

### Version 4

* Nested aggregation
* Error frequency analysis
* strings.Fields()
* strings.Join()
* Log message extraction
* Working with multiple maps
* Basic operational analytics

## Future Enhancements

* Command-line argument support
* JSON export
* Concurrent log processing
* Unit tests
* Sort output by log level frequency
* Sort top errors by occurrence count
* Docker support
* Kubernetes deployment examples

## Example Use Cases

* Application log analysis
* Error trend detection
* Local troubleshooting
* Learning Go data structures
* Building CLI tooling foundations
* Platform engineering practice projects
