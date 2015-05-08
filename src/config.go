package main

import (
	"fmt"
	"log"

	"github.com/fjl/go-couchdb"
)

type Config struct {
    URL string
}

// Client() returns a new client for accessing couchdb.
func (c *Config) Client() (*couchdb.Client, error) {
	client, err := couchdb.NewClient(c.URL, nil)

	if err != nil {
		return nil, fmt.Errorf("Error setting up client: %s", err)
	}

	log.Printf("[INFO] CouchDB Client configured for user: %s", c.URL)

	return client, nil
}
