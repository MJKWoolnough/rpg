//+build !js

package main // import "vimagination.zapto.org/rpg"

import (
	"log"

	_ "vimagination.zapto.org/engine/graphics/gl21"
	_ "vimagination.zapto.org/engine/text/gl21"
	_ "vimagination.zapto.org/engine/windows/glfw32"
)

func main() {
	err := run()
	if err != nil {
		log.Printf("%s", err)
	}

}
