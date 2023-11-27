package main

import (
	"3-Minute-Sleep/pkgs/winInteractions"
	"fmt"
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

		displayActive, err := winInteractions.IsDisplaying()

		if err != nil {
			log.Fatal(err)
		}

		prevResults[i] = displayActive

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

		fmt.Print(i, prevResults)
		time.Sleep(CheckInterval)
		i++

	}

}
