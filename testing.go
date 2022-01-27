package main

import (
	"fmt"
	"github.com/jonasmyklevold/myquote/import"
)

func main() {
	funksjon()
	funksjon2()
}

func funksjon() {
	fmt.Print(ting.TingQuote())
}

func funksjon2() {
	fmt.Print(ting.Ting2Quote())
}
