# Log Analyzer

A simple Go application that analyzes log files and counts:

- INFO messages
- WARN messages
- ERROR messages

## Project Structure

```text
log-analyzer/
├── analyzer/
│   └── analyzer.go
├── logs/
│   └── app.log
├── main.go
├── go.mod
└── README.md
```

## Run

```bash
go run main.go
```

## Sample Output

```text
Log Analysis Result

INFO  : 4
WARN  : 2
ERROR : 2
```