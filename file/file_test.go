package file

import (
	"path/filepath"
	"testing"

	"github.com/sig00rd/simpleserver/helpers"
)

func TestInspectFile(t *testing.T) {

	t.Run("should throw an error when tasked to inspect a file that doesn't exist", func(t *testing.T) {
		path := filePath("../test_files/missingfile.html")

		got := InspectFile(path)

		helpers.AssertError(t, got)
	})

	t.Run("should not throw an error when inspecting a valid html file", func(t *testing.T) {
		path := filePath("../test_files/valid.html")

		got := InspectFile(path)

		helpers.AssertNoError(t, got)
	})

	t.Run("should throw an error when inspecting invalid html file", func(t *testing.T) {
		path := filePath("../test_files/invalid.html")

		got := InspectFile(path)

		helpers.AssertError(t, got)
	})
}

func filePath(path string) string {
	return filepath.FromSlash(path)
}
