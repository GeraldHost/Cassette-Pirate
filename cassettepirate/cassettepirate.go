package cassettepirate

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// First 4 bytes of RIFF header
var magic []byte = []byte("RIFF")

// Sample rate e.g 44.1 samples/s
var sampleRate int = 44100

// How many bits in each sample
var bitsPerSample int = 8

// The actual size of binary bit
var effectiveBitsPerSample int = bitsPerSample * 20

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
func BinaryStringToWav(bytes []byte) []byte {
  resp := make([]byte, 0)
  binStr := BinaryStr(bytes)
  for _, b := range binStr {
    c := 0
    if(b == 49) {
      c = 255
    } 
    for i := 0; i < effectiveBitsPerSample; i++ {
      resp = append(resp, byte(c))
    }
  }
  return resp
}

func BinaryToWav(path, outputFilePath string) {
  fmt.Println("[*] Converting binary file to wav")
  // read file bytes
  file, err := os.Open(path)
  if err != nil {
    fmt.Printf("[!] failed to open file %s\n", path)
  }
  defer file.Close()
  reader := bufio.NewScanner(file)
  
  // set up data section
  data := make([]byte, 0)
  
  // scan the binary file and convert it to a binary string
  for reader.Scan() {
    bytes := reader.Bytes()
    data = append(data, BinaryStringToWav(bytes)...)
  }

  header := WavFileHeader(len(data))
  
  // append header and data sections together
  bin := append(header, data...)

  err = ioutil.WriteFile(outputFilePath, bin, 0644)
  if err != nil {
    fmt.Printf("[!] failed to write wav file %s\n", outputFilePath)
    return
  }

  fmt.Printf("[*] wav file created: %s\n", outputFilePath)
}

