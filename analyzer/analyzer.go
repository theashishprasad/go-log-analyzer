package analyzer

import (
	"bufio"
	"os"
	"strings"
)

type LogStats struct {
	Counts    map[string]int
	TopErrors map[string]int
}

func AnalyzeLog(filePath string) (LogStats, error) {
	stats := LogStats{
		Counts:    make(map[string]int),
		TopErrors: make(map[string]int),
	}

	file, err := os.Open(filePath)
	if err != nil {
		return LogStats{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		if len(fields) == 0 {
			continue
		}

		level := fields[0]

		stats.Counts[level]++

		if level == "ERROR" {
			message := strings.Join(fields[1:], " ")

			if message != "" {
				stats.TopErrors[message]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return LogStats{}, err
	}

	return stats, nil
}