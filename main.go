package main

func main() {
	ctx, err := getContext(640, 480)
	if err != nil {
		logPrintf("error initialising context: %s", err)
		return
	}
	defer ctx.Close()
}
