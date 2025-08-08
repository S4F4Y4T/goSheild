package utils

import (
	"strings"
	"testing"
)

func TestAskChoiceFromReader(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" Hello \n", "hello"},
		{"WORLD\n", "world"},
		{" GoLang  \n", "golang"},
		{"    \n", ""},
	}

	for _, tt := range tests {
		reader := strings.NewReader(tt.input)
		got := askChoiceFromReader(reader)
		if got != tt.expected {
			t.Errorf("askChoiceFromReader(%q) = %q; want %q", tt.input, got, tt.expected)
		}
	}
}
