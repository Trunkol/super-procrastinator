package main

import (
	"fmt"
	"super-procrastinator/hacknews"
	"super-procrastinator/medium"
)

func main() {
	fmt.Println(medium.Stories())
	fmt.Println(hacknews.Stories())
}
