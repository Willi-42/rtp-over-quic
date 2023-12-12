package main

import (
	"log"

	"github.com/Willi-42/rtp-over-quic/cmd"
)

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)
	cmd.Execute()
}
