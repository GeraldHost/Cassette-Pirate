package cassettepirate

import (
  "fmt"
  "strconv"
  "encoding/binary"
)

// get average amplitude from byte array
func averageAmplitude(sample []byte) int {
  sum := 0
  for _, a := range sample {
    sum+=int(a)
  }
  return sum / len(sample)
}

// convery amplitude to it's representitve bit
// amplitudes of value 0-10 are 0 bit
// amplitudes of value 245-255 are 1 bit
// eveything else gets ignored so we return -1
func amplitudeToBit(amplitude int) int {
  if amplitude < 128 {
    return 0
  } else if amplitude > 128 {
    return 1
  } else {
    return -1
  }
}

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

// Convert input bytes to binary bytes eg: 010101
func BinaryStr(bytes []byte) []byte {
  resp := make([]byte, 0)
  for _, b := range bytes {
    bin := fmt.Sprintf("%08b", b)
    resp = append(resp, []byte(bin)...)
  }
  return resp
}

// Stolen: https://stackoverflow.com/questions/48646580/convert-a-bitstring-into-a-byte-array
func BinaryStrAsByteSlice(b []byte) []byte {
    var out []byte
    var str string

    for i := len(b); i > 0; i -= 8 {
        if i-8 < 0 {
            str = string(b[0:i])
        } else {
            str = string(b[i-8 : i])
        }
        v, err := strconv.ParseUint(str, 2, 8)
        if err != nil {
            panic(err)
        }
        out = append([]byte{byte(v)}, out...)
    }
    return out
}

