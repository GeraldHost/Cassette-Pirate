package main

import (
  "fmt"
  "flag"
  cp "cassettepirate/cassettepirate"
)

type config struct {
  binary string
}

func parseFlags() *config {
	c := &config{}
	flag.StringVar(&c.binary, "b", "", "The binary file to convert to wav audio")
	flag.Parse()

	return c
}

func header() {
  fmt.Println(`
       _________
      |   ___   |
      |  o___o  |
      |__/___\__|
    cassette pirate
  `)
}

func main() {
  header()

  config := parseFlags()
  
  if len(config.binary) > 0 {
    cp.BinaryToWav(config.binary)
  }

}
