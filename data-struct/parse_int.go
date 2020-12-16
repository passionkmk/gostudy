package main

import "strconv"

func main() {
	ExampleStrConv()
}

func ExampleStrConv() {
	var i int
	var k int64
	var f float64
	var s string
	var err error

	i, err = strconv.Atoi("350")
}
