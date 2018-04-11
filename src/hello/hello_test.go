package hello_test

import (
	"hello"
	"testing"
)

func TestHello(t *testing.T) {
	h := (hello.NewHello("Hello"))

	go h.ConsumeSays()

	str := h.SayIt("Tom")

	if str != "Hello, Tom." {
		t.Errorf("Hello.sayit isn't \"Hello, Tom.\". It's: %s", str)
	}

	h.SayDelay("\x0225\x01Jerry\x03")
}
