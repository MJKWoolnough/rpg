package main

type Image struct {
	Image  *Texture
	offset Point
	size   Rectangle
}

func (i *Image) OffsetBy(p Point) {
	i.offset = i.offset.Add(p)
}
