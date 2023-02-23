package main

import (
	"github.com/rs/zerolog/log"
	"lpotl-go/http"
)

// TODO: Inject version dynamically
var Version = "v0.1.0"

func main() {
	log.Printf("Starting LPOTL-GO version: %s\n", Version)
	s := http.NewServer()
	s.Init(Version)
	s.Run()
}
