package progress

import (
	"bytes"
	"testing"
)

func TestWriter(t *testing.T) {
	// writer = bytes.NewBuffer([]byte{})
	var buf bytes.Buffer
	writer = &buf

	Set("Test")
	// title="ABC"
	// writer=bufio.NewWriter()
	t.Fatal(string(buf.String()))

}
