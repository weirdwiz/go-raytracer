package main

import (
	"fmt"
	"io"
	"log"
)

func WriteColor(stream io.Writer, v *Vector, maxVal float64) (int, error) {
	r := v.GetX()
	g := v.GetY()
	b := v.GetZ()

	rByte := int(maxVal * r)
	gByte := int(maxVal * g)
	bByte := int(maxVal * b)

	log.Println(rByte, gByte, bByte)
	return stream.Write([]byte(fmt.Sprintf("%d %d %d ", rByte, gByte, bByte)))
}
