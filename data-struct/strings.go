package main

import (
	"fmt"
	"strings"
)

func main() {
	ExamStrCatStrings()
}

func ExamStrCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
}

func ExamStrCarSprint() {
	s := "abc"
	s = fmt.Sprint(s, "def")
	fmt.Println(s)
}

func ExamStrCatStrings() {
	s := "abc"
	s = strings.Join([]string{s, "def"}, " ")
	fmt.Println(s)
}
