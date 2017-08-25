package protocol

import (
	"io"

	"github.com/MJKWoolnough/byteio"
)

const (
	Close byte = iota
	TakeControl
	LayerList
	LayerData
)

type Reader struct {
	byteio.StickyLittleEndianReader
}

func (r *Reader) ReadBytes() []byte {
	if r.Err != nil {
		return nil
	}
	b := make([]byte, r.ReadUint32())
	r.Read(b)
	return b
}

func NewReader(r io.Reader) *Reader {
	return &Reader{byteio.StickyLittleEndianReader{Reader: r}}
}

type Writer struct {
	byteio.StickyWriter
}

func (w *Writer) WriteBytes(s []byte) {
	if w.Err != nil {
		return
	}
	w.WriteUint16(uint16(len(s)))
	w.Write(s)
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{StickyWriter: byteio.StickyLittleEndianWriter{Writer: w}}
}
