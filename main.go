package main

import (
	"github.com/andrei93r/3-Minute-Sleep/pkgs/winInteractions"
	"log"
	"time"
)

const (
	CheckInterval = 60 * time.Second
)

func main() {
	prevResults := [3]bool{true, true, true}

	i := 0

	for {
		if i > len(prevResults)-1 {
			i = 0
		}

		IsDisplaying, err := winInteractions.IsDisplaying()

		if err != nil {
			log.Fatal(err)
		}

		prevResults[i] = IsDisplaying

		if i == len(prevResults)-1 {

			if prevResults[0] == false && prevResults[1] == false && prevResults[2] == false {
				err := winInteractions.SetSuspendState(
					winInteractions.SLEEP_SLEEP,
					winInteractions.CRITICAL)

				if err != nil {
					log.Fatal(err)
				}
			}
		}

		i++
		time.Sleep(CheckInterval)

	}

}
