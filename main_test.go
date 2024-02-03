package main

import (
	"os"
	"testing"
)

// Test the writeJSON function
func TestWriteJSON(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	info := PageInfo{
		URL:   "Test URL",
		Title: "Test Title",
		Text:  "Test Text"}

	// Call the function to test
	err = writeJSON(info, tempFile)
	if err != nil {
		t.Fatal(err)
	}

	// Read the file's contents
	contents, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Check if the contents
	expected := "{\"url\":\"Test URL\",\"title\":\"Test Title\",\"text\":\"Test Text\"}\n"
	if string(contents) != expected {
		t.Errorf("Expected %q, got %q", expected, contents)
	}
}

func TestStripHTML(t *testing.T) {
	// Test the stripHTML function
	input := "<p>Test String.</p>"
	expected := "Test String."
	result := stripHTML(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
