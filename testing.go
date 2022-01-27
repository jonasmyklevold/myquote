package main

import (
	"fmt"
	"github.com/jonasmyklevold/myquote/import"
)

func main() {
	test()
	test1()
}

func test() {
	fmt.Print(ting.TingQuote())
}

func test1() {
	fmt.Print(ting.Ting2Quote())
}
