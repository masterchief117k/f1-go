package parser

import (
	"math"
	"testing"
)

// first go is bing a bitch so we vonver float64 to *float64
func floatPtr(v float64) *float64 {
	return &v
}
func TestParseLapTimeString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    *float64
		expectError bool
	}{
		{
			name:        "Valid sector time",
			input:       "24.120",
			expected:    floatPtr(24.120),
			expectError: false,
		},
		{
			name:        "Valid full lap time",
			input:       "1:22.450",
			expected:    floatPtr(82.450),
			expectError: false,
		},
		{
			name:        "Empty string should return nil without error",
			input:       "",
			expected:    nil,
			expectError: false,
		},
		{
			name:        "Null string should return nil without error",
			input:       "null",
			expected:    nil,
			expectError: false,
		},
		{
			name:        "Manformed time format",
			input:       "1:20:30.123",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Non-numeric characters",
			input:       "1:inbalid",
			expected:    nil,
			expectError: true,
		},
	}
	// let us do test case thing for testcase
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseLapTimeString(tc.input)
			if (err != nil) != tc.expectError {
				t.Fatalf("expepct error: %v, got: %v", tc.expectError, err)
			}
			if !tc.expectError {
				if tc.expected == nil {
					if got != nil {
						t.Errorf("expected nil pointer, got: %v", *got)
					}
				} else {
					if got == nil {
						t.Fatalf("expect %v, got nil pointer", *tc.expected)
					}
					//don't leave bug in testing conver float to abs  so that we dont miss the second sector time
					if math.Abs(*got-*tc.expected) > 0.0001 {
						t.Errorf("expected %v, got %v", *tc.expected, *got)
					}
				}
			}
		})
	}
}

// Test case for lap time
func TestParseLapJSON(t *testing.T) {
	validPayload := []byte(`{
    "lap_number": 44,
    "driver_number": 16,
	"driver_code": "LEC",
    "team_name": "Ferrari",
    "lap_time": "1:22.450",
    "sector1_time": "24.120",
    "sector2_time": "29.810",
    "sector3_time": "28.520",
    "is_deleted": false,
    "compound": "MEDIUM",
    "tyre_age": 12
}`)
	record, err := ParseLapJSON(validPayload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if record == nil {
		t.Fatalf("expected record, got nil")
	}
	if record.Driver.DriverCode != "LEC" {
		t.Errorf("error in metadata: got code %s", record.Driver.DriverCode)
	}
	if record.LapNumber != 44 {
		t.Errorf("error in metadata: got codes lap %d", record.LapNumber)
	}
	if record.Timing.TotalTime == nil {
		t.Fatalf("expected TotalTime pointer, got nil ")
	}
	if math.Abs(*record.Timing.TotalTime-82.450) > 0.0001 {
		t.Errorf("did not recieve expected lap time, got: %v", *record.Timing.TotalTime)
	}
	if record.Timing.Corners == nil {
		t.Fatalf("expected corner times , got nil")
	}

}

func TestParseLapJSON_MalformedJSON(t *testing.T) {
	invalidJSON := []byte(`{ "lap_number": 44, "driver_code":} `)
	record, err := ParseLapJSON(invalidJSON)
	if err == nil {
		t.Fatalf("expected unmarshal error for malformed json , got nil ")
	}
	if record != nil {
		t.Errorf("expected nil record on error, got: %v", record)
	}

}

func TestParseLapJSON_InvalidTimingFormat(t *testing.T) {
	badTimingJSON := []byte(`{
		"lap_number": 44,
		"driver_number": 16,
		"driver_code": "LEC",
		"lap_time": "1:xx.450"
	}`)

	record, err := ParseLapJSON(badTimingJSON)
	if err == nil {
		t.Fatalf("expected timing parse error for corrupted lap time, got nil")
	}
	if record != nil {
		t.Errorf("expected nil record on timing error, got: %v", record)
	}
}
