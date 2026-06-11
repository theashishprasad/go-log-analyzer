package main

import (
	"fmt"
	"os"

	"github.com/theashishprasad/log-analyzer/analyzer"
)

func main() {
	arguments := os.Args

	if len(arguments) != 2 {
		fmt.Println("Usage: go run main.go <log-file>")
		return
	}

	filename := arguments[1]

	stats, err := analyzer.AnalyzeLog(filename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Log Analysis Result")
	fmt.Println()

	total := 0

	for _, val := range stats.Counts {
		total += val
	}

	if total == 0 {
		fmt.Println("Log file has no log levels.")
		return
	}

	fmt.Printf("Total Logs : %d\n\n", total)

	for key, val := range stats.Counts {
		percentage := (float64(val) / float64(total)) * 100

		fmt.Printf("%s : %d (%.2f%%)\n",
			key,
			val,
			percentage,
		)
	}

	fmt.Println("\nTop Errors\n")

	if len(stats.TopErrors) == 0 {
		fmt.Println("No errors found")
		return
	}

	for key, val := range stats.TopErrors {
		fmt.Printf("%s : %d\n", key, val)
	}
}