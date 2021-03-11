package cassettepirate

import (
  "fmt"
  "io/ioutil"
)

// First 4 bytes of RIFF header
var magic []byte = []byte("RIFF")

// Sample rate e.g 44.1 samples/s
var sampleRate int = 44100

// How many bits in each sample
var bitsPerSample int = 8

// Number of channels, 1 = mono, 2 = stereo
var channelCount int = 1

// PCM = 1, 1 byte int
var formatType int = 1

// Static file type
var fileTypeHeader []byte = []byte("WAVE")

// marks the start of the format section
var formatMarker []byte = []byte("fmt ")

// Fuck knows what should this value actually be? 4 byte int
var formatLength int = 16

// Not sure what this is -> (Sample Rate * BitsPerSample * Channels) / 8. 4 byte int
var k int = (bitsPerSample * sampleRate * channelCount) / 8 

// also not sure what this is 1 byte
var q int = (bitsPerSample * channelCount) / 8 

// marks the start of the data section
var dataMarker []byte = []byte("data")

// Create wave file header based on the size of the data
func WavFileHeader(dataSize int) []byte {
  fileSize := dataSize + 44 

  parts := []interface{}{
    magic,
    U32LittleEndianInt(fileSize),
    fileTypeHeader,
    formatMarker,
    U32LittleEndianInt(formatLength),
    formatType,
    channelCount,
    U32LittleEndianInt(sampleRate),
    U32LittleEndianInt(k),
    q,
    bitsPerSample,
    dataMarker,
    U32LittleEndianInt(dataSize),
  }

  resp := make([]byte, 0)
  for _, part := range parts {
    switch part.(type) {
    case int:
      resp = append(resp, U16LittleEndianInt(part.(int))...)
    case []byte:
      resp = append(resp, part.([]byte)...)
    }
  }

  return resp
}

// convery binary data to wav audio
// TODO: currently just returning dummy audio
func BinaryStringToWav() []byte {
  resp := make([]byte, 0)
  bin := BinaryStr("hello world")
  for a := 0; a < 1000000; a++ {
    for _, b := range bin {
      c := 0
      if(b == 49) {
        c = 255
      } 
      for i := 0; i < 8; i++ {
        resp = append(resp, byte(c))
      }
    }
  }
  return resp
}

func BinaryToWav(path string) {
  bin := BinaryStringToWav()
  header := WavFileHeader(len(bin))

  bin = append(header, bin...)

  err := ioutil.WriteFile("test.wav", bin, 0644)
  if err != nil {
    fmt.Println("failed to write wav file")
    return
  }

  fmt.Println("binary to wav")
}

