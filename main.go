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

	for key, val := range stats.Counts {
		fmt.Printf("%s : %d\n", key, val)
	}
}