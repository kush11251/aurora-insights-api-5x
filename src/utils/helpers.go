package utils

import (
	"errors"
)

// IsValid insight checks if the insight is valid
func IsValidInsight(insight Insight) (bool, error) {
	if insight.Id == "" || insight.Description == "" {
		return false, errors.New("invalid insight")
	}
	return true, nil
}