package main

import "fmt"

func main() {
	ExampleModifyBytes()
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

func ExamplePrintBytes2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
}

func ExampleModifyBytes() {
	s := []byte("가나다") // 8bit -> 1byte형 한글은 3바이트
	s[2]++
	fmt.Println(string(s))
}
