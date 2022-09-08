package helpers

import (
	"io"
	"log"
)

// Closer - close connection with err handling
func Closer(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("close error: %v", err)
	}
}

