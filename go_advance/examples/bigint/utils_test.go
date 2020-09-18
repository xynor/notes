package bigint

import (
	"fmt"
	"testing"
)

func TestStr2Big(t *testing.T) {
	s := "10.3454365464534243"
	b, _ := Str2Big(s, 18)
	bs := Big2Str(b, 18)
	fmt.Println(bs)
}
