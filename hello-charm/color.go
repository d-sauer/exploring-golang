package main

import "fmt"

type Color struct {
	r, g, b int
}

func NewRGBColor(r, g, b int) Color {
	return Color{r: r, g: g, b: b}
}

func NewUint32Color(c uint32) Color {
	return Color{
		r: int((c >> 16) & 0xff),
		g: int((c >> 8) & 0xff),
		b: int(c & 0xff),
	}
}

func (c *Color) ToUint32() uint32 {
	return uint32(c.r)<<16 | uint32(c.g)<<8 | uint32(c.b)
}

func (c *Color) ToHex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.r, c.g, c.b)
}
