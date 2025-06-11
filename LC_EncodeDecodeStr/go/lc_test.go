package lc

import (
	"reflect" // For deep comparison of slices
	"testing"
)

func TestEncodeDecodeStrings(t *testing.T) {
	tests := []struct {
		name string
		strs []string
	}{
		{
			name: "Empty list",
			strs: []string{},
		},
		{
			name: "List with one empty string",
			strs: []string{""},
		},
		{
			name: "List with multiple empty strings",
			strs: []string{"", "", ""},
		},
		{
			name: "List with one non-empty string",
			strs: []string{"hello"},
		},
		{
			name: "Basic multiple strings",
			strs: []string{"hello", "world"},
		},
		{
			name: "Strings with numbers",
			strs: []string{"123", "45678"},
		},
		{
			name: "Strings with special characters and delimiters",
			strs: []string{"a#b", "c#d#e", "f#g#h#i"}, // Contains the '#' delimiter
		},
		{
			name: "Strings with spaces",
			strs: []string{"go lang", "is awesome", "!"},
		},
		{
			name: "Strings with Unicode characters",
			strs: []string{"你好", "世界", "Go語言"}, // Example with Chinese characters
		},
		{
			name: "Long strings",
			strs: []string{"abcdefghijklmnopqrstuvwxyz", "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"},
		},
		{
			name: "Mixed length strings",
			strs: []string{"a", "bb", "ccc", "dddd"},
		},
		{
			name: "Mixed cases",
			strs: []string{"Apple", "banana", "Cherry"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 1. Encode the original list of strings
			encodedStr := encode(tt.strs)

			// 2. Decode the encoded string
			decodedStrs := decode(encodedStr)

			// 3. Compare the decoded list with the original list
			if !reflect.DeepEqual(decodedStrs, tt.strs) {
				t.Errorf("Encode/Decode mismatch for input: %v\nEncoded: %q\nDecoded: %v\nExpected: %v",
					tt.strs, encodedStr, decodedStrs, tt.strs)
			} else {
				t.Logf("Test Passed: %s -> %q -> %v", tt.strs, encodedStr, decodedStrs)
			}
		})
	}
}
