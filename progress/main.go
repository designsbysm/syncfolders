package progress

import (
	"fmt"
	"io"
	"os"
)

var title string
var value int
var writer io.Writer

func Clear() {
	fmt.Fprint(writer, "\033[2K\r")
}

func Finish() {
	print()
	fmt.Fprint(writer, "\n")
}

func Increment() {
	value += 1

	if value%100 == 0 {
		print()
	}
}

func Set(msg string) {
	if writer == nil {
		writer = os.Stdout
	}

	title = msg
	value = 0
	print()
}

func print() {
	Clear()
	fmt.Fprintf(writer, "%8s: %d", title, value)
}
