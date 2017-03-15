//+build !js

package main

import (
	"log"

	_ "github.com/MJKWoolnough/engine/graphics/gles2"
	_ "github.com/MJKWoolnough/engine/windows/glfw31"
)

func main() {
	err := run()
	if err != nil {
		log.Printf("%s", err)
	}

}
