package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func extractAndConvert(value string) float64 {
	re := regexp.MustCompile(`(\d+(\.\d+)?)`) // Extract number (including decimals)
	matches := re.FindString(value)

	if matches == "" {
		return 0
	}

	num, err := strconv.ParseFloat(matches, 64)
	if err != nil {
		return 0
	}

	value = strings.ToLower(value)

	if strings.Contains(value, "cr") {
		num *= 100 // Convert Cr to Lakh
	} else if strings.Contains(value, "thousand") {
		num /= 100

	}

	return num
}

func GetLowestPrice(values []string) string {
	if len(values) == 0 {
		return ""
	}

	lowest := values[0]
	minValue := extractAndConvert(values[0])

	for _, val := range values[1:] {
		converted := extractAndConvert(val)
		if converted < minValue {
			minValue = converted
			lowest = val
		}
	}

	return lowest
}
