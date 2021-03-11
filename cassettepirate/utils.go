package cassettepirate

import (
  "fmt"
  "encoding/binary"
)

// unsign 32 bit int in little endian
func U32LittleEndianInt(number int) []byte {
  b := make([]byte, 4)
  binary.LittleEndian.PutUint32(b[0:], uint32(number))
  return b
}

// unsign 16 bit int in little endian
func U16LittleEndianInt(number int) []byte {
  b := make([]byte, 2)
  binary.LittleEndian.PutUint16(b[0:], uint16(number))
  return b
}

// Convert string to binary bytes eg: 010101
func BinaryStr(str string) []byte {
  resp := make([]byte, 0)
  for _, b := range str[:] {
    bin := fmt.Sprintf("%b", b)
    resp = append(resp, []byte(bin)...)
  }
  return resp
}

