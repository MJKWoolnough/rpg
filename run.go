package main

import "fmt"

func run() error {
	ctx, err := getContext(640, 480)
	if err != nil {
		return fmt.Errorf("error initialising context: %s", err)
	}
	defer ctx.Close()
	return nil
}
