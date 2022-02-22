package main

import "fmt"
import "os"

func main() {
	var a, b, c int
	var calc string

	fmt.Printf("<Formula>\n")
	fmt.Printf("ex) 1 + 1 = ? -> 1 1 +\n")
	fmt.Scanf("%d %d %s", &a, &b, &calc)

	if calc == "+" {
		c = a + b
	} else if calc == "-" {
		c = a - b
	} else if calc == "*" {
		c = a * b
	} else if calc == "/" {
		if b == 0 {
			fmt.Printf("I cannot calculate this formula.\n")
			os.Exit(1)
		} else {
			c = a / b
		}
	}

	fmt.Printf("<Answer>\n")
	fmt.Printf("%d %s %d = %d", a, calc, b, c)
}
