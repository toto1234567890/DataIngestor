package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// -----------------------------------------------------------------------------
// Helper Functions
// -----------------------------------------------------------------------------

// ParseFloat parses a string to float64, returns 0 if parsing fails
func ParseFloat(s string) float64 {
	var f float64
	fmt.Sscanf(s, "%f", &f)
	return f
}

// -----------------------------------------------------------------------------

// ParseInt64 parses a string to int64 with error handling
func ParseInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse int64 from '%s': %w", s, err)
	}
	return i, nil
}

// -----------------------------------------------------------------------------

// MaskAPIKey masks sensitive information in URLs (API keys, tokens, etc.)
func MaskAPIKey(endpoint string) string {
	// Mask query parameters that contain sensitive data
	// Common patterns: ?token=xxx, ?apikey=xxx, ?api_key=xxx, &token=xxx, etc.
	if strings.Contains(endpoint, "?") || strings.Contains(endpoint, "&") {
		parts := strings.Split(endpoint, "?")
		if len(parts) > 1 {
			// Keep the base URL, mask the query parameters
			return parts[0] + "?***"
		}
	}
	return endpoint
}
