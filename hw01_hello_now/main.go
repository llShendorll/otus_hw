package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	timeNtp, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("current time: %s\nexact time: %s\n", time.Now(), timeNtp.UTC())
}
