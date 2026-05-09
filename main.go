package main

import "git-pr/internal"

func main() {
	internal.Hello("as")
	internal.Bye("as")
	internal.PrintAs()
	as()
}

func as() {
	return
}
