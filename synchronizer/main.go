package main

import "go.vocdoni.io/dvote/log"

func main() {

	//Need to execute the smc function to send the event

	// create the service to read the event

	// Config load
	cfg, err := Load()
	if err != nil {
		log.Debug("Config Load failed!: ", err)
	}

	// Create Synchronizer
	synchronizer, err := NewSynchronizer(*cfg)
	if err != nil {
		log.Debug("Synchronizer create failed: ", err)
	}

	go func() {
		if err := synchronizer.Sync(); err != nil {
			log.Fatal(err)
		}
	}()
}
