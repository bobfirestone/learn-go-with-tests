// Hello world says "Hello, World"

package hello

import "fmt"

// Hello say hello
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Nachos"))
}
