package helpers

import (
	"io"
	"log"
)

func Closer(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("close error: %v", err)
	}
}

