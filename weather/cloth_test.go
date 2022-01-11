package weather

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetLikelyWeather(t *testing.T) {
	tests := []struct {
		description     string
		weatherStatuses []string
		expected        string
	}{
		{
			description:     "Must return the duplicated status",
			weatherStatuses: []string{"Rain", "Rain", "Snow"},
			expected:        "Rain",
		},
		{
			description:     "Must return Unknown if no duplicated status",
			weatherStatuses: []string{"Rain", "Sun", "Snow"},
			expected:        "Unknown",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			likelyWeather := getLikelyWeather(tc.weatherStatuses)
			assert.Equal(t, tc.expected, likelyWeather)
		})
	}
}
