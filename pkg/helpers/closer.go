package helpers

import (
	"io"
	"log"
)

func Closer(closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Printf("Error occured while close: %v", err)
	}
}
