package replacefun

import (
	"strings"
	"fmt"
	"testing"
)

func TestEvalReplace(t *testing.T) {
	in := strings.Join([]string{
		"다들 그 동안 고생이 많았다.",
		"첫째는 분당에 있는 { 2 ** 4 * 3 }평 아파트를 갖거라.",
		"둘째는 임야 { 10 ** 5 mod 7777 }평을 가져라.",
		"막내는 { 10000 - ( 10 ** 5 mod 7777 ) }평 임야를 갖고.",
		"배기량 { 711 * 8 / 9 }cc의 경운기를 갖거라.",
	}, "\n")


	fmt.Println(EvalReplaceAll(in))
	// Output:
	// 다들 그 동안 고생이 많았다.\n첫째는 분당에 있는 48평 아파트를 갖거라.\n둘째는 임야 6676평을 가져라.\n막내는 3324평 임야를 갖고.\n배기량 632cc의 경운기를 갖거라.

}
