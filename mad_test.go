package mad

import (
	"testing"
	"fmt"
)

func Test(t *testing.T) {
	const filename = "test.mp3"

	decoder, err := New(filename)
	if (err != nil) {
		t.Errorf("Failed to inialize decoder. %s", err)
	}

	fmt.Printf("File name: %s\n", filename)
	fmt.Printf("Length: %d\n", decoder.Length())
	fmt.Printf("Sample rate: %d\n", decoder.SampleRate())
	fmt.Printf("Number of channels: %d\n", decoder.Channels())
	
	decoder.Close()
}