package main

import (
	"fmt"

	"github.com/theashishprasad/log-analyzer/analyzer"
)

func main() {
	stats, err := analyzer.AnalyzeLog("log/app.log")

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
		fmt.Printf("%s : %d (%.2f%%)\n", key, val, percentage)
	}
}