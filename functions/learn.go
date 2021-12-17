package functions

import (
	"fmt"
	"strings"
)

func add(x, y int) int {
	return x + y
}

func converter(title, email string) string {
	s1 := strings.ToUpper(title)
	return s1 + email
}

func sum(numbers ...int) {
	total := 0
	for _, v := range numbers {
		total =+ v
	}
	fmt.Println(total)
}

func Learn() {
	var result = add(10, 5)
	var message = converter("Hello", "mark@gmail.com")
	sum(10, 20, 30)
	fmt.Println(result)
	fmt.Println(message)
}
