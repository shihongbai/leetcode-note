package str

import (
	"fmt"
	"testing"
)

func Test_kmp(t *testing.T) {
	S := "ABABABACABABABCABABABAC"
	P := "ABABAC"

	fmt.Println("Next 数组:", computeNext(P))
	fmt.Println("匹配位置:", kmp(S, P))
}
