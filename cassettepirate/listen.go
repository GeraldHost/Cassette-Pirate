package cassettepirate

import (
  "fmt"
  "github.com/gordonklaus/portaudio"
)

func check(err error) {
  if err != nil {
    fmt.Println(err)
    panic(0)
  }
}

func ListenForInput() {
  framesPerBuffer := make([]byte, 64)
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

  for {
    check(stream.Read())
    fmt.Println([]byte(framesPerBuffer))
  }
}
