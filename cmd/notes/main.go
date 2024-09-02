package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/jMurad/notes/internal/app/notes"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/notes.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := notes.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := notes.Start(config); err != nil {
		log.Fatal(err)
	}
}
