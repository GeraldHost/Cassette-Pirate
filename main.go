package main

import (
  "fmt"
  "flag"
  cp "cassettepirate/cassettepirate"
)

type config struct {
  binary string
  outfile string
}

func parseFlags() *config {
	c := &config{}
	flag.StringVar(&c.binary, "b", "", "The binary file to convert to wav audio")
  flag.StringVar(&c.outfile, "o", "", "The output wav file")
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
    if len(config.outfile) <= 0 {
      fmt.Println("[!] please provide an out file")
      return
    }
    cp.BinaryToWav(config.binary, config.outfile)
  }

}
