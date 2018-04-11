package wrapper

import (
	"bufio"
	"os"
)

func LoadFromStdIn() (int, string) {
	var n int = -1

	var str string = ""

	reader := bufio.NewReader(os.Stdin)

	for {
		strBytes, _, err := reader.ReadLine()

		if err != nil {
			break
		}

		str = str + string(strBytes)
	}

	n = len(str)

	return n, str
}
