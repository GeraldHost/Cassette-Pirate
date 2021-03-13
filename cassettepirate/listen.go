package cassettepirate

import (
  "os"
  "os/signal"
  "io/ioutil"
  "strings"
  "fmt"
  "syscall"
  "regexp"
  "github.com/gordonklaus/portaudio"
)

const terminalClearLine = "\r\x1b[2K"

var re = regexp.MustCompile("[01]+0{1}1{8}")

func check(err error) {
  if err != nil {
    fmt.Println(err)
    panic(0)
  }
}

// convert bits we get from listening the audio to binary string
func BitsToBinStr(bits []byte) string {
  binStr := strings.Trim(strings.Replace(fmt.Sprint(bits), " ", "", -1), "[]")
  return re.ReplaceAllString(binStr, "")
}

func ListenForInput() {
  nSamples := 1
  framesPerBuffer := make([]byte, nSamples * effectiveBitsPerSample)
  portaudio.Initialize()
  stream, err := portaudio.OpenDefaultStream(
    channelCount, 
    0, 
    float64(sampleRate), 
    len(framesPerBuffer), 
    framesPerBuffer)

  if err != nil {
    fmt.Println(err)
    return
  }

  check(stream.Start())
  
  bits := make([]byte, 0)
  c := make(chan os.Signal, 1)
  signal.Notify(c, syscall.SIGINT)
  fmt.Println("[*] listening on default stream. Send Interrupt signal to stop")

  Loop:
  for {
    check(stream.Read())
    bin := ParseInput([]byte(framesPerBuffer))
    bits = append(bits, bin...)
    select {
    case _ = <-c:
      fmt.Printf("%s[*] stopped listening\n", terminalClearLine)
      break Loop
    default:
    }
  }
  
  binStr := BitsToBinStr(bits)
  bytes := BinaryStrAsByteSlice([]byte(binStr))
 
  outputFilePath := "output.bin"
  err = ioutil.WriteFile(outputFilePath, bytes, 0644)
  if err != nil {
    fmt.Printf("[!] failed to write binary file %s\n", outputFilePath)
    return
  }

  fmt.Printf("[*] binary file written to: %s\n", outputFilePath)
}

func ParseInput(input []byte) []byte {
  // need to read each sample by chunking effectiveBitsPerSample
  // each sample will represent either a 1 or a 0
  // we will sum the sample and divide by the lenght to get the average amplitude
  bits := make([]byte, 0)

  for i := 0; i < len(input); i += effectiveBitsPerSample {
    sample := input[i:i+effectiveBitsPerSample]
    amplitude := averageAmplitude(sample)
    fmt.Println(amplitude)
    bit := amplitudeToBit(amplitude)
    if bit != -1 {
      bits = append(bits, byte(bit))
    }
  }

  return bits
}
