package main

import (
	"fmt"
	"strconv"
)

func StatsAverage(matches []Match, stat string) string {
	var total float64
	count := float64(len(matches))

	for _, match := range matches {
		var value float64
		switch stat {
		case "Kills":
			value = parseFloat(match.Stats.Kills)
		case "Deaths":
			value = parseFloat(match.Stats.Deaths)
		case "Assists":
			value = parseFloat(match.Stats.Assists)
		case "HeadshotsPercentage":
			value = parseFloat(match.Stats.HeadshotsPercentage)
		case "Adr":
			value = parseFloat(match.Stats.Adr)
		case "KdRatio":
			value = parseFloat(match.Stats.KdRatio)
		case "KrRatio":
			value = parseFloat(match.Stats.KrRatio)
		default:
			return "Invalid statistic type"
		}
		total += value
	}

	if count == 0 {
		return "No data"
	}

	average := total / count

	if stat == "HeadshotsPercentage" {
		return fmt.Sprintf("%.2f%%", average) // Format as percentage
	}

	return fmt.Sprintf("%.2f", average)
}

func parseFloat(s string) float64 {
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err, "Value:", s)
		return 0.0
	}
	return result
}
