package ddz

import (
	"fmt"
	"testing"
)

func TestHandCopy(t *testing.T) {
	cs := CardSlice{
		Club3,
		Diamond3,
	}

	h := HandParse(cs)
	fmt.Println(h)
}
