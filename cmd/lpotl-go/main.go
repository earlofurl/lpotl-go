package main

import (
	"github.com/earlofurl/lpotl-go/http"
	"github.com/rs/zerolog/log"
)

// TODO: Inject version dynamically
var Version = "v0.1.666"

func main() {
	log.Printf("Starting LPOTL-GO version: %s\n", Version)

	s := http.NewServer()
	s.Init(Version)
	s.Run()
}
