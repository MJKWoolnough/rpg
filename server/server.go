package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/MJKWoolnough/byteio"
	"github.com/MJKWoolnough/httpdir"
	"github.com/MJKWoolnough/rpg/internal/protocol"
	"golang.org/x/net/websocket"
)

var (
	httpPort = flag.Int("http", -1, "http port")
	tcpPort  = flag.Int("tcp", -1, "tcp port")
)

var dir http.FileSystem = httpdir.Default

func ws(c *websocket.Conn) {
	c.PayloadType = websocket.BinaryFrame
	handleConn(c)
}

func handleConn(c net.Conn) {
	r := byteio.StickyReader{Reader: byteio.LittleEndianReader{Reader: c}}
	w := byteio.StickyWriter{byteio.LittleEndianWriter{Writer: c}}
	for {
		switch r.ReadUint8() {
		case protocol.Close:
			c.Close()
			return
		case protocol.LayerList:
			w.WriteUint32(len(layers))
			for _, l := range layers {
				w.WriteUint32(uint32(len(l)))
				w.Write(l)
			}
		case protocol.LayerData:
			layer := r.ReadUint8()
			if layer >= len(layers) {
				w.WriteUint8(0)
				continue
			}
			w.Write(layers[layer])
		}
	}
}

func main() {
	flag.Parse()
	if *httpPort < 0 && *tcpPort < 0 {
		flag.Usage()
		return
	}
	s := new(State)
	var wg sync.WaitGroup
	if *httpPort >= 0 {
		l, err := net.Listen("tcp", fmt.Sprintf(":%d", *httpPort))
		if err != nil {
			log.Fatal(err)
		}
		http.Handle("/ws", websocket.Handler(ws))
		http.Handle("/", http.FileServer(dir))
		wg.Add(1)
		go func() {
			log.Println(http.Serve(l, nil))
			wg.Done()
		}()
		log.Println("HTTP server on %s", l.Addr())
	}
	if *tcpPort >= 0 {
		l, err := net.Listen("tcp", fmt.Sprintf(":%d", *tcpPort))
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		go func() {
			for {
				c, err := l.Accept()
				if err != nil {
					if ne, ok := err.(net.Error); ok {
						if ne.Temporary() {
							continue
						}
					}
					log.Println(err)
					break
				}
				go handleConn(c)
			}
			wg.Done()
		}()
		log.Println("TCP server on %s", l.Addr())
	}
	wg.Wait()
}
