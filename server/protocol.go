package main

import (
	"net"

	"github.com/MJKWoolnough/rpg/internal/protocol"
)

var layers [][]byte

func handleConn(c net.Conn) {
	r := protocol.NewReader(c)
	w := protocol.NewWriter(c)
	for {
		switch r.ReadUint8() {
		case protocol.Close:
			c.Close()
			return
		case protocol.LayerList:
			w.WriteUint32(uint32(len(layers)))
			for _, l := range layers {
				w.WriteBytes(l)
			}
		case protocol.LayerData:
			layer := r.ReadUint8()
			if int(layer) >= len(layers) {
				w.WriteUint8(0)
				continue
			}
			w.WriteBytes(layers[layer])
		}
	}
}
