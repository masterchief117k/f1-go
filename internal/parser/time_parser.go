package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// correcting the time format
func ParseLapTimeString(timeStr string) (*float64, error) {
	if timeStr == "null" || timeStr == "" {

		return nil, nil
	}
	parts := strings.Split(timeStr, ":")
	if len(parts) == 1 {
		sec, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return nil, err
		}
		return &sec, nil
	} else if len(parts) == 2 {
		min, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return nil, err
		}
		sec, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, err
		}
		total := (min * 60) + sec
		return &total, nil
	} else {
		return nil, fmt.Errorf("invalid time format:%s", timeStr)
	}

}
