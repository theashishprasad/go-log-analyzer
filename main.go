package main

import (
	"fmt"

	"github.com/theashishprasad/log-analyzer/analyzer"
)

func main() {
	stats, err := analyzer.AnalyzeLog("logs/app.log")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Log Analysis Result")
	fmt.Println()

	fmt.Printf("INFO  : %d\n", stats.InfoCount)
	fmt.Printf("WARN  : %d\n", stats.WarnCount)
	fmt.Printf("ERROR : %d\n", stats.ErrorCount)
}