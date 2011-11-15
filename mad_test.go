package mad

import (
	"testing"
	// "fmt"
)

func Test(t *testing.T) {
	const filename = "test.mp3"

	decoder, err := New(filename)
	if (err != nil) {
		t.Errorf("Failed to inialize decoder. %s", err)
	}
	
	decoder.Close()
}