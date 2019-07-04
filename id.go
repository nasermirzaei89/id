package id

import (
	"crypto/sha1"
	"fmt"
	"log"
	"time"
)

var id = make(chan string)

func init() {
	go func() {
		h := sha1.New()
		c := []byte(time.Now().String())
		for {
			_, err := h.Write(c)
			if err != nil {
				log.Fatalln(err)
			}

			id <- fmt.Sprintf("%x", h.Sum(nil))
		}
	}()
}

// New returns a new random string
func New() string {
	return <-id
}
