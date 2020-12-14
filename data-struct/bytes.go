package main

import "fmt"

func main() {
	ExamplePrintBytes()
}

func ExamplePrintBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	// Output
	// ea:b0:80:eb:82:98:eb:8b:a4
}
