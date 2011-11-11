// mad is MPEG Audio Decoder.
package mad

import "C"

import (
	"os"
)

// File structure represents MP3 file.
type File struct {
	// Sample rate of the file.
	sampleRate int
	// Number of channels in the current file.
	channels int
}

// Whence type for Seek function.
type Whence int

const (
	// Set new position relative to the start.
	SeekSet = iota
	// Set new position related to the current one.
	SeekCurrent
)

// New opens and initialize MAD decoder and File structure.
func New(filename string) (file *File, err os.Error) {
	// TODO:
	
	return nil, nil
}

// SampleRate returns file's sample rate value.
func (file *File) SampleRate() int {
	return file.sampleRate
}

// Channels returns number of channels for the related audio file.
func (file *File) Channels() int {
	return file.channels
}

// Length returns file's length in seconds.
func (file *File) Length() int {
	// TODO:

	return -1
}

// Seek move decoding position.
// If whence parameter is SeekSet than position parameter should be non-negative
// integer value, which is new decoding position relative to the start of the file.
// If whence equals SeekCurrent than position parameter can be negative as well as positive integer,
// which means new position related to the current one.
func (file *File) Seek(position int, whence Whence) os.Error {
	// TODO:

	return nil
}

// Read returns up to the specified number of bytes of decoded PCM audio.
// Return number of read 16-bit words.
func (file *File) Read(buf []byte) int {
	// TODO:
	
	return -1
}

// Close release resources assigned to File structure and close MAD decoder.
func (file *File) Close() {
	// TODO:
}
