package main

import (
	"fmt"
	"gopkg.in/nishanths/go-hgconfig.v1"
)

func main() {
	value, err := hgconfig.Get("merge-tools.editmerge.premerge")
	fmt.Println("value:", value)
	fmt.Println("error:", err)

	fmt.Println()

	value, err = hgconfig.Username()
	fmt.Println("value:", value)
	fmt.Println("error:", err)

	fmt.Println()

	value, err = hgconfig.Get("foo.bar.baz")
	fmt.Println("value:", value)
	fmt.Println("error:", err)
}
