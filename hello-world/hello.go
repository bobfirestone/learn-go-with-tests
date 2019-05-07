// Hello world says "Hello, World"

package hello

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello say hello
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}
