package heihei_test

import (
	"fmt"
	"heihei"
	"testing"
)

func TestHeiheiBasic(t *testing.T) {
	fmt.Printf("\x02\x01\x03\n")
	heihei.LearnReflect()
}
