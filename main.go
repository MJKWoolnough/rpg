//+build !js

package main

import (
	"log"

	_ "github.com/MJKWoolnough/engine/graphics/gl21"
	_ "github.com/MJKWoolnough/engine/windows/glfw32"
)

func main() {
	err := run()
	if err != nil {
		log.Printf("%s", err)
	}

}
