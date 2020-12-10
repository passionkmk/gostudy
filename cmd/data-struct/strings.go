package main

import (
	"fmt"
)

func main() {
	// MEMO: -
	// Go 는 문자열을 UTF-8로 인코딩, 즉 표현할수 있는 byte수가 가변적이다.
	// 한글의 경우 1글자당 3byte가 필요하다.
	// 따라서 byte 위치인 0, 3, 9가 출력되고
	// 뒤의 숫자는 "가", "나", "다"를 각각 int32의 숫자로 표현한것이다.
	for i, r := range "가나다" {
		fmt.Println(i, r)
	}
	fmt.Println(len("가나다"))

	/**
	0 44032
	3 45208
	6 45796
	9
	**/
}
