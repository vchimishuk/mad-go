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
	defer decoder.Close()

	fmt.Printf("File name: %s\n", filename)
	fmt.Printf("Length: %d\n", decoder.Length())
	fmt.Printf("Sample rate: %d\n", decoder.SampleRate())
	fmt.Printf("Number of channels: %d\n", decoder.Channels())

	
	buf := make([]byte, 4096)
	fmt.Printf("Decoded time:     0")
	for decoder.Read(buf) > 0 {
		fmt.Printf("\b\b\b\b\b%5d", decoder.CurrentPosition())
	}
	fmt.Printf("\n")
}