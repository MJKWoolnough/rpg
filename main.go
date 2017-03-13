//+build !js

package main

import "log"

func main() {
	err := run()
	if err != nil {
		log.Printf("%s", err)
	}
}
