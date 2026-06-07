package analyzer

import (
	"bufio"
	"os"
	"strings"
)

type LogStats struct {
	InfoCount  int
	WarnCount  int
	ErrorCount int
}

func AnalyzeLog(filePath string) (LogStats, error) {
	stats := LogStats{}

	file, err := os.Open(filePath)
	if err != nil {
		return LogStats{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "INFO") {
			stats.InfoCount++
		} else if strings.HasPrefix(line, "WARN") {
			stats.WarnCount++
		} else if strings.HasPrefix(line, "ERROR") {
			stats.ErrorCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return LogStats{}, err
	}

	return stats, nil
}