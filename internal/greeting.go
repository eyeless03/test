package internal

import "fmt"

func Hello(name string) {
	fmt.Println("привет,", name, "!")
}

func Bye(name string) {
	fmt.Printf("Пока, %s!\n", name)
}
