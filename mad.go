// mad is MPEG Audio Decoder.
package mad

// #include <stdlib.h>
// #include "gomad.h"
import "C"

import (
	"os"
	"unsafe"
	"fmt"
)

// File structure represents MP3 file.
type Decoder struct {
	// C decoder structure.
	cDecoder C.struct_gomad_decoder
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
func New(filename string) (file *Decoder, err os.Error) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	file = new(Decoder)
	_, e := C.gomad_open(&(file.cDecoder), cFilename)
	if e != nil {
		return nil, fmt.Errorf("Failed to open file %s: %s", filename, e)
	}

	return file, nil
}

// SampleRate returns file's sample rate value.
func (file *Decoder) SampleRate() int {
	return file.sampleRate
}

// Channels returns number of channels for the related audio file.
func (file *Decoder) Channels() int {
	return file.channels
}

// Length returns file's length in seconds.
func (file *Decoder) Length() int {
	// TODO:

	return -1
}

// Seek move decoding position.
// If whence parameter is SeekSet than position parameter should be non-negative
// integer value, which is new decoding position relative to the start of the file.
// If whence equals SeekCurrent than position parameter can be negative as well as positive integer,
// which means new position related to the current one.
func (file *Decoder) Seek(position int, whence Whence) os.Error {
	// TODO:

	return nil
}

// Read returns up to the specified number of bytes of decoded PCM audio.
// Return number of read 16-bit words.
func (file *Decoder) Read(buf []byte) int {
	// TODO:
	
	return -1
}

// Close release resources assigned to Decoder structure and close MAD decoder.
func (file *Decoder) Close() {
	C.gomad_close(&(file.cDecoder))
}
