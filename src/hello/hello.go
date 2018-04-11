package hello

import "fmt"

type Hello struct {
	greeting string

	queue chan string
}

func NewHello(greeting string) *Hello {
	r := new(Hello)

	r.greeting = greeting

	r.queue = make(chan string, 1024)

	return r
}

func (h *Hello) SayDelay(name string) {
	h.queue <- h.SayIt(name)
}

func (h *Hello) ConsumeSays() {
	var line string

	for {
		line = <-h.queue

		if line == "" {
			break
		}

		fmt.Println(line)
	}
}

func (h *Hello) SayIt(name string) string {
	r := ""

	r = r + "" + h.greeting + ", " + name + "."

	return r
}
