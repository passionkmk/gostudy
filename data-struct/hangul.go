// MARK: - 한글 받침 검사기
package hangul

var (
	start rune = 44032 // '가'의 유니코드
	end   rune = 55204 // '힣' 다음 글자의 유니코드
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false

	return true
}
