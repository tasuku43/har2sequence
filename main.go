package main

import (
	"flag"
	"fmt"
	"har2sequence/pkg/config"
	"har2sequence/pkg/har"
	"log"
)

func main() {
	configPath := flag.String("config", "", "Path to the config file")
	harPath := flag.String("har", "", "Path to the HAR file")
	flag.Parse()

	var c config.Config
	var err error

	if *configPath != "" {
		c, err = config.LoadConfig(*configPath)
		if err != nil {
			log.Fatalf("Error loading loadConfig: %v", err)
		}
	}

	if *harPath == "" {
		log.Fatalf("HAR file path must be specified")
	}

	harData, err := har.LoadHAR(*harPath)
	if err != nil {
		log.Fatalf("Failed to load HAR file: %v", err)
	}

	sequenceDiagram := harData.GenerateSequenceDiagram(c)

	fmt.Println(sequenceDiagram.Render())
}
