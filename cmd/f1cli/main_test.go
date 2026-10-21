package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckSeason(t *testing.T) {
	tests := []struct {
		name    string
		year    int
		wantErr bool
	}{
		{"min valid year", 1950, false},
		{"mid valid year", 2000, false},
		{"max valid year", 2025, false},
		{"just below min", 1949, true},
		{"just above max", 2026, true},
		{"negative year", -1, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := checkSeason(tc.year)
			hasErr := err != nil
			if hasErr != tc.wantErr {
				t.Errorf("checkSeason(%d) got error = %v, want error = %v (err: %v)", tc.year, hasErr, tc.wantErr, err)
			}
		})
	}
}

func TestCheckTop(t *testing.T) {
	tests := []struct {
		name    string
		top     int
		wantErr bool
	}{
		{"valid value min", 1, false},
		{"valid value max", 22, false},
		{"negative", -1, true},
		{"zero check", 0, true},
		{"just above max", 23, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := checkTop(tc.top)
			hasErr := err != nil
			if hasErr != tc.wantErr {
				t.Errorf("checkTop(%d) got error = %v, want error = %v (err: %v)", tc.top, hasErr, tc.wantErr, err)
			}
		})
	}
}

func TestCheckData(t *testing.T) {
	tmpDir := t.TempDir()
	validFile := filepath.Join(tmpDir, "exists.csv")
	if err := os.WriteFile(validFile, []byte("mock data"), 0o644); err != nil {
		t.Fatalf("setup failed: could not create test file: %v", err)
	}

	missingFile := filepath.Join(tmpDir, "missing.csv")

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"file exists", validFile, false},
		{"file missing", missingFile, true},
		{"path is a directory", tmpDir, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := checkData(tc.path)
			hasErr := err != nil
			if hasErr != tc.wantErr {
				t.Errorf("checkData(%q) got error = %v, want error = %v (err: %v)", tc.path, hasErr, tc.wantErr, err)
			}
		})
	}
}
