package main

import "fmt"

func f(from string) {
		for i := 0; i < 3; i++ {
				fmt.Println(from, ":", i)
		}
}

func main() {

		f("direct")

		go f("goroutine1")
		go f("goroutine2")

		go func(msg string) {
				fmt.Println(msg)
		}("randum1")

		go func(msg string) {
				fmt.Println(msg)
		}("randum2")

		var input string
		fmt.Scanln(&input)
		fmt.Println("done")
}

