package main

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteJSONToFile(t *testing.T) {
	tempFile, err := io.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	info := PageInfo{
		URL:   "https://example.com",
		Title: "Test Title",
		Text:  "Test Content",
	}

	// Execute
	err = writeJSON(info, tempFile)
	assert.NoError(t, err)

	// Verify
	tempFile.Seek(0, 0) // Go back to the start of the file
	bytes, err := io.ReadAll(tempFile)
	assert.NoError(t, err)

	var pageInfo PageInfo
	err = json.Unmarshal(bytes, &pageInfo)
	assert.NoError(t, err)

	assert.Equal(t, "Test Title", pageInfo.Title)
	assert.Equal(t, "Test Content", pageInfo.Text)
	assert.Equal(t, "https://example.com", pageInfo.URL)
}
