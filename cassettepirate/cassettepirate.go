package cassettepirate

import (
  "fmt"
  "io/ioutil"
  "encoding/binary"
)

var sampleRate int = 44100
var bitsPerSample int = 8
var channelCount int = 1
var formatType int = 1 // 1 is PCM 1 byte int

// Convert int to byte array padding with n bytes
func U32LittleEndianInt(number int) []byte {
  b := make([]byte, 4)
  binary.LittleEndian.PutUint32(b[0:], uint32(number))
  return b
}

func U16LittleEndianInt(number int) []byte {
  b := make([]byte, 2)
  binary.LittleEndian.PutUint16(b[0:], uint16(number))
  return b
}

func WavFileHeader(dataSize int) []byte {
  magic := []byte("RIFF")
  fileSize := dataSize + 44 // TODO: we need to calculate this once we know the size of the header? 4 byte int 
  fileTypeHeader := []byte("WAVE")
  
  formatMarker := []byte("fmt ")
  formatLength := 16 // TODO: what should this value actually be? 4 byte int
  
  k := (bitsPerSample * sampleRate * channelCount) / 8 // Not sure what this is -> (Sample Rate * BitsPerSample * Channels) / 8. 4 byte int
  q := (bitsPerSample * channelCount) / 8 // also not sure what this is 1 byte

  dataMarker := []byte("data")

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

func BinaryStr(bytes []byte) []byte {
  resp := make([]byte, 0)
  for _, b := range bytes {
    bin := fmt.Sprintf("%b", b)
    resp = append(resp, []byte(bin)...)
  }
  return resp
}

// convery binary data to wav audio
// TODO: currently just returning dummy audio
func BinaryStringToWav() []byte {
  resp := make([]byte, 0)
  bin := BinaryStr([]byte("hello world"))
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

