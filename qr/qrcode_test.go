package main

import (
	"errors"
	"testing"
)

type ErrorWriter struct{}

func (this *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error")
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	w := new(ErrorWriter)
	err := GenerateQRCode(w, "123-45678", Version(1))

	if err == nil || err.Error() != "Expected error" {
		t.Errorf("Error not propagated correctly, got %v", err)
	}
}

func TestVersionDeterminesSize(t *testing.T) {
	table := []struct {
		version  int
		expected int
	}{
		{1, 21},
		{2, 25},
		{6, 41},
		{7, 45},
		{14, 73},
		{40, 177},
	}

	for _, test := range table {
		size := Version(test.version).PatternSize()

		if size != test.expected {
			t.Errorf("Version %2d, expected %3d, but got %3d", test.version, test.expected, size)
		}
	}
}
