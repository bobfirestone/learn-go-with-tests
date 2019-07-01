package injection

import (
	"bytes"
	"fmt"
)

// Greet says Hello
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
